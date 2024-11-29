package checkpayment

import (
	"fmt"
	"time"
)

func (service *CheckOutdatedPayment) handleCheckOutdatedPayment(id int) {
	service.wg.Add(1)
	go func() {
		defer service.wg.Done()
	Loop:
		for {
			select {
			case <-service.ctx.Done():
				fmt.Println("signaled stopped", id)
				break Loop
			default:
				fmt.Println("in looping", id)
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (service *CheckOutdatedPayment) HandleCronCheck() {
	service.handleCheckOutdatedPayment(1)
}
