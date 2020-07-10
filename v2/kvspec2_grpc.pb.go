// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kvspec2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PublicClient is the client API for Public service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicClient interface {
	Query(ctx context.Context, in *ObjectReader, opts ...grpc.CallOption) (*ObjectResult, error)
	Commit(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error)
	BatchCommit(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResult, error)
}

type publicClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicClient(cc grpc.ClientConnInterface) PublicClient {
	return &publicClient{cc}
}

func (c *publicClient) Query(ctx context.Context, in *ObjectReader, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := c.cc.Invoke(ctx, "/kvspec2.Public/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) Commit(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := c.cc.Invoke(ctx, "/kvspec2.Public/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicClient) BatchCommit(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResult, error) {
	out := new(BatchResult)
	err := c.cc.Invoke(ctx, "/kvspec2.Public/BatchCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServer is the server API for Public service.
// All implementations must embed UnimplementedPublicServer
// for forward compatibility
type PublicServer interface {
	Query(context.Context, *ObjectReader) (*ObjectResult, error)
	Commit(context.Context, *ObjectWriter) (*ObjectResult, error)
	BatchCommit(context.Context, *BatchRequest) (*BatchResult, error)
	mustEmbedUnimplementedPublicServer()
}

// UnimplementedPublicServer must be embedded to have forward compatible implementations.
type UnimplementedPublicServer struct {
}

func (*UnimplementedPublicServer) Query(context.Context, *ObjectReader) (*ObjectResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedPublicServer) Commit(context.Context, *ObjectWriter) (*ObjectResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (*UnimplementedPublicServer) BatchCommit(context.Context, *BatchRequest) (*BatchResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCommit not implemented")
}
func (*UnimplementedPublicServer) mustEmbedUnimplementedPublicServer() {}

func RegisterPublicServer(s *grpc.Server, srv PublicServer) {
	s.RegisterService(&_Public_serviceDesc, srv)
}

func _Public_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectReader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvspec2.Public/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).Query(ctx, req.(*ObjectReader))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWriter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvspec2.Public/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).Commit(ctx, req.(*ObjectWriter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Public_BatchCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServer).BatchCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvspec2.Public/BatchCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServer).BatchCommit(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Public_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvspec2.Public",
	HandlerType: (*PublicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Public_Query_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Public_Commit_Handler,
		},
		{
			MethodName: "BatchCommit",
			Handler:    _Public_BatchCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvspec2.proto",
}

// InternalClient is the client API for Internal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalClient interface {
	Prepare(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error)
	Accept(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error)
}

type internalClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalClient(cc grpc.ClientConnInterface) InternalClient {
	return &internalClient{cc}
}

func (c *internalClient) Prepare(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := c.cc.Invoke(ctx, "/kvspec2.Internal/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) Accept(ctx context.Context, in *ObjectWriter, opts ...grpc.CallOption) (*ObjectResult, error) {
	out := new(ObjectResult)
	err := c.cc.Invoke(ctx, "/kvspec2.Internal/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServer is the server API for Internal service.
// All implementations must embed UnimplementedInternalServer
// for forward compatibility
type InternalServer interface {
	Prepare(context.Context, *ObjectWriter) (*ObjectResult, error)
	Accept(context.Context, *ObjectWriter) (*ObjectResult, error)
	mustEmbedUnimplementedInternalServer()
}

// UnimplementedInternalServer must be embedded to have forward compatible implementations.
type UnimplementedInternalServer struct {
}

func (*UnimplementedInternalServer) Prepare(context.Context, *ObjectWriter) (*ObjectResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (*UnimplementedInternalServer) Accept(context.Context, *ObjectWriter) (*ObjectResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (*UnimplementedInternalServer) mustEmbedUnimplementedInternalServer() {}

func RegisterInternalServer(s *grpc.Server, srv InternalServer) {
	s.RegisterService(&_Internal_serviceDesc, srv)
}

func _Internal_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWriter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvspec2.Internal/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).Prepare(ctx, req.(*ObjectWriter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWriter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvspec2.Internal/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).Accept(ctx, req.(*ObjectWriter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Internal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvspec2.Internal",
	HandlerType: (*InternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Internal_Prepare_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _Internal_Accept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvspec2.proto",
}
