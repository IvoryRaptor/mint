package scheduler

import (
	"github.com/IvoryRaptor/skin"
	"github.com/IvoryRaptor/dragonfly"
	"time"
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
	go func() {
		for {
			d := <-s.ch
			switch d {
			case 1:
				for _,topic:=range s.skin.GetTopics(){
					println(topic)
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
