package scheduler

import (
	"github.com/IvoryRaptor/mint"
	"github.com/IvoryRaptor/dragonfly"
	"time"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/IvoryRaptor/postoffice"
	"strings"
)

type Service struct {
	mint      mint.IMint
	ch        chan int
	run       bool
	second    int
	partition int
}

func (s *Service) Config(kernel dragonfly.IKernel, config map[interface{}]interface{}) error {
	s.mint = kernel.(mint.IMint)
	s.ch = make(chan int)
	s.run = false
	s.second = config["second"].(int)
	s.partition = config["partition"].(int)
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
	caller := []byte("default_mint")
	actor := make([]byte, 0)
	go func() {
		for {
			d := <-s.ch
			switch d {
			case 1:
				for _, topic := range s.mint.GetTopics() {
					sp := strings.Split(topic, "_")
					for i := 0; i < s.partition; i++ {
						mes := postoffice.MQMessage{
							Caller:   caller,
							Matrix:   sp[0],
							Device:   sp[1],
							Time:     time.Now().Unix(),
							Resource: "mint",
							Action:   "heart",
							Payload:  []byte{byte(i)},
						}
						payload, _ := proto.Marshal(&mes)
						log.Printf("Publish %s", topic)
						s.mint.Publish(topic, int32(i), actor, payload)
					}
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
	s.mint.RemoveService(s)
}
