package service

import (
	"hierarchy-management/internal/domain"
	"hierarchy-management/internal/errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Authenticate(username, password string) error
}

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Authenticate(username, password string) error {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return errors.NewAuthenticationError("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.NewAuthenticationError("invalid username or password")
	}

	return nil
}
