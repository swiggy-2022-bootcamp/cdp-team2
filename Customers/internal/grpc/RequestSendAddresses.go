package main

import (
	"context"
	"customers/dao/models"
	pb "customers/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
	// "customers/internal/literals"
	// "strconv"
)

func GetAddress( customerId string)([]models.Address) {
 	// conn, _ := grpc.Dial("localhost:"+strconv.Itoa(literals.ADDRESS_PORT), grpc.WithInsecure())
	 	conn, _ := grpc.Dial("localhost:9010", grpc.WithInsecure())

	defer conn.Close()
	c := pb.NewServiceClient(conn)
 	return GetAddressService(c, customerId)
}

func GetAddressService(c pb.ServiceClient, customerId string)( []models.Address ){
	addressRequest := pb.AddressstuRequest{
		CustomerId:customerId,
	}
	fmt.Println(2)
	resp,err:=c.AddressstuService(context.Background(), &addressRequest)
	if err!=nil{
		return []models.Address{};
	}
	addresses:=[]models.Address{}
	for _,val:=range resp.Address{
		addresses=append(addresses,models.Address{val.Firstname,val.Lastname,val.AddressLine1,val.AddressLine2,val.City,val.PostCode,val.CountryCode,})
	}
	return addresses
}
 
func main(){
	fmt.Println(GetAddress("frgt"))
}