package grpc
import (
	model "customer-account/dao/models"
	"context"
	protos "customer-account/internal/proto/reward"
	grpc "google.golang.org/grpc"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	// "customer-account/internal/literals"
	// "strconv"
)
func GetRewardByCustomerId(customer_id string)model.Reward{
	// conn,err:=grpc.Dial("localhost:"+strconv.Itoa(literals.REWARD_PORT),grpc.WithInsecure())
	conn, err := grpc.Dial("35.84.28.237:30219", grpc.WithTransportCredentials(insecure.NewCredentials()))
	reward:=model.Reward{}

	if err != nil {
		fmt.Println("unable to connect to grpc service", err)
		return reward
	}
	defer conn.Close()

	client := protos.NewRewardServiceClient(conn)

	req := &protos.RewardRequest{
		CustomerId: customer_id,
	}
	res, err := client.GetReward(context.Background(), req)
	if err != nil {
 		return reward
	}
reward=model.Reward{res.Reward.CustomerId,res.Reward.Description,int(res.Reward.Points)} 
fmt.Println(reward)
return reward;
return model.Reward{}
}

// func RewardService(c pb.ServiceClient,customer_id string)model.Reward{

// }
// func main(){
// 	GetRewardByCustomerId("1234567890");
// }