package grpc
import (
	model "customer-account/dao/models"
	// "context"
	// pb "customer-account/internal/proto/cart"
	// grpc "google.golang.org/grpc"
	// "fmt"
	// "google.golang.org/grpc/credentials/insecure"
	// "customer-account/internal/literals"
	// "strconv"
)
func GetCartByCustomerId(customer_id string)model.Cart{
	// conn,err:=grpc.Dial(":9000",grpc.WithInsecure())
	// // conn,err:=grpc.Dial("0.tcp.in.ngrok.io:11480",grpc.WithTransportCredentials(insecure.NewCredentials()))
	// fmt.Println("entered",err)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// defer conn.Close()
	// c:=pb.NewCartServiceClient(conn)

	// cartRequest:=pb.CartRequest{
	// 		CustomerId:"133",
	// }
	
	// res,err:=c.GetCart(context.Background(),&cartRequest)
	// fmt.Println("res",res,err)
	// product:=[]model.Product{}
	// for _,val:=range res.Cart.Products{
	// 	product=append(product,model.Product{val.ProductId,int(val.Quantity)})
	// 	fmt.Println(product)
	// }
	// cart:=model.Cart{product}
	// fmt.Println(cart)
	// return cart;
	return model.Cart{}
}
