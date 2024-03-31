package goat

import (
	"errors"
)

// This code defines several custom error types specific to authentication failures.
var (
	// common errors
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrUserNotFound        = errors.New("user not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
	ErrEmailNotProvided    = errors.New("email is required")
	ErrPasswordNotProvided = errors.New("password not provided")

	// Potential additional errors (you can add more as needed)
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("token expired")
	ErrMissingToken     = errors.New("missing token")
	ErrPermissionDenied = errors.New("permission denied")
)
