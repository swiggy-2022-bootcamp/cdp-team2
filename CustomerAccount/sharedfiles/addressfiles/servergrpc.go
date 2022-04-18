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
