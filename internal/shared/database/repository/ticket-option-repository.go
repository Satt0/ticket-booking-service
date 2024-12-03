package repository

import (
	"http-server/internal/shared/database/entities"

	"gorm.io/gorm"
)

type TicketOptionRepository struct {
	db *gorm.DB
}

func NewTicketOptionRepository(db *gorm.DB) *TicketOptionRepository {
	return &TicketOptionRepository{db: db}
}

func (r TicketOptionRepository) FindAll() ([]entities.TicketOption, error) {
	var order []entities.TicketOption
	result := r.db.Find(&order)
	return order, result.Error
}

func (r TicketOptionRepository) FindByID(id uint64) (*entities.TicketOption, error) {
	var order *entities.TicketOption
	result := r.db.First(&order, id)
	return order, result.Error
}

func (r TicketOptionRepository) Create(order entities.TicketOption) (entities.TicketOption, error) {
	result := r.db.Create(&order)
	return order, result.Error
}
