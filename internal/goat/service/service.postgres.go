package service

import (
	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/bontusss/goat/internal/goat/repository"
	"github.com/bontusss/goat/internal/goat/utils"
)

type PostgresServiceImpl struct {
	postgresRepository repository.PostgreSQLUserRepository
}

func NewPostgreSQLUserRepository(repo repository.PostgreSQLUserRepository) goat.UserService {
	return &PostgresServiceImpl{postgresRepository: repo}
}

// DeleteUser implements goat.UserService.
func (p *PostgresServiceImpl) DeleteUser(id uint) error {
	if err := p.postgresRepository.DeleteUser(id); err != nil {
		return err
	}

	return nil
}

// GetUserByID implements goat.UserService.
func (p *PostgresServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := p.postgresRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements goat.UserService.
func (p *PostgresServiceImpl) Login(email string, password string) (*models.User, error) {
	user, err := p.postgresRepository.Login(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Register implements goat.UserService.
func (p *PostgresServiceImpl) Register(user *models.User) error {
	if err := utils.ValidateUser(user); err != nil {
		return err
	}

	if err := p.postgresRepository.Register(user); err != nil {
		return err
	}

	return nil
}

// ResetPassword implements goat.UserService.
func (p *PostgresServiceImpl) ResetPassword(email string, newPassword string) error {
	if err := p.postgresRepository.ResetPassword(email, newPassword); err != nil {
		return err
	}

	return nil
}

// UpdateUser implements goat.UserService.
func (p *PostgresServiceImpl) UpdateUser(user *models.User) error {
	if err := p.postgresRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
