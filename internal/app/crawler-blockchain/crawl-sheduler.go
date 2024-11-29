package crawlerblockchain

import (
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func (service *CrawlerBlockchain) produce() {
	ticker := time.Tick(3 * time.Second)
	defer func() {
		service.wg.Done()
		fmt.Println("Producer is exiting")
	}()
ProducerLoop:
	for {
		select {
		case <-service.ctx.Done():
			break ProducerLoop
		case <-ticker:
			fmt.Println("producer writing message", time.Now().UTC().String())
			err := service.deps.KafkaClient.Writer.WriteMessages(service.ctx, kafka.Message{Topic: "my-topic", Value: []byte(time.Now().UTC().String())})
			if err != nil {
				fmt.Println("write error", err)
			}
		}
	}
}
