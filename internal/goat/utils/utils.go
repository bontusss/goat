package utils

import (
	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
)

func ValidateUser(user *models.User) error {
	if user.Email == "" {
		return goat.ErrEmailNotProvided
	}

	if user.Password == "" {
		return goat.ErrPasswordNotProvided
	}

	return nil
}