package kernel

import (
	"github.com/IvoryRaptor/dragonfly"
	"github.com/IvoryRaptor/postoffice"
	"github.com/golang/protobuf/proto"
	"log"
	"github.com/IvoryRaptor/mint/mq"
	//"github.com/IvoryRaptor/mint"
	"github.com/IvoryRaptor/mint"
	"sync"
)

type Mint struct {
	dragonfly.Kernel
	mq        *mq.Kafka
	Zookeeper *dragonfly.Zookeeper
	matrixMap *sync.Map
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
	mit := mint.MintPayload{}
	err = proto.Unmarshal(msg.Payload, &mit)
	if err != nil {
		log.Println(err.Error())
		return
	}
	t, _ := m.matrixMap.LoadOrStore(msg.Source.Matrix, &sync.Map{})
	matrix := t.(*sync.Map)

	t, _ = matrix.LoadOrStore(msg.Source.Device, &sync.Map{})

	angler := t.(*sync.Map)
	t, _ = angler.LoadOrStore(mit.Number, &sync.Map{})
	p := t.(*sync.Map)
	p.Store(mit.Partition, mit.Time)

	m.matrixMap.Range(func(key, value interface{}) bool {
		s := value.(*sync.Map)
		s.Range(func(key, value interface{}) bool {
			return true
		})
		return true
	})
}

func (m *Mint) SetFields() {
	m.Zookeeper = m.GetService("zookeeper").(*dragonfly.Zookeeper)
	m.mq = m.GetService("mq").(*mq.Kafka)
	m.matrixMap = &sync.Map{}
}

func (m *Mint) GetMatrix() *sync.Map {
	return m.matrixMap
}
