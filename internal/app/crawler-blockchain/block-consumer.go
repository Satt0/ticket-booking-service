package crawlerblockchain

import (
	"fmt"
	"time"
)

func (service *CrawlerBlockchain) consume() {
	reader, closeReader := service.deps.KafkaClient.CreateKafkaReader("my-topic", "my-group-id")
	defer func() {
		closeReader()
		service.wg.Done()
		fmt.Println("consumer is exiting")
	}()
	for {
		m, err := reader.FetchMessage(service.ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), time.Now().UTC().String())
		// handle processing message, take in account a message can be processed twice without creating duplicating actions.
		err_commit := reader.CommitMessages(service.ctx, m)
		if err_commit != nil {
			service.deps.Logger.Out.Error("cannot commit message", m.Offset)
		}
	}
}
