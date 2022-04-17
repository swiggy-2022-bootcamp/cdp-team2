package authsrv

import (
	"fmt"
	"net/http"
	"strings"

	domain "github.com/auth-frontstore-service/internal/core/domain"
	usersrv "github.com/auth-frontstore-service/internal/core/services/usersrv"
	repo "github.com/auth-frontstore-service/internal/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Login(c *gin.Context) (int, gin.H, error) {

	body := LoginBody{}

	if err := c.BindJSON(&body); err != nil {
		return http.StatusBadRequest, nil, err
	}

	isCorrectCredentials, customerId := usersrv.New().VerifyCustomerCredentials(body.Username, body.Password)
	if !isCorrectCredentials {
		return http.StatusNotFound, gin.H{"message": "User Not found"}, nil
	}

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while getting User Id"}, nil
	}

	user := &domain.User{
		CustomerId: id,
		Role:       "Customer",
		ID:         primitive.NewObjectID(),
	}

	token, err := repo.JWTManager.GenerateBasicToken(user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while creating token"}, nil
	}

	return http.StatusOK, gin.H{"token": token}, nil
}

func OAuth(c *gin.Context) (int, gin.H, error) {
	userService := usersrv.New()

	id, err := primitive.ObjectIDFromHex(c.GetString("User"))
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while getting User Id"}, nil
	}
	user := &domain.User{
		CustomerId: id,
		Role:       "Customer",
	}
	//	fmt.Println("Id in OAuth", id)
	token, err := repo.JWTManager.Generate(user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while creating token"}, nil
	}

	if err := userService.UpdateUser(&bson.M{"customerId": user.CustomerId}, &bson.M{"$push": bson.M{"tokens": token}}, id); err != nil {
		fmt.Println(err)
		return http.StatusNotFound, gin.H{"message": "Unable To Push Token"}, nil
	}

	return http.StatusOK, gin.H{"token": token}, nil
}

func Logout(c *gin.Context) (int, gin.H, error) {
	userService := usersrv.New()
	token := strings.Split(c.Request.Header["Authorization"][0], "Bearer ")[1]
	id, err := primitive.ObjectIDFromHex(c.GetString("User"))

	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while getting User Id"}, nil
	}

	if err := userService.UpdateUser(&bson.M{"customerId": id}, &bson.M{"$pull": bson.M{"tokens": token}}, id); err != nil {
		return http.StatusNotFound, gin.H{"message": "Unable To Push Token"}, nil
	}

	return http.StatusOK, gin.H{"message": "Logout Successfull"}, nil
}
