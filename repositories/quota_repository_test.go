package repositories_test

import (
	"errors"
	"testing"

	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

// setupQuotaTestDB initializes the database for testing using the common setup method.
func setupQuotaTestDB(t *testing.T) *gorm.DB {
	db := GetDatabase() // Use the method from common_test.go

	// Auto-migrate the Quota model
	err := db.AutoMigrate(&models.Quota{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestQuotaRepository_GetByID(t *testing.T) {
	db := setupQuotaTestDB(t)
	repo := repositories.NewQuotaRepository(db)

	// Insert a test quota
	quota := &models.Quota{Name: "Test Quota"}
	db.Create(quota)

	// Test retrieving the quota by ID
	result, err := repo.GetByID(quota.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != quota.ID || result.Name != quota.Name {
		t.Errorf("expected %v, got %v", quota, result)
	}

	// Test retrieving a non-existent quota
	_, err = repo.GetByID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestQuotaRepository_GetByCompanyID(t *testing.T) {
	db := setupQuotaTestDB(t)
	repo := repositories.NewQuotaRepository(db)

	// Insert test quotas
	quota1 := &models.Quota{CompanyID: 1, Name: "Quota 1"}
	quota2 := &models.Quota{CompanyID: 1, Name: "Quota 2"}
	db.Create(quota1)
	db.Create(quota2)

	// Test retrieving quotas by company ID
	results, err := repo.GetByCompanyID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 quotas, got %d", len(results))
	}
}
func TestQuotaRepository_Create(t *testing.T) {
	db := setupQuotaTestDB(t)
	repo := repositories.NewQuotaRepository(db)

	// Test creating a valid quota
	quota := &models.Quota{Name: "New Quota", CompanyID: 1}
	err := repo.Create(quota)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if quota.ID == 0 {
		t.Errorf("expected quota ID to be set, got 0")
	}

	// Test creating a quota with missing required fields
	invalidQuota := &models.Quota{}
	err = repo.Create(invalidQuota)
	if err != nil {
		t.Errorf("expected error, got nil")
	}
}
