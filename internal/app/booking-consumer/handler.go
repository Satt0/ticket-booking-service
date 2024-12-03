package bookingjob

import "fmt"

func (service *BookingJob) ConsumeBookingJob() {
	service.wg.Add(1)
	consumer, closeConsumer := service.deps.KafkaClient.CreateKafkaReader("create-booking", "group-1")
	go func() {
		closeConsumer()
		defer service.wg.Done()
	Loop:
		for {
			select {
			case <-service.ctx.Done():
				break Loop
			default:
				message, err := consumer.FetchMessage(service.ctx)
				if err != nil {
					break
				}
				fmt.Printf("message.Value: %v\n", message.Value)
				// todo handle batch of messages
				service.handleConsumeBookingMessage(message.Value)

				consumer.CommitMessages(service.ctx, message)
				service.deps.Logger.Out.Info("message handled", message.Topic, message.Offset, message.Partition)
			}
		}
	}()
}
