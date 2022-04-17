package grpc_services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/protos/shipping"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
)

type Server struct {
	*pb.UnimplementedServiceServer
}

func (s *Server) AddressutsService(ctx context.Context, address *pb.AddressutsRequest) (*pb.AddressutsResponse, error) {

	customerIdStr, firstName, lastName, addressLine1, addressLine2, city, countryId, postCodeStr := address.Address.CustomerId, address.Address.Firstname, address.Address.Lastname, address.Address.Address1, address.Address.Address2, address.Address.City, address.Address.CountryId, address.Address.Postcode

	fmt.Println(customerIdStr, firstName, lastName, addressLine1, addressLine2, city, countryId, postCodeStr)

	customerId, _ := strconv.Atoi(customerIdStr)

	postCode, _ := strconv.Atoi(postCodeStr)

	shippingAddress := models.ShippingAddress{
		CustomerId:   customerId,
		FirstName:    firstName,
		LastName:     lastName,
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
		City:         city,
		CountryCode:  countryId,
		PostCode:     postCode,
	}

	_, err := dao.GetOrderDao().Create(shippingAddress)

	if err != nil {
		return nil, errors.New("error occurred")
	}

	return &pb.AddressutsResponse{
		Response: true,
	}, nil
}
