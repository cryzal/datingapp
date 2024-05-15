package user

import "datingapp/core/entities"

type UserAdapter interface {
	Insert(userEntity *entities.User) error
	Get(email string) (*entities.User, error)
}
