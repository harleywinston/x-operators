package main

import (
	"log"

	syncercmd "github.com/harleywinston/x-operators/cmd/syncer"
	xuicmd "github.com/harleywinston/x-operators/cmd/xui"
)

func main() {
	go log.Fatal(xuicmd.Setupxui())
	go log.Fatal(syncercmd.SetupSyncer())
}
