package repository

import (
	"errors"
	"fmt"
	users_dto "http-server/internal/app/api/handler/orders/dto"
	"http-server/internal/shared/database/entities"
	"http-server/internal/shared/utils"
	"strings"

	"gorm.io/gorm"
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
		// minus stock value
		// todo use bulk update
		generatedQuery := utils.Map(dto.Cart, func(v users_dto.OrderTicketOptionReqDto) string {
			return fmt.Sprintf("(%d,%d)", v.OptionId, v.Amount)
		})
		res := tx.Exec(fmt.Sprintf(`
			UPDATE ticket_options as target
				SET 
					current_capacity = target.current_capacity - c.amount
				FROM (
					VALUES
						%s
				) AS c(id, amount)
			WHERE 
				target.event_id = ? AND 
				c.id = target.option_id AND
				target.current_capacity >= c.amount;
		`, strings.Join(generatedQuery, ",")), dto.EventId)
		if res.RowsAffected != int64(len(uniqueOptionIds)) {
			if res.Error != nil {
				return res.Error
			}
			return errors.New("error in updating ticket stock")
		}
		return tx.Create(&order).Error
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return &entities.Order{}, err
	}
	return &order, nil
}
