package syncer

import (
	"log"
	"time"

	"github.com/harleywinston/x-operators/pkg/services"
	"github.com/harleywinston/x-operators/xui"
)

func registerSyncer() {
	syncerServices := services.SyncerServices{}

	ticker := time.NewTicker(2 * time.Second)

	// go func() {
	for range ticker.C {
		err := syncerServices.Sync()
		if err != nil {
			log.Println(err.Error())
		}
	}
	// }()

	// select {}
}

func InitSyncer() error {
	if err := xui.InitDriver(); err != nil {
		return err
	}
	// if err := helper.InitAPISession(); err != nil {
	// 	return err
	// }
	registerSyncer()

	return nil
}
