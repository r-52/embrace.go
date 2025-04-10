package repositories_test

import (
	"errors"
	"testing"

	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

// setupUserTestDB initializes the database for testing using the common setup method.
func setupUserTestDB(t *testing.T) *gorm.DB {
	db := GetDatabase() // Use the method from common_test.go

	// Auto-migrate the User model
	err := db.AutoMigrate(&models.User{}, &models.Company{}, &models.UserRole{}, &models.TimeEntryType{}, &models.UserProfile{}, &models.TimeEntry{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestUserRepository_GetByID(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert a test user
	user := &models.User{Email: "test@example.com"}
	db.Create(user)

	// Test retrieving the user by ID
	result, err := repo.GetByID(user.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != user.ID || result.Email != user.Email {
		t.Errorf("expected %v, got %v", user, result)
	}

	// Test retrieving a non-existent user
	_, err = repo.GetByID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserRepository_Create(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Test creating a new user
	user := &models.User{Email: "new@example.com"}
	err := repo.Create(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the user was created
	var result models.User
	db.First(&result, user.ID)
	if result.Email != user.Email {
		t.Errorf("expected %v, got %v", user.Email, result.Email)
	}
}

func TestUserRepository_Update(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert a test user
	user := &models.User{Email: "old@example.com"}
	db.Create(user)

	// Update the user
	user.Email = "updated@example.com"
	err := repo.Update(user)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the update
	var result models.User
	db.First(&result, user.ID)
	if result.Email != "updated@example.com" {
		t.Errorf("expected %v, got %v", "updated@example.com", result.Email)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert a test user
	user := &models.User{Email: "delete@example.com"}
	db.Create(user)

	// Delete the user
	err := repo.Delete(user.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the user was deleted
	var result models.User
	err = db.First(&result, user.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserRepository_GetUsersByCompanyID(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert test users
	user1 := &models.User{CompanyID: 1, Email: "user1@example.com"}
	user2 := &models.User{CompanyID: 1, Email: "user2@example.com"}
	db.Create(user1)
	db.Create(user2)

	// Test retrieving users by company ID
	results, err := repo.GetUsersByCompanyID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 users, got %d", len(results))
	}
}

func TestUserRepository_GetByRoleIDAndCompanyID(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert a test user
	user := &models.User{RoleID: 1, CompanyID: 1, Email: "role@example.com"}
	db.Create(user)

	// Test retrieving the user by role ID and company ID
	result, err := repo.GetByRoleIDAndCompanyID(1, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != user.ID || result.Email != user.Email {
		t.Errorf("expected %v, got %v", user, result)
	}

	// Test retrieving a non-existent user
	_, err = repo.GetByRoleIDAndCompanyID(2, 1)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert a test user
	user := &models.User{Email: "email@example.com"}
	db.Create(user)

	// Test retrieving the user by email
	result, err := repo.GetByEmail("email@example.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != user.ID || result.Email != user.Email {
		t.Errorf("expected %v, got %v", user, result)
	}

	// Test retrieving a non-existent user
	_, err = repo.GetByEmail("nonexistent@example.com")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserRepository_GetCountByCompanyID(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert test users
	user1 := &models.User{CompanyID: 1, Email: "user1@example.com"}
	user2 := &models.User{CompanyID: 1, Email: "user2@example.com"}
	user3 := &models.User{CompanyID: 2, Email: "user3@example.com"}
	db.Create(user1)
	db.Create(user2)
	db.Create(user3)

	// Test count for company ID 1
	count, err := repo.GetCountByCompanyID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}

	// Test count for company ID 2
	count, err = repo.GetCountByCompanyID(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("expected count 1, got %d", count)
	}

	// Test count for a non-existent company ID
	count, err = repo.GetCountByCompanyID(999)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}
}
func TestUserRepository_GetPreloadedUserByID(t *testing.T) {
	db := setupUserTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Insert a test user with associations
	user := &models.User{
		Email: "preloaded@example.com",
		// Add any associations here, e.g., Profile, Roles, etc., if applicable
	}
	db.Create(user)

	// Test retrieving the user with preloaded associations by ID
	result, err := repo.GetPreloadedUserByID(user.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != user.ID || result.Email != user.Email {
		t.Errorf("expected %v, got %v", user, result)
	}

	// Test retrieving a non-existent user
	_, err = repo.GetPreloadedUserByID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}
