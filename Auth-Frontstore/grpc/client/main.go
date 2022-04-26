package main

import (
	"context"
	"fmt"
	"log"

	"github.com/auth-frontstore-service/protos/authpb"
	"google.golang.org/grpc"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI3TGVENUp0YUR6YUVuSzhDeEhociJ9.pyxQmVf-L6tq5h_1-YtCqtEMyRshaAClMZiqx_gCcMM"
)

func verifyToken(c authpb.AuthServiceClient) {
	res, err := c.Verify(context.Background(), &authpb.VerifyRequest{
		Token: token,
	})
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Response from Auth: Proceed : %v , CustomerId : %v", res.GetProceed(), res.GetCustomerId())
}

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial(
		"35.84.28.237:30220",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	c := authpb.NewAuthServiceClient(cc)

	verifyToken(c)
}
