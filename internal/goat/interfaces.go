package goat

import (
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/gin-gonic/gin"
)

// UserService defines the interface for user management
type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error) // Get user by ID
	UpdateUser(user *models.User) error        // Update user information
	DeleteUser(id uint) error                  // Delete a user (consider security implications)
	ResetPassword(email, newPassword string) error
	// You can add more methods as needed (e.g., search users)
}

// Authenticator defines the interface for authentication methods
type Authenticator interface {
	Authenticate(c *gin.Context) (*models.User, error)
	// Additional methods for different authentication flows (optional)
	RefreshAuthToken(c *gin.Context) (*models.User, error)             // Refresh JWT token (if using JWT)
	SocialLogin(c *gin.Context, provider string) (*models.User, error) // Social login using providers (optional)
	// You can add more methods for specific authentication flows
}
