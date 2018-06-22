package main

import (
	"log"
	"github.com/IvoryRaptor/mint/kernel"
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/mint/scheduler"
	"github.com/IvoryRaptor/mint/mq"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	s := kernel.Mint{}
	s.New("skin")
	s.Set("matrix","default")
	s.Set("angler", "mint")
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
	s.Zookeeper.Create("/iotnn/default/mint._heart/mint", make([]byte, 0), 0, zk.WorldACL(zk.PermAll))
	s.WaitStop()
}
