package main

import (
	"context"
	"fmt"
	"log"

	"github.com/auth-frontstore-service/protos/authpb"
	"google.golang.org/grpc"
)

const (
	token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MjNjYjlmNTdjYmJlOTliZTMxOTMzNDEifQ.MFfCbN2yo0TgX23b1D5EuTDUoxxX-8SIwiY3f9mOScF_wBkErJUBn217_FpxYCCwbYPzRanPnvfg8k4TPwmE74vJIxztCBxLyUxtp8RTvU482bhM_jZ3NNyTdPAMrxn2U1vKkSC4PkSTfrcXua6i7aqk5-Nq5kb0Zt0Fz4TeoBo"
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
		"localhost:50051",
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
