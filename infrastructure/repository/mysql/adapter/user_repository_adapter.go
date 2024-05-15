package adapter

import (
	"datingapp/core/entities"
	repository "datingapp/infrastructure/repository/mysql"
	"datingapp/infrastructure/repository/mysql/models"
	"strconv"
)

type UserRepositoryAdapter struct {
	UserRepo *repository.RepositoryUser
}

func NewuserRepositoryAdapter(repo *repository.RepositoryUser) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{UserRepo: repo}
}

func (a UserRepositoryAdapter) Insert(userEntity *entities.User) error {
	user := models.UserModel{
		Email:    userEntity.Email,
		Password: userEntity.Password,
	}
	err := a.UserRepo.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func (a UserRepositoryAdapter) Get(email string) (*entities.User, error) {
	userModel, err := a.UserRepo.Get(email)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       strconv.FormatInt(*userModel.ID, 10),
		Email:    userModel.Email,
		Password: userModel.Password,
	}, nil
}
