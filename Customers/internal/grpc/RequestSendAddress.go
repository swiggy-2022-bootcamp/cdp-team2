package grpc

import (
	"context"
	model "customers/dao/models"
	pb "customers/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
	// "customers/internal/literals"
	// "strconv"
		"google.golang.org/grpc/credentials/insecure"

)

func SendAddress(address model.Address, customer_id string) bool {
	// conn, _ := grpc.Dial("localhost:"+strconv.Itoa(literals.ADDRESS_PORT), grpc.WithInsecure())
conn,_:=grpc.Dial("tcp://0.tcp.in.ngrok.io:18160 ",grpc.WithTransportCredentials(insecure.NewCredentials()))

	// conn, _ := grpc.Dial("13b3-49-37-152-145.in.ngrok.io", grpc.WithInsecure())

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
			AddressLine1:   address.AddressLine1,
			AddressLine2:   address.AddressLine2,
			City:       address.City,
			PostCode:   address.PostCode,
			CountryCode:   address.CountryCode,
		},
	}
	fmt.Println(2)
	resp,_:=c.AddressutsService(context.Background(), &addressRequest)
	fmt.Println(resp)
 	return true
}
 
// func main(){
// 	fmt.Println(SendAddress(model.Address{"iday","kiran","erdfg","wesdr","sed",1234,"1234"},"udayid"))
// }