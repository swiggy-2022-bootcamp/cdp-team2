package services

import (
	pb "common/pkg/protos"
	"context"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/pkg/errors"
)

type ProductsServices struct {
	Ctx                context.Context
	ProductsGRPCClient pb.ProductsServicesClient
}

var _ ports.IProductsServices = (*ProductsServices)(nil)

func NewProductsServices(ctx context.Context, productsGrpcClient pb.ProductsServicesClient) *ProductsServices {
	return &ProductsServices{
		Ctx:                ctx,
		ProductsGRPCClient: productsGrpcClient,
	}
}

func (ps *ProductsServices) GetProductList() ([]*domain.Product, *errors.AppError) {
	res, err := ps.ProductsGRPCClient.GetAvailableProducts(ps.Ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, errors.New(err.Error(), http.StatusInternalServerError)
	}
	var _products []*domain.Product
	for _, _gPRCProduct := range res.Products {
		var _product domain.Product
		_product.BindGprcProduct(_gPRCProduct)
		_products = append(_products, &_product)
	}
	return _products, nil
}

func (ps *ProductsServices) GetProductById(productID int64) (*domain.Product, *errors.AppError) {
	res, err := ps.ProductsGRPCClient.GetProductById(ps.Ctx, &pb.ProductIDRequest{
		ProductID: productID,
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}
	var _product *domain.Product
	_product.BindGprcProduct(res.Product)
	return _product, nil
}
