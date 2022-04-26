package grpc_client

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/api"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/config"
	pb2 "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/protos/auth"
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

func VerifyToken(c2 *gin.Context, token string) (*pb2.VerifyResponse, error) {
	conn, err := grpc.Dial(config.GrpcAdd["AUTH_SERVICE"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c2.AbortWithStatusJSON(api.UserErr(fmt.Errorf("Invalid token")))
		fmt.Println(err)
	}
	c := pb2.NewAuthServiceClient(conn)
	return VerifyTokenService(c, token)

}

func VerifyTokenService(c pb2.AuthServiceClient, token string) (*pb2.VerifyResponse, error) {
	verifyRequest := pb2.VerifyRequest{
		Token: token,
	}
	resp, err := c.Verify(context.Background(), &verifyRequest)
	if err != nil {
		return &pb2.VerifyResponse{
			Proceed:    resp.Proceed,
			CustomerId: resp.CustomerId,
		}, nil
	}

	return resp, err
}
