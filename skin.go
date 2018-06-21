package skin

import (
	"github.com/IvoryRaptor/dragonfly/mq"
)

type ISkin interface {
	mq.IArrive
	Publish(topic string, actor []byte, payload []byte) error
	GetTopics() []string
}
