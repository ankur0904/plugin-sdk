// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/pb/destination.proto

package pb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DestinationClient is the client API for Destination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DestinationClient interface {
	// Get an example configuration for the source plugin
	GetExampleConfig(ctx context.Context, in *GetExampleConfig_Request, opts ...grpc.CallOption) (*GetExampleConfig_Response, error)
	// Configure the destination plugin with the given credentials and mode
	Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error)
	// Migrate tables to the given source plugin version
	Migrate(ctx context.Context, in *Migrate_Request, opts ...grpc.CallOption) (*Migrate_Response, error)
	// Write resources
	Write(ctx context.Context, opts ...grpc.CallOption) (Destination_WriteClient, error)
}

type destinationClient struct {
	cc grpc.ClientConnInterface
}

func NewDestinationClient(cc grpc.ClientConnInterface) DestinationClient {
	return &destinationClient{cc}
}

func (c *destinationClient) GetExampleConfig(ctx context.Context, in *GetExampleConfig_Request, opts ...grpc.CallOption) (*GetExampleConfig_Response, error) {
	out := new(GetExampleConfig_Response)
	err := c.cc.Invoke(ctx, "/proto.Destination/GetExampleConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error) {
	out := new(Configure_Response)
	err := c.cc.Invoke(ctx, "/proto.Destination/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Migrate(ctx context.Context, in *Migrate_Request, opts ...grpc.CallOption) (*Migrate_Response, error) {
	out := new(Migrate_Response)
	err := c.cc.Invoke(ctx, "/proto.Destination/Migrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Write(ctx context.Context, opts ...grpc.CallOption) (Destination_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Destination_ServiceDesc.Streams[0], "/proto.Destination/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &destinationWriteClient{stream}
	return x, nil
}

type Destination_WriteClient interface {
	Send(*Write_Request) error
	CloseAndRecv() (*Write_Response, error)
	grpc.ClientStream
}

type destinationWriteClient struct {
	grpc.ClientStream
}

func (x *destinationWriteClient) Send(m *Write_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *destinationWriteClient) CloseAndRecv() (*Write_Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Write_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DestinationServer is the server API for Destination service.
// All implementations must embed UnimplementedDestinationServer
// for forward compatibility
type DestinationServer interface {
	// Get an example configuration for the source plugin
	GetExampleConfig(context.Context, *GetExampleConfig_Request) (*GetExampleConfig_Response, error)
	// Configure the destination plugin with the given credentials and mode
	Configure(context.Context, *Configure_Request) (*Configure_Response, error)
	// Migrate tables to the given source plugin version
	Migrate(context.Context, *Migrate_Request) (*Migrate_Response, error)
	// Write resources
	Write(Destination_WriteServer) error
	mustEmbedUnimplementedDestinationServer()
}

// UnimplementedDestinationServer must be embedded to have forward compatible implementations.
type UnimplementedDestinationServer struct {
}

func (UnimplementedDestinationServer) GetExampleConfig(context.Context, *GetExampleConfig_Request) (*GetExampleConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleConfig not implemented")
}
func (UnimplementedDestinationServer) Configure(context.Context, *Configure_Request) (*Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedDestinationServer) Migrate(context.Context, *Migrate_Request) (*Migrate_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migrate not implemented")
}
func (UnimplementedDestinationServer) Write(Destination_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedDestinationServer) mustEmbedUnimplementedDestinationServer() {}

// UnsafeDestinationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DestinationServer will
// result in compilation errors.
type UnsafeDestinationServer interface {
	mustEmbedUnimplementedDestinationServer()
}

func RegisterDestinationServer(s grpc.ServiceRegistrar, srv DestinationServer) {
	s.RegisterService(&Destination_ServiceDesc, srv)
}

func _Destination_GetExampleConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExampleConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).GetExampleConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destination/GetExampleConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).GetExampleConfig(ctx, req.(*GetExampleConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destination/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Configure(ctx, req.(*Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Migrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Migrate_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Migrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destination/Migrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Migrate(ctx, req.(*Migrate_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DestinationServer).Write(&destinationWriteServer{stream})
}

type Destination_WriteServer interface {
	SendAndClose(*Write_Response) error
	Recv() (*Write_Request, error)
	grpc.ServerStream
}

type destinationWriteServer struct {
	grpc.ServerStream
}

func (x *destinationWriteServer) SendAndClose(m *Write_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *destinationWriteServer) Recv() (*Write_Request, error) {
	m := new(Write_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Destination_ServiceDesc is the grpc.ServiceDesc for Destination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Destination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Destination",
	HandlerType: (*DestinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExampleConfig",
			Handler:    _Destination_GetExampleConfig_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Destination_Configure_Handler,
		},
		{
			MethodName: "Migrate",
			Handler:    _Destination_Migrate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _Destination_Write_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/pb/destination.proto",
}
