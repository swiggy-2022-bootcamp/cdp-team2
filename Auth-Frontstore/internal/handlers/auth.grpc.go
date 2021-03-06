package handlers

import (
	"context"
	"fmt"

	"github.com/auth-frontstore-service/internal/core/domain"
	"github.com/auth-frontstore-service/internal/core/services/usersrv"
	repo "github.com/auth-frontstore-service/internal/repository"
	"github.com/auth-frontstore-service/protos/authpb"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	authpb.UnimplementedAuthServiceServer
}

func (server *GrpcServer) Verify(ctx context.Context, req *authpb.VerifyRequest) (*authpb.VerifyResponse, error) {
	token := req.GetToken()
	ok, err := repo.JWTManager.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot verify token: %v", err)
	} else {
		user := &domain.User{}
		id := ok.ID
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error while getting User Id")
		}
		fmt.Println("Grpc Call CustomerId , token", id, token)
		if err := usersrv.New().FetchUser(&bson.M{"customerId": id, "tokens": token}, user); err != nil {
			return nil, status.Errorf(codes.Internal, "Token Invalid")
		} else {
			res := &authpb.VerifyResponse{
				Proceed:    true,
				CustomerId: ok.ID,
			}
			return res, nil
		}
	}
}
