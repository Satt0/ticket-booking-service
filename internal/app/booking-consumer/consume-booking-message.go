package bookingjob

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	TicketTypeId uint64 `json:"ticketTypeId"`
	EventId      uint64 `json:"eventId"`
	Amount       string `json:"amount"`
}

func (service *BookingJob) handleConsumeBookingMessage(payload []byte) error {
	var x Test
	err := json.Unmarshal(payload, &x)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("x: %v\n", x)
	return nil
}
