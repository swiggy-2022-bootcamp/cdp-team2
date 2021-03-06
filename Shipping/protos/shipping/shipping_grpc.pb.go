// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shipping

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	AddressutsService(ctx context.Context, in *AddressutsRequest, opts ...grpc.CallOption) (*AddressutsResponse, error)
	AddressstuService(ctx context.Context, in *AddressstuRequest, opts ...grpc.CallOption) (*AddressstuResponse, error)
	OrderAddressUpdateService(ctx context.Context, in *OrderAddressUpdateRequest, opts ...grpc.CallOption) (*OrderAddressUpdateResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) AddressutsService(ctx context.Context, in *AddressutsRequest, opts ...grpc.CallOption) (*AddressutsResponse, error) {
	out := new(AddressutsResponse)
	err := c.cc.Invoke(ctx, "/proto.Service/AddressutsService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AddressstuService(ctx context.Context, in *AddressstuRequest, opts ...grpc.CallOption) (*AddressstuResponse, error) {
	out := new(AddressstuResponse)
	err := c.cc.Invoke(ctx, "/proto.Service/AddressstuService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) OrderAddressUpdateService(ctx context.Context, in *OrderAddressUpdateRequest, opts ...grpc.CallOption) (*OrderAddressUpdateResponse, error) {
	out := new(OrderAddressUpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.Service/OrderAddressUpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	AddressutsService(context.Context, *AddressutsRequest) (*AddressutsResponse, error)
	AddressstuService(context.Context, *AddressstuRequest) (*AddressstuResponse, error)
	OrderAddressUpdateService(context.Context, *OrderAddressUpdateRequest) (*OrderAddressUpdateResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) AddressutsService(context.Context, *AddressutsRequest) (*AddressutsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressutsService not implemented")
}
func (UnimplementedServiceServer) AddressstuService(context.Context, *AddressstuRequest) (*AddressstuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressstuService not implemented")
}
func (UnimplementedServiceServer) OrderAddressUpdateService(context.Context, *OrderAddressUpdateRequest) (*OrderAddressUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderAddressUpdateService not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_AddressutsService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressutsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddressutsService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/AddressutsService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddressutsService(ctx, req.(*AddressutsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AddressstuService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressstuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AddressstuService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/AddressstuService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AddressstuService(ctx, req.(*AddressstuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_OrderAddressUpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderAddressUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).OrderAddressUpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Service/OrderAddressUpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).OrderAddressUpdateService(ctx, req.(*OrderAddressUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddressutsService",
			Handler:    _Service_AddressutsService_Handler,
		},
		{
			MethodName: "AddressstuService",
			Handler:    _Service_AddressstuService_Handler,
		},
		{
			MethodName: "OrderAddressUpdateService",
			Handler:    _Service_OrderAddressUpdateService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/shipping.proto",
}
