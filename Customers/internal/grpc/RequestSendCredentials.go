package grpc

import (
	"context"
 	pb "customers/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
	// "customers/internal/literals"
	// "strconv"

)

func SendCredential(username string,password string) (bool ,string){
	conn, _ := grpc.Dial("localhost:9000", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	fmt.Println("1")
	return SendCredentialService(c, username,password)
}

func SendCredentialService(c pb.ServiceClient, username string,password string) (bool,string) {
	credentialRequest := pb.CredentialRequest{
		Credential: &pb.Credential{
			Username:username,
			Password:password,
		},
	}
	fmt.Println(2)
	res,err:=c.CredentialService(context.Background(), &credentialRequest)
	if(err!=nil){
		fmt.Println("Encountered an error while creating the",err.Error())
		return false,err.Error()
	}
	return res.Ispresent,res.Customerid
}


func main(){
	fmt.Println(SendCredential("uday","kiranbakka"))
}
 