package user

type CreateUserRequest struct {
	Email           string `form:"email" json:"email" binding:"required,email" validate:"required,email"`
	Password        string `form:"password" json:"password" binding:"required,min=8" validate:"required,min=8"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"required,min=8" validate:"required,min=8,eqfield=Password"`
	CompanyID       uint   `form:"companyId" json:"companyId" binding:"required,min=1" validate:"required,gte=1"`
	FirstName       string `form:"firstName" json:"firstName" binding:"min=2,max=50" validate:"min=2,max=50"`
	LastName        string `form:"lastName" json:"lastName" binding:"min=2,max=50" validate:"min=2,max=50"`
	Phone           string `form:"phone" json:"phone" binding:"min=10,max=15" validate:"min=10,max=15"`
	Title           string `form:"title" json:"title" binding:"min=2,max=50" validate:"min=2,max=50"`
	Position        string `form:"position" json:"position" binding:"min=2,max=50" validate:"min=2,max=50"`
	Location        string `form:"location" json:"location" binding:"min=2,max=50" validate:"min=2,max=50"`
}
