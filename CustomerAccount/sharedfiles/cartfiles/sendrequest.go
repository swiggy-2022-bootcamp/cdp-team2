package grpc
import (
	model "customer-account/internal/dao"
	"context"
	pb "customer-account/internal/proto"
	grpc "google.golang.org/grpc"
	"fmt"
	"customer-account/internal/literals"
	"strconv"
)
func GetCartByCustomerId(customer_id string)model.Cart{
	conn,_:=grpc.Dial("localhost:"+strconv.Itoa(literals.CART_PORT),grpc.WithInsecure())
	defer conn.Close()
	c:=pb.NewServiceClient(conn)
	return CartService(c,customer_id);	
}

func CartService(c pb.ServiceClient,customer_id string)model.Cart{
	cartRequest:=pb.CartRequest{
			CustomerId:customer_id,
	}
	
	res,_:=c.CartService(context.Background(),&cartRequest)
	product:=[]model.Product{}
	for _,val:=range res.Cart.Product{
		product=append(product,model.Product{val.ProductId,int(val.Quantity)})
		fmt.Println(product)
	}
	cart:=model.Cart{product}
	return cart;
}
// func main(){
// 	GetCartByCustomerId("1234567890");
// }