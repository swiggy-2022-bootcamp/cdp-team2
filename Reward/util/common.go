package util

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/config"
)

type RouterConfig struct {
	WebServerConfig *config.WebServerConfig
}

// func ExtractDetailsFromToken(req *http.Request) string {
// 	tokenStr := req.Header.Get("Authorization")
// 	// extract user id from the JWT token
// 	claims := &jwt.MapClaims{}

// 	tok, _, err := jwt.Parse()
// 	if err != nil {
// 		log.WithError(err).Error("an error occurred while parsing the token")
// 	}

// 	claims := token.Claims.(*jwt.MapClaims)
// 	//userInfo := claims["CustomUserInfo"].(map[string]interface{})
// 	log.Info(claims)
// 	customerId := "abc"
// 	//customerId := userInfo["_id"].(string)
// 	return customerId
// }
