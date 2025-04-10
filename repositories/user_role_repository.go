package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	Database *gorm.DB
}

// NewUserRoleRepository creates a new instance of UserRoleRepository with the provided database connection.
// It takes a *gorm.DB as an argument, which represents the database connection, and returns a pointer to a UserRoleRepository.
func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{
		Database: db,
	}
}

// GetByID retrieves a user role record from the database by its ID.
// It takes an unsigned integer `id` as input and returns a pointer to a
// `models.UserRole` instance and an error. If the user role with the specified
// ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserRoleRepository) GetByID(id uint) (*models.UserRole, error) {
	var userRole models.UserRole
	err := r.Database.First(&userRole, id).Error
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}

// Create inserts a new user role record into the database.
// It takes a pointer to a `models.UserRole` instance as input and returns an error.
// If the create operation fails, it returns a non-nil error.
func (r *UserRoleRepository) Create(userRole *models.UserRole) error {
	err := r.Database.Create(userRole).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing user role record in the database.
// It takes a pointer to a `models.UserRole` instance as input and returns an error.
// If the update operation fails, it returns a non-nil error.
func (r *UserRoleRepository) Update(userRole *models.UserRole) error {
	err := r.Database.Save(userRole).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a user role record from the database by its ID.
// It takes an unsigned integer `id` as input and returns an error.
// If the user role with the specified ID is not found or if the delete operation fails, it returns a non-nil error.
func (r *UserRoleRepository) Delete(id uint) error {
	var userRole models.UserRole
	err := r.Database.First(&userRole, id).Error
	if err != nil {
		return err
	}
	err = r.Database.Delete(&userRole).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByCompanyID retrieves all user roles associated with a specific company ID.
// It takes an unsigned integer `companyID` as input and returns a slice of `models.UserRole`
// instances and an error. If there is a database error, it returns a non-nil error.
func (r *UserRoleRepository) GetByCompanyID(companyID uint) ([]models.UserRole, error) {
	var userRoles []models.UserRole
	err := r.Database.Where("company_id = ?", companyID).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}

// CountByCompanyID counts the number of user roles associated with a specific company ID.
// It takes an unsigned integer `companyID` as input and returns the count as an int64 and an error.
// If there is a database error, it returns a non-nil error.
func (r *UserRoleRepository) CountByCompanyID(companyID uint) (int64, error) {
	var count int64
	err := r.Database.Model(&models.UserRole{}).Where("company_id = ?", companyID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
