package company

import (
	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/models/dto/company"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

type CompanyCreator struct {
	CompanyRepository repositories.CompanyRepositoryInterface
}

type CompanyCreatorInterface interface {
	CreateCompany(req *company.CreateCompanyRequest) (*models.Company, error)
}

func NewCompanyCreator(db *gorm.DB) *CompanyCreator {
	return &CompanyCreator{
		CompanyRepository: repositories.NewCompanyRepository(db),
	}
}

// CreateCompany creates a new company in the database.
// It takes a pointer to a `company.CreateCompanyRequest` instance as input and returns a pointer to a `models.Company` instance and an error.
// If the creation fails, it returns a non-nil error.
func (c *CompanyCreator) CreateCompany(req *company.CreateCompanyRequest) (*models.Company, error) {
	return nil, nil
}
