// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: pkg/proto/iscc/iscc.proto

package iscc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	InitialSizeClassCache_GetPreviousExecutionStats_FullMethodName    = "/buildbarn.iscc.InitialSizeClassCache/GetPreviousExecutionStats"
	InitialSizeClassCache_UpdatePreviousExecutionStats_FullMethodName = "/buildbarn.iscc.InitialSizeClassCache/UpdatePreviousExecutionStats"
)

// InitialSizeClassCacheClient is the client API for InitialSizeClassCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InitialSizeClassCacheClient interface {
	GetPreviousExecutionStats(ctx context.Context, in *GetPreviousExecutionStatsRequest, opts ...grpc.CallOption) (*PreviousExecutionStats, error)
	UpdatePreviousExecutionStats(ctx context.Context, in *UpdatePreviousExecutionStatsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type initialSizeClassCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewInitialSizeClassCacheClient(cc grpc.ClientConnInterface) InitialSizeClassCacheClient {
	return &initialSizeClassCacheClient{cc}
}

func (c *initialSizeClassCacheClient) GetPreviousExecutionStats(ctx context.Context, in *GetPreviousExecutionStatsRequest, opts ...grpc.CallOption) (*PreviousExecutionStats, error) {
	out := new(PreviousExecutionStats)
	err := c.cc.Invoke(ctx, InitialSizeClassCache_GetPreviousExecutionStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *initialSizeClassCacheClient) UpdatePreviousExecutionStats(ctx context.Context, in *UpdatePreviousExecutionStatsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, InitialSizeClassCache_UpdatePreviousExecutionStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitialSizeClassCacheServer is the server API for InitialSizeClassCache service.
// All implementations should embed UnimplementedInitialSizeClassCacheServer
// for forward compatibility
type InitialSizeClassCacheServer interface {
	GetPreviousExecutionStats(context.Context, *GetPreviousExecutionStatsRequest) (*PreviousExecutionStats, error)
	UpdatePreviousExecutionStats(context.Context, *UpdatePreviousExecutionStatsRequest) (*emptypb.Empty, error)
}

// UnimplementedInitialSizeClassCacheServer should be embedded to have forward compatible implementations.
type UnimplementedInitialSizeClassCacheServer struct {
}

func (UnimplementedInitialSizeClassCacheServer) GetPreviousExecutionStats(context.Context, *GetPreviousExecutionStatsRequest) (*PreviousExecutionStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviousExecutionStats not implemented")
}
func (UnimplementedInitialSizeClassCacheServer) UpdatePreviousExecutionStats(context.Context, *UpdatePreviousExecutionStatsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreviousExecutionStats not implemented")
}

// UnsafeInitialSizeClassCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InitialSizeClassCacheServer will
// result in compilation errors.
type UnsafeInitialSizeClassCacheServer interface {
	mustEmbedUnimplementedInitialSizeClassCacheServer()
}

func RegisterInitialSizeClassCacheServer(s grpc.ServiceRegistrar, srv InitialSizeClassCacheServer) {
	s.RegisterService(&InitialSizeClassCache_ServiceDesc, srv)
}

func _InitialSizeClassCache_GetPreviousExecutionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviousExecutionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitialSizeClassCacheServer).GetPreviousExecutionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InitialSizeClassCache_GetPreviousExecutionStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitialSizeClassCacheServer).GetPreviousExecutionStats(ctx, req.(*GetPreviousExecutionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InitialSizeClassCache_UpdatePreviousExecutionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePreviousExecutionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitialSizeClassCacheServer).UpdatePreviousExecutionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InitialSizeClassCache_UpdatePreviousExecutionStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitialSizeClassCacheServer).UpdatePreviousExecutionStats(ctx, req.(*UpdatePreviousExecutionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InitialSizeClassCache_ServiceDesc is the grpc.ServiceDesc for InitialSizeClassCache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InitialSizeClassCache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.iscc.InitialSizeClassCache",
	HandlerType: (*InitialSizeClassCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPreviousExecutionStats",
			Handler:    _InitialSizeClassCache_GetPreviousExecutionStats_Handler,
		},
		{
			MethodName: "UpdatePreviousExecutionStats",
			Handler:    _InitialSizeClassCache_UpdatePreviousExecutionStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/iscc/iscc.proto",
}
