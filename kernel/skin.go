package kernel

import (
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/postoffice"
	"github.com/IvoryRaptor/dragonfly/mq"
	"github.com/golang/protobuf/proto"
	"log"
)

type Skin struct {
	dragonfly.Kernel
	mq        mq.IMQ
	zookeeper *dragonfly.Zookeeper
}

func (skin *Skin) GetTopics() []string {
	result := make([]string, 0)
	for key, v := range skin.zookeeper.GetChildes() {
		for _, topic := range v.GetKeys() {
			result = append(result, key+"_"+topic)
		}
	}
	return result
}

func (skin *Skin) Arrive(data []byte) {
	msg := postoffice.MQMessage{}
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (skin *Skin) Publish(topic string, actor []byte, payload []byte) error {
	return skin.mq.Publish(topic, actor, payload)
}

func (skin *Skin) SetFields() {
	skin.zookeeper = skin.GetService("zookeeper").(*dragonfly.Zookeeper)
	skin.mq = skin.GetService("mq").(mq.IMQ)
}
