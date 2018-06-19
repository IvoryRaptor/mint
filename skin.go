package skin

import (
	"github.com/IvoryRaptor/dragonfly"
)

type ISkin interface {
	dragonfly.IKernel
	Publish(topic string, payload []byte) error
	Arrive(msg *MQMessage)
	GetTopics() []string
}
