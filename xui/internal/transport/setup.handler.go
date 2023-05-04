package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-operators/xui/internal/models"
	"github.com/harleywinston/x-operators/xui/internal/service"
)

type SetupHandlers struct {
	service service.SetupServices
}

func (h *SetupHandlers) AddClientHandler(ctx *gin.Context) {
	var user models.UserModel
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Not valid user",
			"err":     err,
		})
		return
	}

	if err := h.service.AddClientService(user); err != nil {
		ctx.JSON(500, gin.H{
			"message": "internal error",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("user %v added", user.Email),
	})
}

func (h *SetupHandlers) DeleteClientHandler(ctx *gin.Context) {
	var user models.UserModel
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Not valid user",
			"err":     err,
		})
		return
	}

	if err := h.service.DeleteClientService(user); err != nil {
		ctx.JSON(500, gin.H{
			"message": "internal error",
			"err":     err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("user %v deleted", user.Email),
	})
}
