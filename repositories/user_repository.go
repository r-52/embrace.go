package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *gorm.DB
}

type UserRepositoryInterface interface {
	// GetByID retrieves a user record from the database by its ID.
	// It takes an unsigned integer `id` as input and returns a pointer to a `models.User` instance and an error.
	GetByID(id uint) (*models.User, error)

	// GetPreloadedUserByID retrieves a user record with all associations preloaded by its ID.
	// It takes an unsigned integer `id` as input and returns a pointer to a `models.User` instance and an error.
	GetPreloadedUserByID(id uint) (*models.User, error)

	// Create inserts a new user record into the database.
	// It takes a pointer to a `models.User` instance as input and returns an error.
	Create(user *models.User) error

	// Update updates an existing user record in the database.
	// It takes a pointer to a `models.User` instance as input and returns an error.
	Update(user *models.User) error

	// Delete removes a user record from the database by its ID.
	// It takes an unsigned integer `id` as input and returns an error.
	Delete(id uint) error

	// GetUsersByCompanyID retrieves all users associated with a specific company ID.
	// It takes an unsigned integer `companyID` as input and returns a slice of pointers to `models.User` instances and an error.
	GetUsersByCompanyID(companyID uint) ([]*models.User, error)

	// GetCountByCompanyID returns the count of users associated with a specific company ID.
	// It takes an unsigned integer `companyID` as input and returns an integer count and an error.
	GetCountByCompanyID(companyID uint) (int64, error)

	// GetByRoleIDAndCompanyID retrieves a user record by role ID and company ID.
	// It takes two unsigned integers `roleID` and `companyID` as input and returns a pointer to a `models.User` instance and an error.
	GetByRoleIDAndCompanyID(roleID, companyID uint) (*models.User, error)

	// GetByEmail retrieves a user record from the database by its email.
	// It takes a string `email` as input and returns a pointer to a `models.User` instance and an error.
	GetByEmail(email string) (*models.User, error)
}

// NewUserRepository creates a new instance of UserRepository with the provided database connection.
// It takes a *gorm.DB as an argument, which represents the database connection, and returns a pointer to a UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

// GetByID retrieves a user record from the database by its ID.
// It takes an unsigned integer `id` as input and returns a pointer to a
// `models.User` instance and an error. If the user with the specified
// ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.Database.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// It takes an unsigned integer `id` as input and returns a pointer to a
// `models.User` instance and an error. If the user with the specified
// ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserRepository) GetPreloadedUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.Database.Preload(clause.Associations).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create inserts a new user record into the database.
// It takes a pointer to a `models.User` instance as input and returns an error.
// If the create operation fails, it returns a non-nil error.
func (r *UserRepository) Create(user *models.User) error {
	err := r.Database.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing user record in the database.
// It takes a pointer to a `models.User` instance as input and returns an error.
// If the update operation fails, it returns a non-nil error.
func (r *UserRepository) Update(user *models.User) error {
	err := r.Database.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a user record from the database by its ID.
// It takes an unsigned integer `id` as input and returns an error.
// If the user with the specified ID is not found or if the delete operation fails, it returns a non-nil error.
func (r *UserRepository) Delete(id uint) error {
	var user models.User
	err := r.Database.First(&user, id).Error
	if err != nil {
		return err
	}
	err = r.Database.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// It takes an unsigned integer `companyID` as input and returns a slice of pointers to `models.User` instances and an error.
// If there are no users associated with the specified company ID or if there is a database error, it returns a non-nil error.
func (r *UserRepository) GetUsersByCompanyID(companyID uint) ([]*models.User, error) {
	var users []*models.User
	err := r.Database.Where("company_id = ?", companyID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetCountByCompanyID returns the count of users associated with a specific company ID.
// It takes an unsigned integer `companyID` as input and returns an integer count and an error.
// If there is a database error, it returns a non-nil error.
func (r *UserRepository) GetCountByCompanyID(companyID uint) (int64, error) {
	var count int64
	err := r.Database.Model(&models.User{}).Where("company_id = ?", companyID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// It takes two unsigned integers `roleID` and `companyID` as input and returns a pointer to a `models.User` instance and an error.
// If the user with the specified role ID and company ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserRepository) GetByRoleIDAndCompanyID(roleID, companyID uint) (*models.User, error) {
	var user models.User
	err := r.Database.Where("role_id = ? AND company_id = ?", roleID, companyID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user record from the database by its email.
// It takes a string `email` as input and returns a pointer to a `models.User` instance and an error.
// If the user with the specified email is not found or if there is a database error, it returns a non-nil error.
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.Database.Where("email LIKE ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
