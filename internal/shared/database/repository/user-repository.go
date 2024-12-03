package repository

import (
	"database/sql"
	"errors"
	"fmt"
	users_dto "http-server/internal/app/api/handler/orders/dto"
	"http-server/internal/shared/database/entities"
	"http-server/internal/shared/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r OrderRepository) FindAll() ([]entities.Order, error) {
	var users []entities.Order
	result := r.db.Find(&users)
	return users, result.Error
}

func (r OrderRepository) FindByID(id uint64) (*entities.Order, error) {
	var order *entities.Order
	result := r.db.First(&order, id)
	return order, result.Error
}

func (r OrderRepository) Create(order entities.Order) (entities.Order, error) {
	result := r.db.Create(&order)
	return order, result.Error
}

func (r OrderRepository) SaveOrderAndOutbox(dto users_dto.CreateOrderReqDto) (*entities.Order, error) {
	order := entities.Order{
		Status: "pending",
		UserId: 999,
	}
	order.SetTicketOptions(utils.Map(dto.Cart, func(v users_dto.OrderTicketOptionReqDto) *entities.TicketOrderOption {
		return &entities.TicketOrderOption{
			EventId:         dto.EventId,
			Amount:          v.Amount,
			TotalPrice:      123,
			DiscountedPrice: 999,
		}
	}))
	uniqueOptionIds := utils.Map(dto.Cart, func(v users_dto.OrderTicketOptionReqDto) uint64 { return v.OptionId })
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var options []entities.TicketOption
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("option_id IN(?) AND event_id = ?",
				uniqueOptionIds, dto.EventId).Find(&options).Error; err != nil {
			return err
		}
		if len(options) == 0 || (len(uniqueOptionIds) != len(options)) {
			return errors.New("missmatch options count")
		}
		// minus stock value
		for _, v := range dto.Cart {
			res := tx.Exec(`UPDATE ticket_options SET current_capacity = current_capacity - ? WHERE id = ? AND current_capacity >= ?;`,
				v.Amount, v.OptionId, v.Amount)
			if res.RowsAffected != 1 {
				if res.Error != nil {
					return res.Error
				}
				return errors.New("error in updating ticket stock")
			}
		}
		return tx.Create(&order).Error
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return &entities.Order{}, err
	}
	return &order, nil
}
