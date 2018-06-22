package mq

import (
	"github.com/IvoryRaptor/dragonfly/mq"
	"github.com/IvoryRaptor/dragonfly"
)

type Kafka struct {
	mq.Kafka
}

func (k *Kafka) Publish(topic string, partition int32, actor []byte, payload []byte) error {
	return k.KafkaPublish(topic, partition, actor, payload)
}

func (k *Kafka) Config(kernel dragonfly.IKernel, config map[interface{}]interface{}) error {
	err := k.KafkaConfig(kernel, config)
	if err != nil {
		return err
	}
	return nil
}
