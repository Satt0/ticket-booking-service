package entities

import (
	"encoding/json"

	"gorm.io/datatypes"
)

type TicketOrderOption struct {
	EventId         uint64  `json:"eventId"`
	OptionId        uint64  `json:"sptionId"`
	Amount          uint64  `json:"amount"`
	TotalPrice      float64 `json:"totalPrice"`
	DiscountedPrice float64 `json:"discountedPrice"`
}
type Order struct {
	BaseModel
	UserId        uint64         `gorm:"int8;not null" json:"userId"`
	EventId       uint64         `gorm:"int8;not null" json:"eventId"`
	Status        string         `gorm:"size:255;not null" json:"status"`
	TicketOptions datatypes.JSON `gorm:"jsonb;not null" json:"ticketOptions"`
	// Gender        string `gorm:"size:255;not null" json:"gender"`
	// Balance       string `gorm:"size:255;not null" json:"balance"`
}

func (o *Order) SetTicketOptions(t []*TicketOrderOption) (*Order, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return &Order{}, err
	}
	o.TicketOptions = datatypes.JSON(bytes)
	return o, nil
}
func (o *Order) GetTicketOptions() ([]TicketOrderOption, error) {
	var x []TicketOrderOption
	err := json.Unmarshal(o.TicketOptions, &x)
	if err != nil {
		return nil, err
	}
	return x, nil
}
