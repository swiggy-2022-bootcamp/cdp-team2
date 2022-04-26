package grpchandlers

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/reward/protos"
	orderpb "github.com/swiggy-2022-bootcamp/cdp-team2/reward/protos/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) RedeemReward(ctx context.Context, reward *pb.RedeemRewardRequest) (*pb.RedeemRewardResponse, error) {
	orderReqpb := &orderpb.GetOrderRequest{
		OrderId: reward.OrderId,
	}

	// call to Order MS to get list of products
	conn, grpcErr := grpc.Dial("35.84.28.237:30208", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if grpcErr != nil {
		log.WithError(grpcErr).Error("unable to establish grpc connection")
		return nil, errors.New("an internal error occurred")
	}

	defer conn.Close()
	client := orderpb.NewServiceClient(conn)

	orderResppb, grpcErr := client.GetOrderService(context.Background(), orderReqpb)
	if grpcErr != nil {
		log.WithError(grpcErr).Error("an error occurred while getting the proto response from order ms")
		return nil, grpcErr
	}

	totalRewardInOrder := 0
	for _, product := range orderResppb.Order.ProductDesc {
		totalRewardInOrder += (int(product.Quantity) * int(product.Reward))
	}

	dao := dao.GetDynamoDAO()
	rewardModel, err := dao.GetReward(orderResppb.Order.CustomerId)
	if err != nil {
		log.Error("Error while getting the reward: ", err.ErrorMessage)
		return nil, errors.New("error occurred")
	}

	currentReward := rewardModel.Points
	totalAvailableReward := currentReward + int32(totalRewardInOrder)
	totalPrice := orderResppb.Order.TotalPrice
	rewardApplied := reward.Points
	finalPrice := totalPrice

	// applied reward by user is greater than the total price
	if rewardApplied > int32(totalPrice) {
		log.Error("reward applied is greater than the total price of the given order: ", orderResppb.Order.OrderId, " for user ", orderResppb.Order.CustomerId)
		return nil, errors.New("reward applied is greater than the total price of the given order")
	}

	// sufficient reward available for user
	if rewardApplied <= totalAvailableReward {
		// update the final price
		finalPrice = totalPrice - float32(rewardApplied)
		totalAvailableReward -= rewardApplied

		log.Info("updating reward points for the user ", orderResppb.Order.CustomerId)
		// update reward table with updated reward points
		dao.AddReward(models.Reward{
			CustomerId:  orderResppb.Order.CustomerId,
			Description: "Updated reward after applying reward points",
			Points:      totalAvailableReward,
		})
	} else {
		log.Error("reward applied is greater than the reward balance for user ", orderResppb.Order.CustomerId)
		return nil, errors.New("reward applied is greater than reward balance")
	}

	// call to Order MS to update the finalPrice of the order
	updateFinalPriceOrderReq := &orderpb.OrderPayedPriceUpdateRequest{
		OrderId:    reward.OrderId,
		PayedPrice: float32(finalPrice),
	}

	_, grpcErr = client.OrderPayedPriceUpdateService(context.Background(), updateFinalPriceOrderReq)
	if grpcErr != nil {
		log.WithError(grpcErr).Error("an error occurred while getting the proto response from order ms")
		return nil, grpcErr
	}
	redeemRewardResp := &pb.RedeemRewardResponse{
		FinalPrice: int32(finalPrice),
	}

	// return the final response to calling microservice
	return redeemRewardResp, nil
}
