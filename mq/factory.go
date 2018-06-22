package mq

import (
	"github.com/IvoryRaptor/dragonfly"
)

type Factory struct {
}

func (f *Factory) GetName() string {
	return "mq"
}

func (f *Factory) Create(kernel dragonfly.IKernel, config map[interface{}]interface{}) (dragonfly.IService, error) {
	m := &Kafka{}
	m.Config(kernel, config)
	return m, nil
}
