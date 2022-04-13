package repository

import (
	"fmt"
	"io/ioutil"
	"log"

	domain "github.com/auth-admin-service/internal/core/domain"

	"github.com/golang-jwt/jwt"
)

type jwtManager struct {
	publicKey  []byte
	privateKey []byte
}

type userClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
	ID       string `json:"_id"`
}

var JWTManager *jwtManager

func init() {
	prvKey, err := ioutil.ReadFile("keys/private-key")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := ioutil.ReadFile("keys/public-key")
	if err != nil {
		log.Fatalln(err)
	}
	JWTManager = &jwtManager{
		privateKey: prvKey,
		publicKey:  pubKey,
	}
}

func (manager *jwtManager) Generate(user *domain.User) (string, error) {
	claims := userClaims{
		Username: user.Username,
		Role:     user.Role,
		ID:       user.ID.Hex(),
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(manager.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil

}

func (manager *jwtManager) Verify(accessToken string) (*userClaims, error) {

	key, err := jwt.ParseRSAPublicKeyFromPEM(manager.publicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	token, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}

			return key, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
