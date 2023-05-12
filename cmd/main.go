package main

import (
	"log"

	syncercmd "github.com/harleywinston/x-operators/cmd/syncer"
)

func main() {
	log.Fatal(syncercmd.SetupSyncer())
}
