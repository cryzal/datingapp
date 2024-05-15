package user

import "datingapp/core/entities"

type Service interface {
	Create(user *entities.User) error
	Login(email, password string) (*entities.User, error)
}
