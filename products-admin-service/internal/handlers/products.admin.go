package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlersImpl struct {
	// TODO: Need to inject product services
}

func NewHandlers() *HandlersImpl {
	return &HandlersImpl{}
}

func (h *HandlersImpl) Health(gctx *gin.Context) {
	contentType := "text/plain; charset=utf-8"
	gctx.Data(http.StatusOK, contentType, []byte("ok"))
}
