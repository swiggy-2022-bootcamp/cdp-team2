package responses

import "github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/domain"

type MessageResponse struct {
	Message string `json:"message"`
}

type ProductMessage struct {
	Product *domain.Product `json:"product"`
}
type ProductsMessage struct {
	Products []*domain.Product `json:"products"`
}
