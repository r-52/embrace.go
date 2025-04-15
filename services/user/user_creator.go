package user

import (
	"errors"

	"github.com/r-52/embrace/models"
	users "github.com/r-52/embrace/models/dto/user"
	"github.com/r-52/embrace/repositories"
	"gorm.io/gorm"
)

type UserCreator struct {
	userRepository *repositories.UserRepository
}

func NewUserCreator(db *gorm.DB) *UserCreator {
	return &UserCreator{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (userCreator *UserCreator) CreateUser(req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	// TODO: validate struct
	maybeUser, err := userCreator.userRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if maybeUser != nil {
		return nil, errors.New("E1000")
	}

	passwordService := NewPasswordService(req.Password)
	hashedPassword, err := passwordService.HashPassword()
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		CompanyID: req.CompanyID,
		UserProfile: models.UserProfile{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Phone:     req.Phone,
			Location:  req.Location,
			Title:     req.Title,
			Position:  req.Position,
		},
		Role: models.UserRole{
			InternalUsage: 1,
			Name:          "admin",
			CompanyID:     req.CompanyID,
		},
	}

	err = userCreator.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	return &users.CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
