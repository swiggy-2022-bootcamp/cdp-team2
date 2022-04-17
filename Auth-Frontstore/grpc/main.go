package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/auth-frontstore-service/internal/handlers"
	"github.com/auth-frontstore-service/protos/authpb"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &handlers.GrpcServer{
		authpb.UnimplementedAuthServiceServer{},
	})

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Block Untill Signal is Received
	<-ch

	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of program")
}
