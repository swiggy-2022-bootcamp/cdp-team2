package grpc

import (
	"context"
	model "customers/internal/dao"
	pb "customers/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
	"customers/internal/literals"
	"strconv"
)

func SendAddress(address model.Address, customer_id string) bool {
	conn, _ := grpc.Dial("localhost:"+strconv.Itoa(literals.ADDRESS_PORT), grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	fmt.Println("1")
	return SendAddressService(c, address,customer_id)
}

func SendAddressService(c pb.ServiceClient, address model.Address, customer_id string) bool {
	addressRequest := pb.AddressutsRequest{
		Address: &pb.Address{
			CustomerId: customer_id,
			Firstname:  address.Firstname,
			Lastname:   address.Lastname,
			Address1:   address.Address1,
			Address2:   address.Address2,
			City:       address.City,
			Postcode:   address.Postcode,
			CountryId:   address.CountryId,
		},
	}
	fmt.Println(2)
	c.AddressutsService(context.Background(), &addressRequest)
	return true
}
 