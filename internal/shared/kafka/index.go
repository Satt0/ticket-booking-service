package kafka

import (
	"context"
	"fmt"
	"http-server/internal/shared/logger"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
)

type KafkaClient struct {
	Writer *kafka.Writer
}

func NewKafkaClient(lc fx.Lifecycle, logger *logger.Logger) *KafkaClient {
	writer := &kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		BatchSize: 1,
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			err := writer.Close()
			if err == nil {
				fmt.Println("KAFKA_WRITER closed")
			} else {
				fmt.Println("KAFKA_WRITER not closed ", err.Error())
			}
			return nil
		},
	})
	return &KafkaClient{
		Writer: writer,
	}
}

// consumer must be inited when used
func (k *KafkaClient) CreateKafkaReader(topic string, groupId string) (*kafka.Reader, func()) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          []string{"localhost:9092"},
		Topic:            topic,
		GroupID:          groupId,
		ReadBatchTimeout: time.Duration(1 * time.Second),
		MaxBytes:         10e6, // 10MB

	})
	return r, func() {
		err := r.Close()
		if err != nil {
			fmt.Printf("KAFKA_READER for topic %s group %s CANNOT closed\n", topic, groupId)
		} else {
			fmt.Printf("KAFKA_READER for topic %s group %s closed\n", topic, groupId)
		}
	}
}

var KafkaModuleFx = fx.Provide(NewKafkaClient)
