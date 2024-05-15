package infrastructure

import (
	"datingapp/infrastructure/repository/mysql"
	repository "datingapp/infrastructure/repository/mysql"
	"datingapp/infrastructure/repository/mysql/adapter"
	"datingapp/shared/config"
)

type UserGateway struct {
	UserAdapter *adapter.UserRepositoryAdapter
}

func NewUserGateway(cfg *config.Config) *UserGateway {
	/// example mysql
	db := repository.Connect(cfg)

	userRepository := mysql.Newuser(db)
	return &UserGateway{
		UserAdapter: adapter.NewuserRepositoryAdapter(userRepository),
	}
}
