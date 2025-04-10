package repositories_test

import (
	"errors"
	"testing"

	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

// setupUserRoleTestDB initializes the database for testing using the common setup method.
func setupUserRoleTestDB(t *testing.T) *gorm.DB {
	db := GetDatabase() // Use the method from common_test.go

	// Auto-migrate the UserRole model
	err := db.AutoMigrate(&models.UserRole{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestUserRoleRepository_GetByID(t *testing.T) {
	db := setupUserRoleTestDB(t)
	repo := repositories.NewUserRoleRepository(db)

	// Insert a test user role
	userRole := &models.UserRole{Name: "Admin"}
	db.Create(userRole)

	// Test retrieving the user role by ID
	result, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != userRole.ID || result.Name != userRole.Name {
		t.Errorf("expected %v, got %v", userRole, result)
	}

	// Test retrieving a non-existent user role
	_, err = repo.GetByID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserRoleRepository_Create(t *testing.T) {
	db := setupUserRoleTestDB(t)
	repo := repositories.NewUserRoleRepository(db)

	// Test creating a new user role
	userRole := &models.UserRole{Name: "Editor"}
	err := repo.Create(userRole)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
