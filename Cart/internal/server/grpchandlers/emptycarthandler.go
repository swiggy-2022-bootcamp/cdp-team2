package grpchandlers

import (
	"context"
	"errors"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/cart/protos"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type Server struct{}

func (s *Server) EmptyCart(ctx context.Context, cartReq *pb.CartRequest) (*pb.EmptyCartResponse, error) {
	dao := dao.GetDynamoDAO()

	err := dao.EmptyCart(cartReq.CustomerId)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err.ErrorMessage,
		}).Error("an error occurred while emptying the cart for customer: ", cartReq.CustomerId)
		return &pb.EmptyCartResponse{}, errors.New("an error occurred while emptying the cart")
	}

	return &pb.EmptyCartResponse{}, nil
}
