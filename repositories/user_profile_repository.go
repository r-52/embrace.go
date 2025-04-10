package repositories

import (
	"github.com/r-52/embrace/models"
	"gorm.io/gorm"
)

type UserProfileRepository struct {
	Database *gorm.DB
}

// NewUserProfileRepository creates a new instance of UserProfileRepository with the provided database connection.
// It takes a *gorm.DB as an argument, which represents the database connection, and returns a pointer to a UserProfileRepository.
func NewUserProfileRepository(db *gorm.DB) *UserProfileRepository {
	return &UserProfileRepository{
		Database: db,
	}
}

// GetByID retrieves a user profile record from the database by its ID.
// It takes an unsigned integer `id` as input and returns a pointer to a
// `models.UserProfile` instance and an error. If the user profile with the specified
// ID is not found or if there is a database error, it returns a non-nil error.
func (r *UserProfileRepository) GetByID(id uint) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.Database.First(&userProfile, id).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

// Create inserts a new user profile record into the database.
// It takes a pointer to a `models.UserProfile` instance as input and returns an error.
// If the create operation fails, it returns a non-nil error.
func (r *UserProfileRepository) Create(userProfile *models.UserProfile) error {
	err := r.Database.Create(userProfile).Error
	if err != nil {
		return err
	}
	return nil
}

// Update updates an existing user profile record in the database.
// It takes a pointer to a `models.UserProfile` instance as input and returns an error.
// If the update operation fails, it returns a non-nil error.
func (r *UserProfileRepository) Update(userProfile *models.UserProfile) error {
	err := r.Database.Save(userProfile).Error
	if err != nil {
		return err
	}
	return nil
}

// It takes a string `slug` as input and returns a pointer to a
// `models.UserProfile` instance and an error. If the user profile with the specified
// slug is not found or if there is a database error, it returns a non-nil error.
func (r *UserProfileRepository) GetBySlug(slug string) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.Database.Where("slug = ?", slug).First(&userProfile).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}
