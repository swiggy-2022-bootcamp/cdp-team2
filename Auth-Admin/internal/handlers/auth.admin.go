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

func (h *HandlersImpl) Login(gctx *gin.Context) {
	code, msg, err := authsrv.Login(gctx)
	if err != nil {
		gctx.AbortWithError(code, err)
		return
	} else {
		gctx.IndentedJSON(code, msg)
		return
	}
}
