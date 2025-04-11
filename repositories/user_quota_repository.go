package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
)

type UserQuotaRepository struct {
	Database *gorm.DB
}

func NewUserQuotaRepository(db *gorm.DB) *UserQuotaRepository {
	return &UserQuotaRepository{
		Database: db,
	}
}

type UserQuotaRepositoryInterface interface {
	GetByID(id uint) (*models.UserQuota, error)
	GetByUserID(userID uint) ([]models.UserQuota, error)
	Create(userQuota *models.UserQuota) error
	Update(userQuota *models.UserQuota) error
	Delete(id uint) error
	GetByUserIDAndQuotaID(userID, quotaID uint) (*models.UserQuota, error)
	CountByUserID(userID uint) (int64, error)
	GetByUserIDAndQuotaName(userID uint, quotaName string) (*models.UserQuota, error)
}

// GetByID retrieves a UserQuota record from the database by its ID.
// It takes an unsigned integer `id` as input and returns a pointer to a `models.UserQuota` instance and an error.
// If the UserQuota with the specified ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserQuotaRepository) GetByID(id uint) (*models.UserQuota, error) {
	var userQuota models.UserQuota
	err := r.Database.First(&userQuota, id).Error
	if err != nil {
		return nil, err
	}
	return &userQuota, nil
}

// GetByUserID retrieves all UserQuota records associated with a specific user ID.
// It takes an unsigned integer `userID` as input and returns a slice of `models.UserQuota` instances and an error.
// If there are no UserQuota records for the specified user ID or if there is a database error, it returns a non-nil error.
func (r *UserQuotaRepository) GetByUserID(userID uint) ([]models.UserQuota, error) {
	var userQuotas []models.UserQuota
	err := r.Database.Where("user_id = ?", userID).Find(&userQuotas).Error
	if err != nil {
		return nil, err
	}
	return userQuotas, nil
}

// Create inserts a new UserQuota record into the database.
// It takes a pointer to a `models.UserQuota` instance as input and returns an error.
// If the create operation fails, it returns a non-nil error.
func (r *UserQuotaRepository) Create(userQuota *models.UserQuota) error {
	err := r.Database.Create(userQuota).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing UserQuota record in the database.
// It takes a pointer to a `models.UserQuota` instance as input and returns an error.
// If the update operation fails, it returns a non-nil error.
func (r *UserQuotaRepository) Update(userQuota *models.UserQuota) error {
	err := r.Database.Save(userQuota).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a UserQuota record from the database by its ID.
// It takes an unsigned integer `id` as input and returns an error.
// If the UserQuota with the specified ID is not found or if the delete operation fails, it returns a non-nil error.
func (r *UserQuotaRepository) Delete(id uint) error {
	var userQuota models.UserQuota
	err := r.Database.First(&userQuota, id).Error
	if err != nil {
		return err
	}
	err = r.Database.Delete(&userQuota).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByUserIDAndQuotaID retrieves a UserQuota record by user ID and quota ID.
// It takes two unsigned integers `userID` and `quotaID` as input and returns a pointer to a `models.UserQuota` instance and an error.
// If the UserQuota with the specified user ID and quota ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserQuotaRepository) GetByUserIDAndQuotaID(userID, quotaID uint) (*models.UserQuota, error) {
	var userQuota models.UserQuota
	err := r.Database.Where("user_id = ? AND quota_id = ?", userID, quotaID).First(&userQuota).Error
	if err != nil {
		return nil, err
	}
	return &userQuota, nil
}

// CountByUserID counts the number of UserQuota records associated with a specific user ID.
// It takes an unsigned integer `userID` as input and returns the count as an int64 and an error.
// If there is a database error, it returns a non-nil error.
func (r *UserQuotaRepository) CountByUserID(userID uint) (int64, error) {
	var count int64
	err := r.Database.Model(&models.UserQuota{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetByUserIDAndQuotaName retrieves a UserQuota record by user ID and quota name.
// It takes an unsigned integer `userID` and a string `quotaName` as input and returns a pointer to a `models.UserQuota` instance and an error.
// If the UserQuota with the specified user ID and quota name is not found or if there is a database error, it returns a non-nil error.
func (r *UserQuotaRepository) GetByUserIDAndQuotaName(userID uint, quotaName string) (*models.UserQuota, error) {
	var userQuota models.UserQuota
	err := r.Database.Where("user_id = ? AND quota_name = ?", userID, quotaName).First(&userQuota).Error
	if err != nil {
		return nil, err
	}
	return &userQuota, nil
}
