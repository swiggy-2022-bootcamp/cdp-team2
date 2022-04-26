package controllers

import (
	authpb "common/protos/authpb"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/api"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/grpc"
)

type AuthController struct {
	AuthClient *grpc.AuthClient
}

func NewAuthController() (AuthController, error) {
	ac, err := grpc.NewAuthClient(context.Background())
	return AuthController{ac}, err
}

func (ac *AuthController) CheckAuth(c *gin.Context) {
	token := c.Request.Header["Token"][0]

	res, err := ac.AuthClient.Verify(ac.AuthClient.CtxWithTimeOut(), &authpb.VerifyRequest{
		Token: token,
	})

	if (err != nil) || !(res.GetProceed()) {
		log.Printf("[Error] invalid token")
		c.AbortWithStatusJSON(api.UserErr(fmt.Errorf("Invalid token")))
		return
	}

	log.Printf("Response from Auth: Proceed : %v , CustomerId : %v", res.GetProceed(), res.GetCustomerId())

	c.Set("CustomerId", res.GetCustomerId())
	c.Next()
}
