package repository

import (
	"datingapp/core/entities"
	"errors"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) Insert(userEntity *entities.User) error {

	return nil
}

func (_m *UserRepository) Get(email string) (*entities.User, error) {
	if email == "agung@gmail.com" {
		return &entities.User{
			ID:       "1",
			Email:    "agung@gmail.com",
			Password: "$2a$10$zC.TZJgzB514/AvQ9elfleLz7CcvTYx65UudQwfDKL5aXsOkyHRRq",
		}, nil
	}

	return nil, errors.New("not found")
}
