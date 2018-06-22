package kernel

import (
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/postoffice"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/IvoryRaptor/mint/mq"
	//"github.com/IvoryRaptor/mint"
	mint "github.com/IvoryRaptor/mint"
)

type Mint struct {
	dragonfly.Kernel
	mq        *mq.Kafka
	Zookeeper *dragonfly.Zookeeper
}

func (m *Mint) GetTopics() []string {
	result := make([]string, 0)
	for key, v := range m.Zookeeper.GetChildes() {
		for _, topic := range v.GetKeys() {
			result = append(result, key+"_"+topic)
		}
	}
	return result
}

func (m *Mint) Publish(topic string, partition int32, actor []byte, payload []byte) error {
	return m.mq.Publish(topic, partition, actor, payload)
}

func (m *Mint) Arrive(data []byte) {
	msg := postoffice.MQMessage{}
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}

	println(msg.Source.Matrix, msg.Source.Device)
	println(msg.Destination.Matrix, msg.Destination.Device)
	println(msg.Resource, msg.Action)

	mit := mint.MintPayload{}
	err = proto.Unmarshal(msg.Payload, &mit)
	if err != nil {
		log.Println(err.Error())
		return
	}
	println("--------------")
	println(mit.Partition, mit.Number, mit.Time)
}

func (m *Mint) SetFields() {
	m.Zookeeper = m.GetService("zookeeper").(*dragonfly.Zookeeper)
	m.mq = m.GetService("mq").(*mq.Kafka)
}
