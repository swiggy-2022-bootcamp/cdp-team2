package grpc

import (
	"context"
 	pb "customer-account/internal/proto"
	"fmt"
	grpc "google.golang.org/grpc"
)

func SendCredential(username string,password string) bool {
	conn, _ := grpc.Dial("localhost:8082", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewServiceClient(conn)
	fmt.Println("1")
	return SendCredentialService(c, username,password)
}

func SendCredentialService(c pb.ServiceClient, username string,password string) bool {
	credentialRequest := pb.CredentialRequest{
		Credential: &pb.Credential{
			Username:username,
			Password:password,
		},
	}
	fmt.Println(2)
	res,err:=c.CredentialService(context.Background(), &credentialRequest)
	if(err!=nil){
		fmt.Println("Encountered an error while creating the")
		return false
	}
	return res.Ispresent
}

 