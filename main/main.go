package main

import (
	"log"
	"github.com/IvoryRaptor/skin/skin"
)

func main() {
	s := skin.Skin{
		ConfigFile: "./config/postoffice/config.yaml",
	}
	err := k.Config()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = k.Start()
	if err != nil {
		log.Fatalf(err.Error())
	}
	k.WaitStop()
}
