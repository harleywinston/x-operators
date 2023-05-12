package syncercmd

import (
	syncer "github.com/harleywinston/x-operators/pkg/app"
)

func SetupSyncer() error {
	return syncer.InitSyncer()
}
