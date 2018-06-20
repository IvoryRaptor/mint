package kernel

import (
	"github.com/IvoryRaptor/skin/mq"
	"github.com/IvoryRaptor/skin"
	"github.com/IvoryRaptor/dragonfly"
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

func (skin *Skin) Arrive(msg *skin.MQMessage) {

}

func (skin *Skin) Publish(topic string, payload []byte) error {
	return nil
}

func (skin *Skin) SetFields() {
	skin.zookeeper = skin.GetService("zookeeper").(*dragonfly.Zookeeper)
}
