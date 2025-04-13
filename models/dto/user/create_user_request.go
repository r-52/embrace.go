package user

type CreateUserRequest struct {
	Email     string `form:"email" json:"email" binding:"required,email" validate:"required,email"`
	Password  string `form:"password" json:"password" binding:"required,min=8" validate:"required,min=8"`
	CompanyID uint   `form:"companyId" json:"companyId" binding:"required,min=1" validate:"required,gte=1"`
}
