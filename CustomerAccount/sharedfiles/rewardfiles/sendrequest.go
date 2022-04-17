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
func GetRewardByCustomerId(customer_id string)model.Reward{
	conn,_:=grpc.Dial("localhost:"+strconv.Itoa(literals.REWARD_PORT),grpc.WithInsecure())
	defer conn.Close()
	c:=pb.NewServiceClient(conn)
	return RewardService(c,customer_id);	
}

func RewardService(c pb.ServiceClient,customer_id string)model.Reward{
	rewardRequest:=pb.RewardRequest{
			CustomerId:customer_id,
	}
	fmt.Println("*************")
	res,_:=c.RewardService(context.Background(),&rewardRequest)
	fmt.Println("*************",res)

	reward:=model.Reward{res.Reward.CustomerId,res.Reward.Description,int(res.Reward.Points)} 
	fmt.Println(reward)
	return reward;
}
// func main(){
// 	GetRewardByCustomerId("1234567890");
// }