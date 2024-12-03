package bookingjob

func (service *BookingJob) handleConsumeBookingMessage(payload interface{}) error {
	// check if user has any pending order

	// save order to db, include outbox record

	// commit
	return nil
}
