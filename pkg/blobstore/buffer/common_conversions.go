package buffer

import (
	"io"

	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/buildbarn/bb-storage/pkg/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func intoWriterViaChunkReader(r ChunkReader, w io.Writer) error {
	defer r.Close()

	for {
		chunk, err := r.Read()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		if _, err := w.Write(chunk); err != nil {
			return err
		}
	}
}

func readAtViaChunkReader(r ChunkReader, p []byte, off int64) (int, error) {
	r = newOffsetChunkReader(r, off)
	defer r.Close()

	// Copy chunks into the output buffer.
	nTotal := 0
	for len(p) > 0 {
		chunk, err := r.Read()
		if err == io.EOF {
			return nTotal, err
		}
		if err != nil {
			return 0, err
		}
		nCopied := copy(p, chunk)
		nTotal += nCopied
		p = p[nCopied:]
	}

	// Continue reading the rest of the chunk to force checksum
	// validation.
	for {
		_, err := r.Read()
		if err == io.EOF {
			return nTotal, nil
		}
		if err != nil {
			return 0, err
		}
	}
}

func toByteSliceViaChunkReader(r ChunkReader, digest digest.Digest, maximumSizeBytes int) ([]byte, error) {
	defer r.Close()

	expectedSizeBytes := digest.GetSizeBytes()
	if expectedSizeBytes > int64(maximumSizeBytes) {
		return nil, status.Errorf(codes.InvalidArgument, "Buffer is %d bytes in size, while a maximum of %d bytes is permitted", expectedSizeBytes, maximumSizeBytes)
	}

	data := make([]byte, 0, expectedSizeBytes)
	for {
		// NB: The `digest` is embedded in the ChunkReader, so we do not need it here.
		chunk, err := r.Read()
		if err == io.EOF {
			return data, nil
		} else if err != nil {
			return nil, err
		}
		data = append(data, chunk...)
	}
}

func cloneCopyViaByteSlice(b Buffer, maximumSizeBytes int) (Buffer, Buffer) {
	data, err := b.ToByteSlice(maximumSizeBytes)
	if err != nil {
		return NewBufferFromError(err).CloneCopy(maximumSizeBytes)
	}
	return NewValidatedBufferFromByteSlice(data).CloneCopy(maximumSizeBytes)
}

func toProtoViaByteSlice(b Buffer, m proto.Message, maximumSizeBytes int) (proto.Message, error) {
	data, err := b.ToByteSlice(maximumSizeBytes)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, m); err != nil {
		return nil, util.StatusWrapWithCode(err, codes.InvalidArgument, "Failed to unmarshal message")
	}
	return m, nil
}
