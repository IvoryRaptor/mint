package skin

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Skin struct {
	host         string
	ConfigFile   string
	run          bool
	producer *kafka.Producer
	consumer *kafka.Consumer

}

func (skin *Skin)GetHost() string{

}


func (skin *Skin)Start() error{

}

func (skin *Skin)Publish(topic string) error {
	deliveryChan := make(chan kafka.Event)

	for i := 0; i < 10; i++ {
		fmt.Printf("i 的值为: %d\n", i)
	}

	err := skin.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 1},
			Key:            []byte("heartbeat"),
			Value:          []byte{},
		},
		deliveryChan)
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	if err != nil {
		fmt.Println(topic, err.Error())
		return err
	}
	return nil
}

func (skin *Skin)GetTopics(matrix string, action string) ([]string, bool){

}

func (skin *Skin)Close(device string){

}
