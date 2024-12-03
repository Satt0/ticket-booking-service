package entities

type TicketOption struct {
	BaseModel
	EventId  uint64 `gorm:"int8;not null" json:"eventId"`
	OptionId uint64 `gorm:"int8;not null" json:"optionId"`

	TotalCapacity   uint64  `gorm:"int8;not null" json:"status"`
	CurrentCapacity uint64  `gorm:"int8;not null" json:"currentCapacity"`
	Price           float64 `gorm:"numeric:100,2;not null" json:"price"`
}
