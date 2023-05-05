package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-operators/xui/consts"
	"github.com/harleywinston/x-operators/xui/internal/models"
	"github.com/harleywinston/x-operators/xui/internal/service"
)

type SetupHandlers struct {
	service service.SetupServices
}

func (h *SetupHandlers) AddClientHandler(ctx *gin.Context) {
	var user models.UserModel
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.service.AddClientService(user)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.JSON(consts.ADD_SUCCESS.Code, gin.H{
		"message": consts.ADD_SUCCESS.Message,
		"detail":  fmt.Sprintf(`User email: %s`, user.Email),
	})
}

func (h *SetupHandlers) DeleteClientHandler(ctx *gin.Context) {
	var user models.UserModel
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(consts.BIND_JSON_ERROR.Code, gin.H{
			"message": consts.BIND_JSON_ERROR.Message,
			"detail":  err.Error(),
		})
		return
	}

	err = h.service.DeleteClientService(user)
	if err != nil {
		if e, ok := err.(*consts.CustomError); ok {
			ctx.JSON(e.Code, gin.H{
				"message": e.Message,
				"detail":  e.Detail,
			})
		}
		return
	}

	ctx.JSON(consts.DELETE_SUCCESS.Code, gin.H{
		"message": consts.DELETE_SUCCESS.Message,
		"detail":  fmt.Sprintf(`User email: %s`, user.Email),
	})
}
