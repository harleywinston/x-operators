package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-operators/xui/internal/service"
)

type SetupHandlers struct {
	service service.SetupServices
}

func (h *SetupHandlers) AddClientHandler(ctx *gin.Context) {
	// return h.service.AddClientService()
}

func (h *SetupHandlers) DeleteClientHandler(ctx *gin.Context) {
	// return h.service.DeleteClientService()
}
