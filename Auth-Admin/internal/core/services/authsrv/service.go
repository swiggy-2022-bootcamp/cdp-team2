package authsrv

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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

	body := LoginBody{}

	if err := c.BindJSON(&body); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user := &domain.User{}
	userService := usersrv.New()
	if err := userService.FetchUser(&bson.M{"username": body.Username}, user); err != nil {
		return http.StatusNotFound, gin.H{"message": "User Not found"}, nil
	}

	if user == nil || !userService.IsCorrectPassword(user, body.Password) {
		return http.StatusNotFound, gin.H{"message": "Password Incorrect"}, nil
	}

	token, err := repo.JWTManager.GenerateBasicToken(user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while creating token"}, nil
	}

	return http.StatusOK, gin.H{"token": token}, nil
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

func OAuth(c *gin.Context) (int, gin.H, error) {
	userService := usersrv.New()
	user := &domain.User{}
	id, err := primitive.ObjectIDFromHex(c.GetString("User"))
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while getting User Id"}, nil
	}
	if err := userService.FetchUser(&bson.M{"_id": id}, user); err != nil {
		return http.StatusNotFound, gin.H{"message": "User Not found"}, nil
	}

	token, err := repo.JWTManager.Generate(user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Error while creating token"}, nil
	}

	if err := userService.UpdateUser(&bson.M{"_id": user.ID}, &bson.M{"$push": bson.M{"tokens": token}}); err != nil {
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

	if err := userService.UpdateUser(&bson.M{"_id": id}, &bson.M{"$pull": bson.M{"tokens": token}}); err != nil {
		return http.StatusNotFound, gin.H{"message": "Unable To Push Token"}, nil
	}

	return http.StatusOK, gin.H{"message": "Logout Successfull"}, nil
}
