package grpc_client

import (
	"context"
	"fmt"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/config"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/protos/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//func SetAddressClient(ctx context.Context) (*CartClient, error) {
//	conn, err := grpc.Dial(config.GrpcAdd["CART_SERVICE"], grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Printf("Error creating grpc cart Client : %s", err.Error())
//		return nil, err
//	}
//
//	cartClient := cartpb.NewCartServiceClient(conn)
//
//	ctxx, cancel := context.WithCancel(ctx)
//
//}

func SetAddress(orderId string, addressId int) (*pb.OrderAddressUpdateResponse, error) {
	// conn, _ := grpc.Dial("localhost:"+strconv.Itoa(literals.ADDRESS_PORT), grpc.WithInsecure())
	fmt.Println("config.GrpcAdd[\"CART_SERVICE\"]", config.GrpcAdd["CART_SERVICE"])
	conn, err := grpc.Dial("0.tcp.in.ngrok.io:18600", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		return nil, nil
	}

	return resp, err
}
