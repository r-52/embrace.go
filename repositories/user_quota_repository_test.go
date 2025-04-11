package repositories_test

import (
	"errors"
	"testing"

	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

// setupUserQuotaDB initializes the database for testing using the common setup method.
func setupUserQuotaDB(t *testing.T) *gorm.DB {
	db := GetDatabase() // Use the method from common_test.go

	// Auto-migrate the UserQuota model
	err := db.AutoMigrate(&models.UserQuota{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestUserQuotaRepository_GetByID(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Insert a test UserQuota
	userQuota := &models.UserQuota{UserID: 1, QuotaID: 1}
	db.Create(userQuota)

	// Test retrieving the UserQuota by ID
	result, err := repo.GetByID(userQuota.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != userQuota.ID || result.UserID != userQuota.UserID || result.QuotaID != userQuota.QuotaID {
		t.Errorf("expected %v, got %v", userQuota, result)
	}

	// Test retrieving a non-existent UserQuota
	_, err = repo.GetByID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserQuotaRepository_GetByUserID(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Insert test UserQuotas
	userQuota1 := &models.UserQuota{UserID: 1, QuotaID: 1}
	userQuota2 := &models.UserQuota{UserID: 1, QuotaID: 2}
	db.Create(userQuota1)
	db.Create(userQuota2)

	// Test retrieving UserQuotas by UserID
	results, err := repo.GetByUserID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 UserQuotas, got %d", len(results))
	}
}

func TestUserQuotaRepository_Create(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Test creating a new UserQuota
	userQuota := &models.UserQuota{UserID: 1, QuotaID: 1}
	err := repo.Create(userQuota)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the UserQuota was created
	var result models.UserQuota
	db.First(&result, userQuota.ID)
	if result.UserID != userQuota.UserID || result.QuotaID != userQuota.QuotaID {
		t.Errorf("expected %v, got %v", userQuota, result)
	}
}

func TestUserQuotaRepository_Update(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Insert a test UserQuota
	userQuota := &models.UserQuota{UserID: 1, QuotaID: 1}
	db.Create(userQuota)

	// Update the UserQuota
	userQuota.QuotaID = 2
	err := repo.Update(userQuota)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the update
	var result models.UserQuota
	db.First(&result, userQuota.ID)
	if result.QuotaID != 2 {
		t.Errorf("expected QuotaID 2, got %d", result.QuotaID)
	}
}

func TestUserQuotaRepository_Delete(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Insert a test UserQuota
	userQuota := &models.UserQuota{UserID: 1, QuotaID: 1}
	db.Create(userQuota)

	// Delete the UserQuota
	err := repo.Delete(userQuota.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the UserQuota was deleted
	var result models.UserQuota
	err = db.First(&result, userQuota.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestUserQuotaRepository_GetByUserIDAndQuotaID(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Insert a test UserQuota
	userQuota := &models.UserQuota{UserID: 1, QuotaID: 1}
	db.Create(userQuota)

	// Test retrieving the UserQuota by UserID and QuotaID
	result, err := repo.GetByUserIDAndQuotaID(1, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != userQuota.ID || result.UserID != userQuota.UserID || result.QuotaID != userQuota.QuotaID {
		t.Errorf("expected %v, got %v", userQuota, result)
	}
}

func TestUserQuotaRepository_CountByUserID(t *testing.T) {
	db := setupUserQuotaDB(t)
	repo := repositories.UserQuotaRepository{Database: db}

	// Insert test UserQuotas
	userQuota1 := &models.UserQuota{UserID: 1, QuotaID: 1}
	userQuota2 := &models.UserQuota{UserID: 1, QuotaID: 2}
	db.Create(userQuota1)
	db.Create(userQuota2)

	// Test counting UserQuotas by UserID
	count, err := repo.CountByUserID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}
