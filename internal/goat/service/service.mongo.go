package service

import (
	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/bontusss/goat/internal/goat/repository"
	"github.com/bontusss/goat/internal/goat/utils" // Add the correct import path for utils
)

type MongoServiceImpl struct {
	mongoRepository repository.MongoDBUserRepository
}

func NewMongoService(repo repository.MongoDBUserRepository) goat.UserService {
	return &MongoServiceImpl{repo}
}

func (s *MongoServiceImpl) Register(user *models.User) error {
	// validate user data
	if err := utils.ValidateUser(user); err != nil {
		return err
	}
	if err := s.mongoRepository.Register(user); err != nil {
		return err
	}

	return nil
}

// DeleteUser implements goat.UserService.
func (s *MongoServiceImpl) DeleteUser(id uint) error {
	if err := s.mongoRepository.DeleteUser(id); err != nil {
		return err
	}

	return nil
}

// GetUserByID implements goat.UserService.
func (s *MongoServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := s.mongoRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements goat.UserService.
func (s *MongoServiceImpl) Login(email string, password string) (*models.User, error) {
	user, err := s.mongoRepository.Login(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ResetPassword implements goat.UserService.
func (s *MongoServiceImpl) ResetPassword(email string, newPassword string) error {
	if err := s.mongoRepository.ResetPassword(email, newPassword); err != nil {
		return err
	}

	return nil
}

// UpdateUser implements goat.UserService.
func (s *MongoServiceImpl) UpdateUser(user *models.User) error {
	if err := s.mongoRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
