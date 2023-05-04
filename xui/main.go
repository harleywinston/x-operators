package xui

import (
	"github.com/gin-gonic/gin"

	"github.com/harleywinston/x-operators/xui/internal/transport"
)

func registerHandlers() error {
	r := gin.Default()

	setupHandlers := transport.SetupHandlers{}
	r.GET("/add", setupHandlers.AddClientHandler)
	r.GET("/delete", setupHandlers.DeleteClientHandler)
	return nil
}

func InitApp() error {
	if err := registerHandlers(); err != nil {
		return err
	}
	return nil
}
