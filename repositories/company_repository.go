package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	Database *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		Database: db,
	}
}
func (r *CompanyRepository) GetByID(id uint) (*models.Company, error) {
	var company models.Company
	err := r.Database.First(&company, id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}
func (r *CompanyRepository) Create(company *models.Company) error {
	err := r.Database.Create(company).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *CompanyRepository) Update(company *models.Company) error {
	err := r.Database.Save(company).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *CompanyRepository) Delete(id uint) error {
	var company models.Company
	err := r.Database.First(&company, id).Error
	if err != nil {
		return err
	}
	err = r.Database.Delete(&company).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *CompanyRepository) GetByUserID(userID uint) (*models.Company, error) {
	var company models.Company
	err := r.Database.Joins("JOIN users ON users.company_id = companies.id").Where("users.id = ?", userID).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}
