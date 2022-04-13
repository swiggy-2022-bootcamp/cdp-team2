package authsrv

import (
	"context"
	"fmt"
	"net/http"
	"time"

	domain "github.com/auth-admin-service/internal/core/domain"
	usersrv "github.com/auth-admin-service/internal/core/services/usersrv"
	repo "github.com/auth-admin-service/internal/repository"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body := LoginBody{}

	if err := c.BindJSON(&body); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user := &domain.User{}
	userService := usersrv.New()
	db := repo.ConnectDB().DataStore
	if err := db.Collection("user").FindOne(ctx, bson.M{"username": body.Username}).Decode(&user); err != nil {
		return http.StatusNotFound, gin.H{"message": "User Not found"}, nil
	}

	if user == nil || !userService.IsCorrectPassword(user, body.Password) {
		return http.StatusNotFound, gin.H{"message": "Password Incorrect"}, nil
	}

	token, err := repo.JWTManager.Generate(user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while creating token"}, nil
	}

	return http.StatusOK, gin.H{"user": user, "token": token}, nil
}

func Signup(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body := SignUpBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := repo.ConnectDB().DataStore
	userService := usersrv.New()
	user, err := userService.NewUser(body.Username, body.Password, body.Role)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res, err := db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("can not convert to oid %v", err)})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"username": body.Username, "Id": oid.Hex()})
}

func CheckAuth(c *gin.Context) {
	user := c.GetString("User")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Auth Working", "user": user})
}
