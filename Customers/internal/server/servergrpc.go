package server

import (
 	pb "customers/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
	"net"
	"context"
 	// "strconv"
	"customers/config"
	"customers/internal/handlegrpc"
	// "customers/internal/literals"
	// db "customers/config"
)
type server struct{
	pb.UnimplementedServiceServer
}

func (*server)  CredentialService(ctx context.Context,req *pb.CredentialRequest)(*pb.CredentialResponse,error){
	username:=req.Credential.Username
	password:=req.Credential.Password
	fmt.Println("Credential Service called username and password are :",username,password)
 	customer_id,ispresent:=handlegrpc.CheckCredentials(username,password)
	return &(pb.CredentialResponse{Ispresent:ispresent,Customerid:customer_id}),nil
}
func (*server) AddressstuService(ctx context.Context,req *pb.AddressstuRequest)(*pb.AddressstuResponse,error){
	customer_id:=req.CustomerId;
	fmt.Println(customer_id)

	// fetch the data from DB

	addresses:=[]*pb.Address{}
	// write a for loop to add the values into this array
	addresses=append(addresses,&pb.Address{CustomerId:"rf837fh",Firstname:"uday",Lastname:"Kiran",AddressLine1:"house numer 1",AddressLine2:"ngo street",City:"kamareddy",PostCode:503111,CountryCode:"+91"})
	return &(pb.AddressstuResponse{Address:addresses}),nil
}

func (*server)  AddressutsService(ctx context.Context,req *pb.AddressutsRequest)(*pb.AddressutsResponse,error){
	fmt.Println("address service called")
	
	customerId:=req.Address.CustomerId
	firstname:=req.Address.Firstname
	lastname:=req.Address.Lastname
	addressLine1:=req.Address.AddressLine1
	addressLine2:=req.Address.AddressLine2
	city:=req.Address.City
	postcode:=req.Address.PostCode
	countryId:=req.Address.CountryCode

	fmt.Println(customerId,firstname,lastname,addressLine1,addressLine2,city,postcode,countryId);
	resp:=&(pb.AddressutsResponse{Response:true})
 	return resp,nil
}
func(*server) CartService(ctx context.Context,req *pb.CartRequest )(*pb.CartResponse,error){
    fmt.Println("Cart Service called")
    id_s:=req.GetCustomerId().GetId()
	fmt.Println(id_s)

	//  User your database here to fetch products in the cart

	// fetch products here of the customer whose id is id_s
	// convert golang data types into grpc types   

	//


	products:=[]*pb.Product{};
	products=append(products,&pb.Product{ProductId:int32(1),Quantity:int32(2),Id:int32(3)})
    products=append(products,&pb.Product{ProductId:int32(4),Quantity:int32(5),Id:int32(6)})
	
	
	result1:=pb.Cart{
        Product:products,
    }
    cartResponse:=pb.CartResponse{
        Cart:&result1,
    }
    return &(cartResponse),nil
}

func(*server) RewardService(ctx context.Context,req *pb.RewardRequest )(*pb.RewardResponse,error){
    fmt.Println("Reward Service called")
    id_s:=req.GetCustomerId().GetId()
	fmt.Println(id_s)

	//  User your database here to fetch products in the reward

	// fetch reward here of the customer whose id is id_s
	// convert golang data types into grpc types   

	//


	reward:=pb.Reward{CustomerId:"12345",Description:"Silver level points",Points:int32(32),Id:int32(32)};
 	
 
    rewardResponse:=pb.RewardResponse{
        Reward:&reward,
    }
    return &(rewardResponse),nil
}



func GrpcServer(){
	fmt.Println("Server started")
   // config.InitKafka()
    opts:=[]grpc.ServerOption{}
    s:=grpc.NewServer(opts...)
    pb.RegisterServiceServer(s,&server{})

    Listener,err:=net.Listen("tcp",":"+config.Server["GRPC_PORT"])
	if err!=nil{
		fmt.Println("Error while starting Grpc Server",err.Error())
	}
    s.Serve(Listener)


}