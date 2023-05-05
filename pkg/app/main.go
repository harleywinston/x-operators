package main

import (
	"time"

	"github.com/harleywinston/x-operators/pkg/services"
)

func registerSyncer() error {
	syncerServices := services.SyncerServices{}

	ticker := time.NewTicker(30 * time.Minute)

	for range ticker.C {
		err := syncerServices.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func InitApp() error {
	if err := registerSyncer(); err != nil {
		return err
	}
	return nil
}
