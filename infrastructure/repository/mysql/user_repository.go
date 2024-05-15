package mysql

import (
	"datingapp/infrastructure/repository/mysql/models"

	"gorm.io/gorm"
)

type (
	RepositoryUser struct {
		*gorm.DB
	}
)

func Newuser(db *gorm.DB) *RepositoryUser {
	return &RepositoryUser{db}
}

func (p *RepositoryUser) Insert(product *models.UserModel) error {
	store := p.Create(product)

	if err := store.Error; err != nil {
		return err
	}
	return nil
}

func (p *RepositoryUser) Get(email string) (models.UserModel, error) {
	user := models.UserModel{}
	store := p.Take(&user)

	if err := store.Error; err != nil {
		return user, err
	}
	return user, nil
}
