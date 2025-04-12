package company

type CreateCompanyRequest struct {
	Name string `json:"name" binding:"required"`
}
