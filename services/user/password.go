package user

import "github.com/alexedwards/argon2id"

type PasswordService struct {
	password string
}

func NewPasswordService(password string) *PasswordService {
	return &PasswordService{
		password: password,
	}
}

// HashPassword generates a secure hash of the user's password using the Argon2id algorithm.
// It utilizes the default parameters provided by the argon2id package to ensure a balance
// between security and performance. The function returns the hashed password as a string
// and an error if the hashing process fails.
func (p *PasswordService) HashPassword() (string, error) {
	hash, err := argon2id.CreateHash(p.password, argon2id.DefaultParams)
	return hash, err
}

// ComparePassword compares a plain text password with a hashed password to determine if they match.
// It uses the Argon2id algorithm for secure password comparison.
//
// Parameters:
//   - hashedPassword: The hashed password to compare against.
//
// Returns:
//   - bool: True if the passwords match, false otherwise.
//   - error: An error if the comparison fails or encounters an issue.
func (p *PasswordService) ComparePassword(hashedPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(p.password, hashedPassword)
	return match, err
}
