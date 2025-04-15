package user_test

import (
	"testing"

	"github.com/r-52/embrace/services/user"
)

func TestPasswordService_HashPassword(t *testing.T) {
	password := "securepassword123"
	passwordService := user.NewPasswordService(password)

	// Teste die HashPassword-Methode
	hash, err := passwordService.HashPassword()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if hash == "" {
		t.Fatalf("expected a non-empty hash, got an empty string")
	}
}

func TestPasswordService_ComparePassword(t *testing.T) {
	password := "securepassword123"
	passwordService := user.NewPasswordService(password)
	hash, err := passwordService.HashPassword()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	match, err := passwordService.ComparePassword(hash)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !match {
		t.Fatalf("expected passwords to match, but they did not")
	}
}
