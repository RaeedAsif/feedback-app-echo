package utils

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword function to hash password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword function to verify hashed password
func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func IsEmailValid(email string) bool {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return addr.Address == email
}

func IsValidPasswordCheck(password string) error {
	if len(password) < 8 {
		return errors.New("password is less than 8 characters")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`)
	if !hasUpper.MatchString(password) {
		return errors.New("password should have at least one uppercase letter")
	}

	hasLower := regexp.MustCompile(`[a-z]`)
	if !hasLower.MatchString(password) {
		return errors.New("password should have at least one lowercase letter")

	}

	hasDigit := regexp.MustCompile(`\d`)
	if !hasDigit.MatchString(password) {
		return errors.New("password should have at least one digit")

	}

	hasSpecial := regexp.MustCompile(`[!@#\$%^&*()]`)
	if !hasSpecial.MatchString(password) {
		return errors.New("password should have at least one special character")
	}

	return nil
}
