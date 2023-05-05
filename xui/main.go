package xui

import (
	"github.com/harleywinston/x-operators/xui/consts"
)

// func registerHandlers() error {
// 	r := gin.Default()
//
// 	setupHandlers := transport.SetupHandlers{}
// 	r.POST("/add", setupHandlers.AddClientHandler)
// 	r.DELETE("/delete", setupHandlers.DeleteClientHandler)
// 	r.GET("/list", setupHandlers.ListInboundsHandler)
//
// 	if err := r.Run(); err != nil {
// 		return err
// 	}
// 	return nil
// }

func InitApp() error {
	if err := consts.InitApiSession(); err != nil {
		return err
	}
	// if err := registerHandlers(); err != nil {
	// 	return err
	// }
	return nil
}
