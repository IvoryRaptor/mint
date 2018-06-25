package mint

import (
	"github.com/IvoryRaptor/dragonfly/mq"
	"sync"
)

type IMint interface {
	mq.IArrive
	GetTopics() []string
	Publish(topic string, partition int32, actor []byte, payload []byte) error
	GetMatrix() *sync.Map
}
