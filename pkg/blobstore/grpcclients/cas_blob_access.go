package grpcclients

import (
	"context"
	"io"
	"log"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/buildbarn/bb-storage/pkg/blobstore"
	"github.com/buildbarn/bb-storage/pkg/blobstore/buffer"
	"github.com/buildbarn/bb-storage/pkg/blobstore/slicing"
	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/buildbarn/bb-storage/pkg/util"
	"github.com/google/uuid"

	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type casBlobAccess struct {
	byteStreamClient                bytestream.ByteStreamClient
	contentAddressableStorageClient remoteexecution.ContentAddressableStorageClient
	capabilitiesClient              remoteexecution.CapabilitiesClient
	uuidGenerator                   util.UUIDGenerator
	readChunkSize                   int
	maximumSizeBytes                int
}

// NewCASBlobAccess creates a BlobAccess handle that relays any requests
// to a GRPC service that implements the bytestream.ByteStream and
// remoteexecution.ContentAddressableStorage services. Those are the
// services that Bazel uses to access blobs stored in the Content
// Addressable Storage.
func NewCASBlobAccess(client grpc.ClientConnInterface, uuidGenerator util.UUIDGenerator, readChunkSize, maximumSizeBytes int) blobstore.BlobAccess {
	return &casBlobAccess{
		byteStreamClient:                bytestream.NewByteStreamClient(client),
		contentAddressableStorageClient: remoteexecution.NewContentAddressableStorageClient(client),
		capabilitiesClient:              remoteexecution.NewCapabilitiesClient(client),
		uuidGenerator:                   uuidGenerator,
		readChunkSize:                   readChunkSize,
		maximumSizeBytes:                maximumSizeBytes,
	}
}

// A resumable bytestream reader.
// This takes the digest and creates the stream when the data is read,
// and can the resume partial reads inside the `Read` function.
type resumableByteStreamChunkReader struct {
	byteStreamClient bytestream.ByteStreamClient
	cancel context.CancelFunc
	ctx context.Context
	digest digest.Digest
	maximumSizeBytes                int
}

func (r *resumableByteStreamChunkReader) Read() ([]byte, error) {
	readOffset := 0
	resumeAttemptCount := 5 // TODO: add configuration.
	chunks := make([]*[]byte, 0, resumeAttemptCount)

	expectedSizeBytes := r.digest.GetSizeBytes()
	if expectedSizeBytes > int64(r.maximumSizeBytes) {
		return nil, status.Errorf(codes.InvalidArgument, "Buffer is %d bytes in size, while a maximum of %d bytes is permitted", expectedSizeBytes, r.maximumSizeBytes)
	}
	data := make([]byte, 0, expectedSizeBytes)
	errors := []error{}

	log.Printf("Resumable bystream read of digest %#v.\n", r.digest)
	for i := 0; i < resumeAttemptCount; i++ {
		// TODO: less ugly do-while control flow.
		if i > 0 {
			if readOffset == 0 {
				// Nothing was read in the first attempt.
				break
				// TODO: error handling.
				// Here, or at the end of the loop?
				// If `readOffset` is not incremented in an attempt we can exit.
			}

			log.Printf("Resuming bytestream read of digest %#v at offset %v.\n", r.digest, readOffset)
		}

		client, err := r.byteStreamClient.Read(r.ctx, &bytestream.ReadRequest{
			ResourceName: r.digest.GetByteStreamReadPath(remoteexecution.Compressor_IDENTITY),
			ReadOffset: int64(readOffset),
		})

		if err != nil {
			r.cancel()
			return nil, util.StatusWrapf(err, "Could not open bytestream for digest %#v.", r.digest)
		}

		// Try to read the full stream
		// Just like `toByteSliceViaChunkReader`.
		withinAttempt := make([]byte, 0, expectedSizeBytes)
		for {
			chunk, err := client.Recv()
			// NB: Presumable the final Read returns empty data and EOF,
			// so we do not need to append anything in that case,
			// symmetric with `toByteSliceViaChunkReader`'s loop.
			if err == io.EOF {
				// Reached the end, can return data.
				return data, nil
			} else if err != nil {
				// Retry with a new stream.
				errors = append(errors, err)
				break
			}
			withinAttempt = append(withinAttempt, chunk.Data...)
		}
		readOffset += len(withinAttempt)

		chunks = append(chunks, &withinAttempt)
		data = append(data, withinAttempt...)
	}

	// Could not resume the download, return all errors seen.
	return nil, util.StatusFromMultiple(errors)
}

func (r *resumableByteStreamChunkReader) Close() {
	// This reader has no background state, the stream is closed within `Read`.
}

type byteStreamChunkReader struct {
	client bytestream.ByteStream_ReadClient
	cancel context.CancelFunc
}

func (r *byteStreamChunkReader) Read() ([]byte, error) {
	chunk, err := r.client.Recv()
	if err != nil {
		return nil, err
	}
	return chunk.Data, nil
}

func (r *byteStreamChunkReader) Close() {
	r.cancel()
	for {
		if _, err := r.client.Recv(); err != nil {
			break
		}
	}
}

func (ba *casBlobAccess) Get(ctx context.Context, digest digest.Digest) buffer.Buffer {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	return buffer.NewCASBufferFromChunkReader(digest, &resumableByteStreamChunkReader{
		cancel: cancel,
		byteStreamClient: ba.byteStreamClient,
		ctx: ctxWithCancel,
		digest: digest,
		maximumSizeBytes:                ba.maximumSizeBytes,
	}, buffer.BackendProvided(buffer.Irreparable(digest)))
}

func (ba *casBlobAccess) GetFromComposite(ctx context.Context, parentDigest, childDigest digest.Digest, slicer slicing.BlobSlicer) buffer.Buffer {
	b, _ := slicer.Slice(ba.Get(ctx, parentDigest), childDigest)
	return b
}

func (ba *casBlobAccess) Put(ctx context.Context, digest digest.Digest, b buffer.Buffer) error {
	r := b.ToChunkReader(0, ba.readChunkSize)
	defer r.Close()

	ctxWithCancel, cancel := context.WithCancel(ctx)
	client, err := ba.byteStreamClient.Write(ctxWithCancel)
	if err != nil {
		cancel()
		return err
	}

	resourceName := digest.GetByteStreamWritePath(uuid.Must(ba.uuidGenerator()), remoteexecution.Compressor_IDENTITY)
	writeOffset := int64(0)
	for {
		if data, err := r.Read(); err == nil {
			// Non-terminating chunk.
			if client.Send(&bytestream.WriteRequest{
				ResourceName: resourceName,
				WriteOffset:  writeOffset,
				Data:         data,
			}) != nil {
				cancel()
				_, err := client.CloseAndRecv()
				return err
			}
			writeOffset += int64(len(data))
			resourceName = ""
		} else if err == io.EOF {
			// Terminating chunk.
			if client.Send(&bytestream.WriteRequest{
				ResourceName: resourceName,
				WriteOffset:  writeOffset,
				FinishWrite:  true,
			}) != nil {
				cancel()
				_, err := client.CloseAndRecv()
				return err
			}
			_, err := client.CloseAndRecv()
			cancel()
			return err
		} else if err != nil {
			cancel()
			client.CloseAndRecv()
			return err
		}
	}
}

func (ba *casBlobAccess) FindMissing(ctx context.Context, digests digest.Set) (digest.Set, error) {
	// Partition all digests by digest function, as the
	// FindMissingBlobs() RPC can only process digests for a single
	// instance name and digest function.
	perFunctionDigests := map[digest.Function][]*remoteexecution.Digest{}
	for _, digest := range digests.Items() {
		digestFunction := digest.GetDigestFunction()
		perFunctionDigests[digestFunction] = append(perFunctionDigests[digestFunction], digest.GetProto())
	}

	missingDigests := digest.NewSetBuilder()
	for digestFunction, blobDigests := range perFunctionDigests {
		// Call FindMissingBlobs() for each digest function.
		request := remoteexecution.FindMissingBlobsRequest{
			InstanceName:   digestFunction.GetInstanceName().String(),
			BlobDigests:    blobDigests,
			DigestFunction: digestFunction.GetEnumValue(),
		}
		response, err := ba.contentAddressableStorageClient.FindMissingBlobs(ctx, &request)
		if err != nil {
			return digest.EmptySet, err
		}

		// Convert results back.
		for _, proto := range response.MissingBlobDigests {
			blobDigest, err := digestFunction.NewDigestFromProto(proto)
			if err != nil {
				return digest.EmptySet, err
			}
			missingDigests.Add(blobDigest)
		}
	}
	return missingDigests.Build(), nil
}

func (ba *casBlobAccess) GetCapabilities(ctx context.Context, instanceName digest.InstanceName) (*remoteexecution.ServerCapabilities, error) {
	cacheCapabilities, err := getCacheCapabilities(ctx, ba.capabilitiesClient, instanceName)
	if err != nil {
		return nil, err
	}

	// Only return fields that pertain to the Content Addressable
	// Storage. Don't set 'max_batch_total_size_bytes', as we don't
	// issue batch operations. The same holds for fields related to
	// compression support.
	return &remoteexecution.ServerCapabilities{
		CacheCapabilities: &remoteexecution.CacheCapabilities{
			DigestFunctions: digest.RemoveUnsupportedDigestFunctions(cacheCapabilities.DigestFunctions),
		},
	}, nil
}
