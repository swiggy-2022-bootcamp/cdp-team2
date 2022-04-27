package grpc

import (
	"context"
 	pb "customer-account/internal/proto/grpc"
	"fmt"
	grpc "google.golang.org/grpc"
	"customer-account/config"
)

func SendCredential(username string,password string) bool {
	conn, _ := grpc.Dial("localhost:"+config.Server["GRPC_PORT"], grpc.WithInsecure())
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
		fmt.Println("Encountered an error while creating the",err)
		return false
	}
	return res.Ispresent
}

func main(){
	SendCredential("uday","kiran")
}