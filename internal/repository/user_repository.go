package repository

import (
	"errors"
	"hierarchy-management/internal/domain"
	"os"
)

type userRepository struct {
	users map[string]string
}

func NewUserRepository() domain.UserRepository {

	username := os.Getenv("AUTH_USERNAME")
	hashedPassword := os.Getenv("AUTH_PASSWORD")

	users := map[string]string{
		username: hashedPassword,
	}

	return &userRepository{users: users}
}

func (r *userRepository) GetUser(username string) (*domain.User, error) {
	if hashedPassword, ok := r.users[username]; ok {
		return &domain.User{
			Username: username,
			Password: hashedPassword,
		}, nil
	}
	return nil, errors.New("user not found")
}
