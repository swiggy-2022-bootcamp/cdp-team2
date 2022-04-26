package util

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/config"
)

type RouterConfig struct {
	WebServerConfig *config.WebServerConfig
}

func ExtractDetailsFromToken(req *http.Request) string {
	token, _ := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, keyLookupFunc)

	// extract user id from the JWT token
	claims := token.Claims.(jwt.MapClaims)
	userInfo := claims["CustomUserInfo"].(map[string]interface{})

	customerId := userInfo["_id"].(string)
	log.Info(customerId)

	return customerId
}

// keyLookupFunc returns the public key for JWT authentication
func keyLookupFunc(*jwt.Token) (interface{}, error) {
	return VerifyKey, nil
}
