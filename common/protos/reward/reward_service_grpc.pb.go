// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: reward_service.proto

package reward

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

// RewardServiceClient is the client API for RewardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RewardServiceClient interface {
	// Get the total reward currently associated with a customer with given customerId
	GetReward(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardResponse, error)
	// Calculates the final pricing after deducting the reward/points from the total price for the user
	RedeemReward(ctx context.Context, in *RedeemRewardRequest, opts ...grpc.CallOption) (*RedeemRewardResponse, error)
}

type rewardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRewardServiceClient(cc grpc.ClientConnInterface) RewardServiceClient {
	return &rewardServiceClient{cc}
}

func (c *rewardServiceClient) GetReward(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardResponse, error) {
	out := new(RewardResponse)
	err := c.cc.Invoke(ctx, "/reward.RewardService/GetReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardServiceClient) RedeemReward(ctx context.Context, in *RedeemRewardRequest, opts ...grpc.CallOption) (*RedeemRewardResponse, error) {
	out := new(RedeemRewardResponse)
	err := c.cc.Invoke(ctx, "/reward.RewardService/RedeemReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RewardServiceServer is the server API for RewardService service.
// All implementations must embed UnimplementedRewardServiceServer
// for forward compatibility
type RewardServiceServer interface {
	// Get the total reward currently associated with a customer with given customerId
	GetReward(context.Context, *RewardRequest) (*RewardResponse, error)
	// Calculates the final pricing after deducting the reward/points from the total price for the user
	RedeemReward(context.Context, *RedeemRewardRequest) (*RedeemRewardResponse, error)
	mustEmbedUnimplementedRewardServiceServer()
}

// UnimplementedRewardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRewardServiceServer struct {
}

func (UnimplementedRewardServiceServer) GetReward(context.Context, *RewardRequest) (*RewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReward not implemented")
}
func (UnimplementedRewardServiceServer) RedeemReward(context.Context, *RedeemRewardRequest) (*RedeemRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemReward not implemented")
}
func (UnimplementedRewardServiceServer) mustEmbedUnimplementedRewardServiceServer() {}

// UnsafeRewardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RewardServiceServer will
// result in compilation errors.
type UnsafeRewardServiceServer interface {
	mustEmbedUnimplementedRewardServiceServer()
}

func RegisterRewardServiceServer(s grpc.ServiceRegistrar, srv RewardServiceServer) {
	s.RegisterService(&RewardService_ServiceDesc, srv)
}

func _RewardService_GetReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServiceServer).GetReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reward.RewardService/GetReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServiceServer).GetReward(ctx, req.(*RewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RewardService_RedeemReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeemRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServiceServer).RedeemReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reward.RewardService/RedeemReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServiceServer).RedeemReward(ctx, req.(*RedeemRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RewardService_ServiceDesc is the grpc.ServiceDesc for RewardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RewardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reward.RewardService",
	HandlerType: (*RewardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReward",
			Handler:    _RewardService_GetReward_Handler,
		},
		{
			MethodName: "RedeemReward",
			Handler:    _RewardService_RedeemReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reward_service.proto",
}