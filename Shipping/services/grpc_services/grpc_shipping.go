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

	customerIdStr, firstName, lastName, addressLine1, addressLine2, city, countryId, postCode := address.Address.CustomerId, address.Address.Firstname, address.Address.Lastname, address.Address.AddressLine1, address.Address.AddressLine2, address.Address.City, address.Address.CountryCode, address.Address.PostCode

	fmt.Println(customerIdStr, firstName, lastName, addressLine1, addressLine2, city, countryId, postCode)

	customerId, _ := strconv.Atoi(customerIdStr)

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

func (s *Server) AddressstuService(ctx context.Context, req *pb.AddressstuRequest) (*pb.AddressstuResponse, error) {
	customerIdStr := req.CustomerId
	fmt.Println(customerIdStr)

	customerId, _ := strconv.Atoi(customerIdStr)

	addressResponse, err := dao.GetOrderDao().GetByCustomerId(customerId)

	if err != nil {
		return nil, errors.New("error occurred")
	}

	addresses := []*pb.Address{}

	for i := 0; i < len(addressResponse); i++ {
		addresses = append(addresses, &pb.Address{CustomerId: string(addressResponse[i].CustomerId),
			Firstname:    addressResponse[i].FirstName,
			Lastname:     addressResponse[i].LastName,
			AddressLine1: addressResponse[i].AddressLine1,
			AddressLine2: addressResponse[i].AddressLine2,
			City:         addressResponse[i].City,
			PostCode:     addressResponse[i].PostCode,
			CountryCode:  addressResponse[i].CountryCode})

	}

	return &(pb.AddressstuResponse{Address: addresses}), nil
}
