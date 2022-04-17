package server

import (
 	pb "customers/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
	"net"
	"context"
 	// "strconv"
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

func (*server)  AddressutsService(ctx context.Context,req *pb.AddressutsRequest)(*pb.AddressutsResponse,error){
	customer_id:=req.Address.CustomerId
	firstname:=req.Address.Firstname
	lastname:=req.Address.Lastname
	address1:=req.Address.Address1
	address2:=req.Address.Address2
	city:=req.Address.City
	postcode:=req.Address.Postcode
	country_id:=req.Address.CountryId

	fmt.Println(customer_id,firstname,lastname,address1,address2,city,postcode,country_id);
	return &(pb.AddressutsResponse{Response:true}),nil
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

    Listener,_:=net.Listen("tcp",":9000")
    s.Serve(Listener)

}