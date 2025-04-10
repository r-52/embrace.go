package repositories_test

import (
	"errors"
	"testing"

	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

// setupCompanyTestDB initializes an in-memory SQLite database for testing.
func setupCompanyTestDB(t *testing.T) *gorm.DB {
	db := GetDatabase()

	// Auto-migrate the Company model
	err := db.AutoMigrate(&models.Company{}, &models.User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestCompanyRepository_GetByID(t *testing.T) {
	db := setupCompanyTestDB(t)
	repo := repositories.NewCompanyRepository(db)

	// Insert a test company
	company := &models.Company{Name: "Test Company"}
	db.Create(company)

	// Test retrieving the company by ID
	result, err := repo.GetByID(company.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != company.ID || result.Name != company.Name {
		t.Errorf("expected %v, got %v", company, result)
	}

	// Test retrieving a non-existent company
	_, err = repo.GetByID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestCompanyRepository_Create(t *testing.T) {
	db := setupCompanyTestDB(t)
	repo := repositories.NewCompanyRepository(db)

	// Test creating a new company
	company := &models.Company{Name: "New Company"}
	err := repo.Create(company)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the company was created
	var result models.Company
	db.First(&result, company.ID)
	if result.Name != company.Name {
		t.Errorf("expected %v, got %v", company.Name, result.Name)
	}
}

func TestCompanyRepository_Update(t *testing.T) {
	db := setupCompanyTestDB(t)
	repo := repositories.NewCompanyRepository(db)

	// Insert a test company
	company := &models.Company{Name: "Old Name"}
	db.Create(company)

	// Update the company
	company.Name = "Updated Name"
	err := repo.Update(company)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the update
	var result models.Company
	db.First(&result, company.ID)
	if result.Name != "Updated Name" {
		t.Errorf("expected %v, got %v", "Updated Name", result.Name)
	}
}

func TestCompanyRepository_Delete(t *testing.T) {
	db := setupCompanyTestDB(t)
	repo := repositories.NewCompanyRepository(db)

	// Insert a test company
	company := &models.Company{Name: "Test Company"}
	db.Create(company)

	id := company.ID
	// Delete the company
	err := repo.Delete(company.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the company was deleted
	var result models.Company
	err = db.First(&result, id).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}

func TestCompanyRepository_GetByUserID(t *testing.T) {
	db := setupCompanyTestDB(t)
	repo := repositories.NewCompanyRepository(db)

	// Insert a test company and user
	company := &models.Company{Name: "Test Company"}
	db.Create(company)
	user := &models.User{CompanyID: 1, Email: "Test@User.com"}
	db.Create(user)

	// Test retrieving the company by user ID
	result, err := repo.GetByUserID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result.ID != company.ID || result.Name != company.Name {
		t.Errorf("expected %v, got %v", company, result)
	}

	// Test retrieving a company for a non-existent user
	_, err = repo.GetByUserID(999)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected ErrRecordNotFound, got %v", err)
	}
}
