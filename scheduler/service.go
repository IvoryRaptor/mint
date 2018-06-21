package scheduler

import (
	"github.com/IvoryRaptor/skin"
	"github.com/IvoryRaptor/dragonfly"
	"time"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/IvoryRaptor/postoffice"
)

type Service struct {
	skin   skin.ISkin
	ch     chan int
	run    bool
	second int
}

func (s *Service) Config(kernel dragonfly.IKernel, config map[interface{}]interface{}) error {
	s.skin = kernel.(skin.ISkin)
	s.ch = make(chan int)
	s.run = false
	s.second = config["second"].(int)
	return nil
}

func (s *Service) Start() error {
	s.run = true
	go func() {
		for s.run {
			s.ch <- 1
			time.Sleep(time.Duration(s.second) * time.Second)
		}
	}()
	actor := make([]byte, 0)
	go func() {
		for {
			d := <-s.ch
			switch d {
			case 1:
				mes := postoffice.MQMessage{
					Source: &postoffice.Address{
						Matrix: "default",
						Device: "skin",
					},
					Destination: &postoffice.Address{
						Matrix: "",
						Device: "",
					},
					Resource: "skin",
					Action:   "heart",
					Payload:  make([]byte, 0),
				}
				payload, _ := proto.Marshal(&mes)
				for _, topic := range s.skin.GetTopics() {
					log.Printf("Publish %s", topic)
					s.skin.Publish(topic, actor, payload)
				}
			case -1:
				goto END
			default:

			}
		}
	END:
	}()
	return nil
}

func (s *Service) Stop() {
	s.run = false
	s.ch <- -1
	close(s.ch)
	s.skin.RemoveService(s)
}
