package app

import (
	"datingapp/core/service"
	"datingapp/infrastructure"
	"datingapp/interface/rest/user"
	"datingapp/shared/config"
	"datingapp/shared/driver"
	"datingapp/shared/protocol/rest"
)

type AppUser struct {
	httpHandler *rest.EchoHTTPHandler
	route       driver.Router
}

func (c AppUser) RunApplication() {
	c.route.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewUser() func() driver.RegistryContract {
	return func() driver.RegistryContract {
		cfg := config.ReadConfig("APP_USER_ADDRESS")

		httpHandler := rest.NewEchoHTTPHandlerDefault(cfg)
		datasource := infrastructure.NewUserGateway(cfg)

		return &AppUser{
			httpHandler: &httpHandler,
			route: &user.Routes{
				HTTPHandler: httpHandler,
				Config:      cfg,
				Port:        service.UserServiceNew(datasource.UserAdapter),
			},
		}

	}
}
