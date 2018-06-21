package main

import (
	"log"
	"github.com/IvoryRaptor/skin/kernel"
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/skin/scheduler"
	"github.com/IvoryRaptor/skin/mq"
)

func main() {
	s := kernel.Skin{}
	s.New("skin")
	err := dragonfly.Builder(
		&s,
		[]dragonfly.IServiceFactory{
			&scheduler.Factory{},
			&dragonfly.ZookeeperFactory{},
			&mq.Factory{},
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
