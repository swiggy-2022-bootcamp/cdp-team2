package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/literals"
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
