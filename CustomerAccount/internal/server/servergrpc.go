package server

import (
	"context"
	// "customer-account/internal/literals"
	pb "customer-account/internal/proto"
	"fmt"
	"net"
	// "strconv"

	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceServer
}

func (*server) AddressutsService(ctx context.Context, req *pb.AddressutsRequest) (*pb.AddressutsResponse, error) {
	customer_id := req.Address.CustomerId
	firstname := req.Address.Firstname
	lastname := req.Address.Lastname
	address1 := req.Address.Address1
	address2 := req.Address.Address2
	city := req.Address.City
	postcode := req.Address.Postcode
	country_id := req.Address.CountryId

	fmt.Println(customer_id, firstname, lastname, address1, address2, city, postcode, country_id)
	return &(pb.AddressutsResponse{Response: true}), nil
}
func (*server) CartService(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {
	fmt.Println("Cart Service called")
	id_s := req.CustomerId
	fmt.Println(id_s)

	//  User your database here to fetch products in the cart

	// fetch products here of the customer whose id is id_s
	// convert golang data types into grpc types

	//

	products := []*pb.Product{}
	products = append(products, &pb.Product{ProductId:"id1", Quantity: int32(2)})
	products = append(products, &pb.Product{ProductId: "id2", Quantity: int32(5)})
	fmt.Println(products)
	result1 := pb.Cart{
		Product: products,
	}
	cartResponse := pb.CartResponse{
		Cart: &result1,
	}
	return &(cartResponse), nil
}

func (*server) RewardService(ctx context.Context, req *pb.RewardRequest) (*pb.RewardResponse, error) {
	fmt.Println("Reward Service called")
	id_s := req.CustomerId
	fmt.Println(id_s)

	//  User your database here to fetch products in the reward

	// fetch reward here of the customer whose id is id_s
	// convert golang data types into grpc types

	//

	reward := pb.Reward{CustomerId: "12345", Description: "Silver level points", Points: int32(32)}

	rewardResponse := pb.RewardResponse{
		Reward: &reward,
	}
	return &(rewardResponse), nil
}

func StartGrpc() {

	fmt.Println("GRPC Server started")
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterServiceServer(grpcServer, &server{})
	Listener, err := net.Listen("tcp", ":9000")
	if err!=nil{
		fmt.Println("Failed to listen to the server",err)
	}
	if err:=grpcServer.Serve(Listener); err!=nil{
		fmt.Println("Error------",err)
	}

}
