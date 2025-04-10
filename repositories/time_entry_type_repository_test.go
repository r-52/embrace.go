package repositories_test

import (
	"errors"
	"testing"

	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&models.TimeEntryType{}); err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}
	return db
}

func TestCountByCompanyID(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewTimeEntryTypeRepository(db)

	companyID := uint(1)
	timeEntryTypes := []models.TimeEntryType{
		{CompanyID: companyID, Name: "Type 1"},
		{CompanyID: companyID, Name: "Type 2"},
		{CompanyID: 2, Name: "Type 3"},
	}
	for _, x := range timeEntryTypes {
		if err := db.Create(&x).Error; err != nil {
			t.Fatalf("failed to seed database: %v", err)
		}
	}

	count, err := repo.CountByCompanyID(companyID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}

	emptyCompanyID := uint(3)
	count, err = repo.CountByCompanyID(emptyCompanyID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}
}

func TestGetByCompanyID(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewTimeEntryTypeRepository(db)

	companyID := uint(1)
	timeEntryTypes := []models.TimeEntryType{
		{CompanyID: companyID, Name: "Type 1"},
		{CompanyID: companyID, Name: "Type 2"},
		{CompanyID: 2, Name: "Type 3"},
	}
	for _, x := range timeEntryTypes {
		if err := db.Create(&x).Error; err != nil {
			t.Fatalf("failed to seed database: %v", err)
		}
	}

	result, err := repo.GetByCompanyID(companyID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 results, got %d", len(result))
	}
	if result[0].Name != "Type 1" || result[1].Name != "Type 2" {
		t.Errorf("unexpected result: %+v", result)
	}

	emptyCompanyID := uint(3)
	result, err = repo.GetByCompanyID(emptyCompanyID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected 0 results, got %d", len(result))
	}
}

func TestDelete(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewTimeEntryTypeRepository(db)

	timeEntryType := models.TimeEntryType{CompanyID: 1, Name: "Type 1"}
	if err := db.Create(&timeEntryType).Error; err != nil {
		t.Fatalf("failed to seed database: %v", err)
	}

	if err := repo.Delete(timeEntryType.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result models.TimeEntryType
	err := db.First(&result, timeEntryType.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record not found error, got %v", err)
	}

	nonExistentID := uint(999)
	err = repo.Delete(nonExistentID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record not found error, got %v", err)
	}
}

func TestUpdate(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewTimeEntryTypeRepository(db)

	timeEntryType := models.TimeEntryType{CompanyID: 1, Name: "Original Name"}
	if err := db.Create(&timeEntryType).Error; err != nil {
		t.Fatalf("failed to seed database: %v", err)
	}

	timeEntryType.Name = "Updated Name"
	if err := repo.Update(&timeEntryType); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result models.TimeEntryType
	if err := db.First(&result, timeEntryType.ID).Error; err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "Updated Name" {
		t.Errorf("expected name 'Updated Name', got '%s'", result.Name)
	}
}

func TestCreate(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewTimeEntryTypeRepository(db)

	timeEntryType := models.TimeEntryType{CompanyID: 1, Name: "New Type"}
	if err := repo.Create(&timeEntryType); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if timeEntryType.ID == 0 {
		t.Fatalf("expected ID to be set, got 0")
	}

	var result models.TimeEntryType
	if err := db.First(&result, timeEntryType.ID).Error; err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "New Type" || result.CompanyID != 1 {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestGetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewTimeEntryTypeRepository(db)

	timeEntryType := models.TimeEntryType{CompanyID: 1, Name: "Test Type"}
	if err := db.Create(&timeEntryType).Error; err != nil {
		t.Fatalf("failed to seed database: %v", err)
	}

	result, err := repo.GetByID(timeEntryType.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil || result.ID != timeEntryType.ID || result.Name != "Test Type" {
		t.Errorf("unexpected result: %+v", result)
	}

	nonExistentID := uint(999)
	result, err = repo.GetByID(nonExistentID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record not found error, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result, got %+v", result)
	}
}
