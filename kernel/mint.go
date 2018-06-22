package kernel

import (
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/postoffice"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/IvoryRaptor/mint/mq"
)

type Mint struct {
	dragonfly.Kernel
	mq        *mq.Kafka
	Zookeeper *dragonfly.Zookeeper
}

func (mint *Mint) GetTopics() []string {
	result := make([]string, 0)
	for key, v := range mint.Zookeeper.GetChildes() {
		for _, topic := range v.GetKeys() {
			result = append(result, key+"_"+topic)
		}
	}
	return result
}

func (mint *Mint) Publish(topic string, partition int32, actor []byte, payload []byte) error {
	return mint.mq.Publish(topic, partition, actor, payload)
}

func (mint *Mint) Arrive(data []byte) {
	msg := postoffice.MQMessage{}
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (mint *Mint) SetFields() {
	mint.Zookeeper = mint.GetService("zookeeper").(*dragonfly.Zookeeper)
	mint.mq = mint.GetService("mq").(*mq.Kafka)
}
