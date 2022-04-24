package grpc_client

import (
	"context"
	"fmt"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/config"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/protos/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetAddress(orderId string, addressId int) (*pb.OrderAddressUpdateResponse, error) {
	conn, err := grpc.Dial(config.GrpcAdd["ORDER_SERVICE"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	return SetAddressService(c, orderId, addressId)
}

func SetAddressService(c pb.ServiceClient, orderId string, addressId int) (*pb.OrderAddressUpdateResponse, error) {
	addressRequest := pb.OrderAddressUpdateRequest{
		OrderId:   orderId,
		AddressId: int32(addressId),
	}
	resp, err := c.OrderAddressUpdateService(context.Background(), &addressRequest)
	fmt.Println(resp, err)
	if err != nil {
		return &pb.OrderAddressUpdateResponse{
			Response: false,
		}, nil
	}

	return resp, err
}
