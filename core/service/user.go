package service

import (
	"datingapp/core/entities"
	port "datingapp/core/port/user"
	"datingapp/shared/utils/jwthelper"
	"errors"
)

type UserService struct {
	UserRepositoryAdapter port.UserAdapter
}

func UserServiceNew(userAdapterPort port.UserAdapter) *UserService {
	return &UserService{userAdapterPort}
}

func (p *UserService) Create(user *entities.User) error {
	user.HashPassword()

	err := p.UserRepositoryAdapter.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserService) Login(email, password string) (*entities.User, error) {
	user, err := p.UserRepositoryAdapter.Get(email)
	if err != nil {
		return nil, err
	}

	isValidPassword := user.CheckPasswordHash(password)
	if !isValidPassword {
		return nil, errors.New("Password Invalid")
	}

	credentials := make([]string, 0)
	additionalParam := make(map[string]interface{}, 0)
	tokenData := jwthelper.Login{
		ID:    user.ID,
		Email: user.Email,
	}
	TokenData, err := jwthelper.GenerateNewToken(tokenData, credentials, additionalParam)
	if err != nil {
		return nil, err
	}

	user.SetToken(TokenData.Access)

	return user, nil
}
