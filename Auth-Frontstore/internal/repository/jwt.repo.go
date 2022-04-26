package repository

import (
	"fmt"
	"time"

	config "github.com/auth-frontstore-service/config"
	domain "github.com/auth-frontstore-service/internal/core/domain"
	"github.com/golang-jwt/jwt"
)

type jwtManager struct {
	publicKey     string
	privateKey    string
	secretKey     string
	tokenDuration time.Duration
}

type userClaims struct {
	jwt.StandardClaims
	ID string `json:"_id"`
}

var JWTManager *jwtManager

func init() {
	JWTManager = &jwtManager{
		privateKey:    config.SecretKey,
		publicKey:     config.SecretKey,
		secretKey:     config.SecretKey,
		tokenDuration: time.Minute * 5,
	}
}

func (manager *jwtManager) GenerateBasicToken(user *domain.User) (string, error) {
	claims := userClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		ID: user.CustomerId,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(manager.secretKey))
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (manager *jwtManager) Generate(user *domain.User) (string, error) {
	claims := userClaims{
		ID: user.CustomerId,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(manager.privateKey))
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil

}

func (manager *jwtManager) VerifyBasicToken(accessToken string) (*userClaims, error) {

	token, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}
			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	//	fmt.Println("Verify Basic Token", claims.ID)
	return claims, nil
}

func (manager *jwtManager) Verify(accessToken string) (*userClaims, error) {

	// key, err := jwt.ParseRSAPublicKeyFromPEM(manager.publicKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("validate: parse key: %w", err)
	// }

	token, err := jwt.ParseWithClaims(
		accessToken,
		&userClaims{},
		func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}

			return []byte(manager.publicKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	//	fmt.Println("Claims", claims.ID)

	return claims, nil
}
