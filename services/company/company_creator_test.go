package company_test

import (
	"testing"

	"github.com/r-52/embrace/models"
	dto "github.com/r-52/embrace/models/dto/company"
	"github.com/r-52/embrace/models/dto/user"
	"github.com/r-52/embrace/repositories"
	srv "github.com/r-52/embrace/services/company"
	"gorm.io/gorm"
)

func setupDb() *gorm.DB {
	db := repositories.GetDatabase()
	db.AutoMigrate(&models.Company{}, &models.User{}, &models.UserProfile{}, &models.UserRole{})
	return db
}

func TestCompanyCreator_Create_Company_With_Success(t *testing.T) {
	db := setupDb()
	companyCreator := srv.NewCompanyCreator(db)
	company, err := companyCreator.CreateCompany(&dto.CreateCompanyRequest{
		Name:        "Test Company",
		Description: "Test Description",
		Website:     "https://test.com",
		User: &user.CreateUserRequest{
			Email:     "test@test.com",
			Password:  "password",
			FirstName: "Test",
			LastName:  "User",
			Phone:     "1234567890",
			Location:  "Test Location",
			Title:     "Test Title",
			Position:  "Test Position",
		},
	})

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if company == nil {
		t.Errorf("expected company to be created, got nil")
	}
}
