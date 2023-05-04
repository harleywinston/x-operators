package xui

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-operators/xui/configs"
	"github.com/harleywinston/x-operators/xui/internal/transport"
)

func registerHandlers() error {
	r := gin.Default()

	setupHandlers := transport.SetupHandlers{}
	r.POST("/add", setupHandlers.AddClientHandler)
	r.DELETE("/delete", setupHandlers.DeleteClientHandler)

	if err := r.Run(); err != nil {
		return err
	}
	return nil
}

func InitApp() error {
	if err := configs.InitApiSession(); err != nil {
		return err
	}
	if err := registerHandlers(); err != nil {
		return err
	}
	return nil
}
