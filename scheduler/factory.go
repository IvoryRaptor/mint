package scheduler

import "github.com/IvoryRaptor/dragonfly"

type Factory struct {
}

func (f * Factory)GetName() string{
	return "scheduler"
}

func (f * Factory)Create(kernel dragonfly.IKernel,config map[interface {}]interface{}) (dragonfly.IService,error) {
	result := Service{}
	result.Config(kernel, config)
	return &result, nil
}
