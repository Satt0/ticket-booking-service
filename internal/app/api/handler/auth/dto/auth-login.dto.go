package auth_dto

type LoginDto struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"id" binding:"required"`
}
