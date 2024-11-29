package repository

import (
	"errors"
	"http-server/internal/shared/database/entities"
	"math/big"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) FindAll() ([]entities.Users, error) {
	var users []entities.Users
	result := r.db.Find(&users)
	return users, result.Error
}

func (r UserRepository) FindByID(id uint64) (*entities.Users, error) {
	var user *entities.Users
	result := r.db.First(&user, id)
	return user, result.Error
}
func (r UserRepository) FindByEmail(email string) (*entities.Users, error) {
	var user *entities.Users
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r UserRepository) Create(user entities.Users) (entities.Users, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r UserRepository) Update(user entities.Users) error {
	return r.db.Save(&user).Error
}

func (r UserRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Users{}, id).Error
}

func (r UserRepository) UpdateBalanceAtomic(id uint64) (*entities.Users, error) {
	var user *entities.Users
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, id).Error; err != nil {
			return err
		}
		oldBalance, ok := new(big.Int).SetString(user.Balance, 10)
		if !ok {
			return errors.New("canot add balance")
		}
		newBalance := oldBalance.Add(oldBalance, big.NewInt(10)).String()
		user.Balance = newBalance
		return tx.Exec("UPDATE users set balance = ? where id = ?", user.Balance, user.ID).Error
	})
	if err != nil {
		return &entities.Users{}, errors.New("cannot update user balance")
	}
	return user, nil
}
