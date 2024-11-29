package user_services

import (
	"http-server/internal/shared/database/entities"
	"http-server/internal/shared/database/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: *userRepo,
	}
}

func (us *UserService) GetAllUser() []entities.Users {
	users, _ := us.userRepo.FindAll()
	return users
}
func (us *UserService) GetUserById(id uint64) (*entities.Users, string) {
	user, err := us.userRepo.FindByID(id)
	if err != nil {
		return &entities.Users{}, "Not found"
	}
	return user, ""
}
