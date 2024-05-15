package user

import (
	port "datingapp/core/port/user"
	"datingapp/interface/rest/user/handler"
	"datingapp/shared/config"
	"datingapp/shared/protocol/rest"
)

type Routes struct {
	HTTPHandler rest.EchoHTTPHandler
	Config      *config.Config
	Port        port.Service
}

func (r *Routes) RegisterRouter() {
	userHandler := handler.New(r.Port)

	userRoute := r.HTTPHandler.Framework.Group("/user")
	userRoute.POST("", userHandler.Create)
	userRoute.POST("/login", userHandler.Login)

}
