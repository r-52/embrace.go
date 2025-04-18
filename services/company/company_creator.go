package company

import (
	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/models/dto/company"
	"github.com/r-52/embrace/repositories"
	"github.com/r-52/embrace/services/user"
	"gorm.io/gorm"
)

type CompanyCreator struct {
	companyRepository *repositories.CompanyRepository
	userCreator       user.UserCreator
}

type CompanyCreatorInterface interface {
	CreateCompany(req *company.CreateCompanyRequest) (*models.Company, error)
}

func NewCompanyCreator(db *gorm.DB) *CompanyCreator {
	return &CompanyCreator{
		companyRepository: repositories.NewCompanyRepository(db),
		userCreator:       *user.NewUserCreator(db),
	}
}

// CreateCompany creates a new company in the database.
// It takes a pointer to a `company.CreateCompanyRequest` instance as input and returns a pointer to a `models.Company` instance and an error.
// If the creation fails, it returns a non-nil error.
func (c *CompanyCreator) CreateCompany(req *company.CreateCompanyRequest) (*models.Company, error) {
	company := &models.Company{
		Name:         req.Name,
		Description:  req.Description,
		Website:      req.Website,
		PrimaryEmail: req.User.Email,
	}
	err := c.companyRepository.Create(company)
	if err != nil {
		return nil, err
	}
	req.User.CompanyID = company.ID

	_, err = c.userCreator.CreateUser(req.User)
	if err != nil {
		return nil, err
	}

	return company, nil
}
