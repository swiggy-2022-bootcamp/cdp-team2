package grpc
import (
	model "customer-account/dao/models"
	"customer-account/config"
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
	conn, err := grpc.Dial("localhost:"+config.Server["REWARD_PORT"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	reward:=model.Reward{}

	if err != nil {
		fmt.Println("unable to connect to grpc service", err)
		return reward
	}
	defer conn.Close()

	client := protos.NewRewardServiceClient(conn)

	req := &protos.RewardRequest{
		CustomerId: "123",
	}
	res, err := client.GetReward(context.Background(), req)
	if err != nil {
		fmt.Println("Error1:",res,err)
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