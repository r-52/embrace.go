package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
)

type QuotaRepository struct {
	Database *gorm.DB
}
type QuotaRepositoryInterface interface {
	GetByID(id uint) (*models.Quota, error)
	GetByCompanyID(companyID uint) ([]models.Quota, error)
	Create(quota *models.Quota) error
	Update(quota *models.Quota) error
	Delete(id uint) error
	GetByCompanyIDAndName(companyID uint, name string) (*models.Quota, error)
	CountByCompanyID(companyID uint) (int64, error)
}

func NewQuotaRepository(db *gorm.DB) *QuotaRepository {
	return &QuotaRepository{
		Database: db,
	}
}
func (r *QuotaRepository) GetByID(id uint) (*models.Quota, error) {
	var quota models.Quota
	err := r.Database.First(&quota, id).Error
	if err != nil {
		return nil, err
	}
	return &quota, nil
}
func (r *QuotaRepository) GetByCompanyID(companyID uint) ([]models.Quota, error) {
	var quotas []models.Quota
	err := r.Database.Where("company_id = ?", companyID).Find(&quotas).Error
	if err != nil {
		return nil, err
	}
	return quotas, nil
}
func (r *QuotaRepository) Create(quota *models.Quota) error {
	err := r.Database.Create(quota).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *QuotaRepository) Update(quota *models.Quota) error {
	err := r.Database.Save(quota).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *QuotaRepository) Delete(id uint) error {
	var quota models.Quota
	err := r.Database.First(&quota, id).Error
	if err != nil {
		return err
	}
	err = r.Database.Delete(&quota).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *QuotaRepository) GetByCompanyIDAndName(companyID uint, name string) (*models.Quota, error) {
	var quota models.Quota
	err := r.Database.Where("company_id = ? AND name = ?", companyID, name).First(&quota).Error
	if err != nil {
		return nil, err
	}
	return &quota, nil
}
func (r *QuotaRepository) CountByCompanyID(companyID uint) (int64, error) {
	var count int64
	err := r.Database.Model(&models.Quota{}).Where("company_id = ?", companyID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
