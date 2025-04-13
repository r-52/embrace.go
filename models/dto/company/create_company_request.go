package company

type CreateCompanyRequest struct {
	Name string `form:"name" json:"name" binding:"required" validate:"required,min=1,max=255"`
}
