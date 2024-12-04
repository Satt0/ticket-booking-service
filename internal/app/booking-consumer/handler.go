package bookingjob

import "fmt"

func (service *BookingJob) ConsumeBookingJob() {
	service.wg.Add(1)
	consumer, closeConsumer := service.deps.KafkaClient.CreateKafkaReader("create_ticket_type", "my-group-1")
	defer func() {
		closeConsumer()
		defer service.wg.Done()

	}()
Loop:
	for {
		select {
		case <-service.ctx.Done():
			fmt.Println("exiting")
			break Loop
		default:
			fmt.Println("fetching messages")
			message, err := consumer.FetchMessage(service.ctx)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				break
			}
			service.handleConsumeBookingMessage(message.Value)

			consumer.CommitMessages(service.ctx, message)
			service.deps.Logger.Out.Info("message handled", message.Topic, message.Offset, message.Partition)
		}
	}
}
