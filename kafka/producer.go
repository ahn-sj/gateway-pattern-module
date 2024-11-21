package kafka

import (
	"gateway-module/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

const (
	_allAcks = "all"
)

type Producer struct {
	cfg      config.Producer
	Producer *kafka.Producer
}

func NewProducer(cfg config.Producer) Producer {
	url := cfg.URL
	id := cfg.ClientID
	acks := cfg.Acks
	if acks == "" {
		acks = _allAcks
	}
	conf := &kafka.ConfigMap{
		"bootstrap.servers": url,  // kafka broker url
		"client.id":         id,   // produce client id
		"acks":              acks, // 0, 1, all
	}

	producer, err := kafka.NewProducer(conf)
	if err != nil {
		panic(err.Error()) // 모듈 자체를 사용할 수 없는 상황
	}
	return Producer{
		cfg:      cfg,
		Producer: producer,
	}
}

func (p Producer) SendEvent(v []byte) { // 외부에서 직렬화 / 역직렬화를 처리할 예정이라 바이트로 받음
	topic := p.cfg.Topic

	err := p.Producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: v,
		}, nil)
	if err != nil {
		log.Println("Failed to produce message", string(v))
	} else {
		log.Println("Successed to produce message", string(v))
	}
}
