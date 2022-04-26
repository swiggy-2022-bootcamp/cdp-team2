package middlewares

import (
	"fmt"
	grpc_client "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/services/grpc_client_services"
	"net/http"
	_ "net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuthMiddleware(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) > 0 {
		access_token := strings.Split(token[0], "Bearer ")[1]
		fmt.Println(access_token)

		resp, err := grpc_client.VerifyToken(c, access_token)
		if err == nil {
			fmt.Println("res", resp)
			c.Next()

		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Token Invalid"})
			return
		}

	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Token Invalid"})
		return
	}

}
