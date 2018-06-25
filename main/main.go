package main

import (
	"log"
	"github.com/IvoryRaptor/mint/kernel"
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/mint/scheduler"
	"github.com/IvoryRaptor/mint/mq"
	"github.com/IvoryRaptor/mint/http"
)

func main() {
	s := kernel.Mint{}
	s.New("mint")
	s.Set("matrix","default")
	s.Set("angler", "mint")
	err := dragonfly.Builder(
		&s,
		[]dragonfly.IServiceFactory{
			&scheduler.Factory{},
			&dragonfly.ZookeeperFactory{},
			&mq.Factory{},
			&http.Factory{},
		})
	s.SetFields()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
	s.WaitStop()
}
