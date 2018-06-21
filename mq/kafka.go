package mq

import (
	"github.com/IvoryRaptor/dragonfly/mq"
	"github.com/IvoryRaptor/dragonfly"
)

type Kafka struct {
	mq.Kafka
	partition int
}

func (k * Kafka)Publish(topic string,actor []byte,payload []byte) error {
	for i := 0; i < k.partition; i++ {
		k.KafkaPublish(topic, int32(i), actor, payload)
	}
	return nil
}

func (k *Kafka) Config(kernel dragonfly.IKernel, config map[interface{}]interface{}) error {
	err := k.KafkaConfig(kernel, config)
	if err!=nil{
		return err
	}
	k.partition = config["partition"].(int)
	return nil
}
