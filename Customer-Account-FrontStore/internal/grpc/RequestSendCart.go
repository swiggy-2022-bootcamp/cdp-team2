package grpc

import (
	"context"
	model "customer-account/dao/models"
	pb "customer-account/internal/proto/cart"
	"fmt"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "customer-account/internal/literals"
	// "strconv"
)

func GetCartByCustomerId(customer_id string) model.Cart {
	conn, err := grpc.Dial("35.84.28.237:30217", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn,err:=grpc.Dial("0.tcp.in.ngrok.io:11480",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return model.Cart{}
	}
	defer conn.Close()
	c := pb.NewCartServiceClient(conn)

	cartRequest := pb.CartRequest{
		CustomerId: customer_id,
	}

	res, err := c.GetCart(context.Background(), &cartRequest)
	if err != nil {
		return model.Cart{}
	}
	product := []model.Product{}
	for _, val := range res.Cart.Products {
		product = append(product, model.Product{val.ProductId, int(val.Quantity)})
		fmt.Println(product)
	}
	cart := model.Cart{product}
	fmt.Println(cart)
	return cart
	return model.Cart{}
}

// func main(){
// 	GetCartByCustomerId("24567")
// }
