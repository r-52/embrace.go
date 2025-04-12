package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
)

type TimeEntryTypeRepository struct {
	Database *gorm.DB
}

type TimeEntryTypeRepositoryInterface interface {
	GetByID(id uint) (*models.TimeEntryType, error)
	Create(timeEntryType *models.TimeEntryType) error
	Update(timeEntryType *models.TimeEntryType) error
	Delete(id uint) error
	GetByCompanyID(companyID uint) ([]models.TimeEntryType, error)
	CountByCompanyID(companyID uint) (int64, error)
}

func NewTimeEntryTypeRepository(db *gorm.DB) *TimeEntryTypeRepository {
	return &TimeEntryTypeRepository{
		Database: db,
	}
}

func (r *TimeEntryTypeRepository) GetByID(id uint) (*models.TimeEntryType, error) {
	var timeEntryType models.TimeEntryType
	err := r.Database.First(&timeEntryType, id).Error
	if err != nil {
		return nil, err
	}
	return &timeEntryType, nil
}

func (r *TimeEntryTypeRepository) Create(timeEntryType *models.TimeEntryType) error {
	err := r.Database.Create(timeEntryType).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TimeEntryTypeRepository) Update(timeEntryType *models.TimeEntryType) error {
	err := r.Database.Save(timeEntryType).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *TimeEntryTypeRepository) Delete(id uint) error {
	var timeEntryType models.TimeEntryType
	err := r.Database.First(&timeEntryType, id).Error
	if err != nil {
		return err
	}
	err = r.Database.Delete(&timeEntryType).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TimeEntryTypeRepository) GetByCompanyID(companyID uint) ([]models.TimeEntryType, error) {
	var timeEntryTypes []models.TimeEntryType
	err := r.Database.Where("company_id = ?", companyID).Find(&timeEntryTypes).Error
	if err != nil {
		return nil, err
	}
	return timeEntryTypes, nil
}

func (r *TimeEntryTypeRepository) CountByCompanyID(companyID uint) (int64, error) {
	var count int64
	err := r.Database.Model(&models.TimeEntryType{}).Where("company_id = ?", companyID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
