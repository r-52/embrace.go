package user

type CreateUserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	CompanyID uint   `json:"companyId"`
}
