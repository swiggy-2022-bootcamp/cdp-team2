package grpchandlers

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/cart/protos"
)

func (s *Server) GetCart(ctx context.Context, cartReq *pb.CartRequest) (*pb.CartResponse, error) {
	dao := dao.GetDynamoDAO()

	customerId := cartReq.CustomerId
	cartModel, err := dao.GetCart(customerId)
	if err != nil {
		log.Error("Error while getting the cart: ", err.ErrorMessage)
		return nil, errors.New("error occurred")
	}

	pbCartResponse := &pb.CartResponse{Cart: &pb.Cart{
		CustomerId: customerId,
	}}

	productsPb := []*pb.Product{}
	for _, prod := range cartModel.Products {
		productPb := &pb.Product{
			ProductId: prod.ProductId,
			Quantity:  prod.Quantity,
		}
		productsPb = append(productsPb, productPb)
	}

	pbCartResponse.Cart.Products = productsPb

	return pbCartResponse, nil
}
