package service

import (
	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/bontusss/goat/internal/goat/repository"
	"github.com/bontusss/goat/internal/goat/utils"
)

type MysqlServiceImpl struct {
	MysqlRepository repository.MySQLUserRepository
}

func NewMysqlService(repo repository.MySQLUserRepository) goat.UserService {
	return &MysqlServiceImpl{repo}
}

// DeleteUser implements goat.UserService.
func (m *MysqlServiceImpl) DeleteUser(id uint) error {
	if err := m.MysqlRepository.DeleteUser(id); err != nil {
		return err
	}

	return nil
}

// GetUserByID implements goat.UserService.
func (m *MysqlServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := m.MysqlRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements goat.UserService.
func (m *MysqlServiceImpl) Login(email string, password string) (*models.User, error) {
	user, err := m.MysqlRepository.Login(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Register implements goat.UserService.
func (m *MysqlServiceImpl) Register(user *models.User) error {
	if err := utils.ValidateUser(user); err != nil {
		return err
	}

	if err := m.MysqlRepository.Register(user); err != nil {
		return err
	}

	return nil
}

// ResetPassword implements goat.UserService.
func (m *MysqlServiceImpl) ResetPassword(email string, newPassword string) error {
	if err := m.MysqlRepository.ResetPassword(email, newPassword); err != nil {
		return err
	}

	return nil
}

// UpdateUser implements goat.UserService.
func (m *MysqlServiceImpl) UpdateUser(user *models.User) error {
	if err := m.MysqlRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
