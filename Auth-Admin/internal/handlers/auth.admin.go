package handlers

import (
	"net/http"

	authsrv "github.com/auth-admin-service/internal/core/services/authsrv"
	"github.com/auth-admin-service/internal/literals"
	"github.com/gin-gonic/gin"
)

type HandlersImpl struct {
	// TODO: Need to inject product services
}

func NewHandlers() *HandlersImpl {
	return &HandlersImpl{}
}

func HandlerWrapper(handler func(*gin.Context) (int, gin.H, error), gctx *gin.Context) {
	code, msg, err := handler(gctx)
	if err != nil {
		gctx.AbortWithError(code, err)
		return
	} else {
		gctx.IndentedJSON(code, msg)
		return
	}
}

// Health godoc
// @Summary      Health Check Route
// @Description  API to check products-admin health
// @Tags         Health
// @Accept       json
// @Produce      plain/text
// @Success      200  {object} 	string
// @Router       / [get]
func (h *HandlersImpl) Health(gctx *gin.Context) {
	contentType := "text/plain; charset=utf-8"
	gctx.Data(http.StatusOK, contentType, []byte(literals.HEALTH_MESSAGE))
}

// Login godoc
// @Summary      Login Route
// @Description  API to Login Users
// @Tags         Login
// @Accept       json
// @Produce      plain/text
// @Success      200  {object} 	string
// @Router       /login [post]
// @Param username body string true "Username/email of the user"
// @Param password body string true "Password of the user"
func (h *HandlersImpl) Login(gctx *gin.Context) {
	HandlerWrapper(authsrv.Login, gctx)
}

// OAuth godoc
// @Summary      OAuth Access Token
// @Description  API to Generate access token using basic token provided by login API
// @Tags         OAuth
// @Accept       json
// @Produce      plain/text
// @Success      200  {object} 	string
// @Router       /oAuth/token [post]
// @Param Authorization header string true "Insert your basic token" default(Basic <Add access token here>)
func (h *HandlersImpl) OAuth(gctx *gin.Context) {
	HandlerWrapper(authsrv.OAuth, gctx)
}

// Logout godoc
// @Summary      Logout Route
// @Description  API to delete acess token from user databse and logout user
// @Tags         Logout
// @Accept       json
// @Produce      plain/text
// @Success      200  {object} 	string
// @Router       /logout [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (h *HandlersImpl) Logout(gctx *gin.Context) {
	HandlerWrapper(authsrv.Logout, gctx)
}
