package company

import "github.com/r-52/embrace/models/dto/user"

type CreateCompanyRequest struct {
	Name string                  `form:"name" json:"name" binding:"required" validate:"required,min=1,max=255"`
	User *user.CreateUserRequest `form:"user" json:"user" binding:"required" validate:"required"`
}
