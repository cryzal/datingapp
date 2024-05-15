package handler

import (
	"datingapp/core/entities"
	port "datingapp/core/port/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service port.Service
}

func New(service port.Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) Create(c echo.Context) error {
	defaultResponse := DefaultResponse{}
	request := CreatePayload{}
	err := c.Bind(&request)
	if err != nil {
		defaultResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, defaultResponse)
	}

	userEntity := entities.User{
		Email:    request.Email,
		Password: request.Password,
	}

	err = h.service.Create(&userEntity)
	if err != nil {
		defaultResponse.Message = err.Error()
		return c.JSON(400, defaultResponse)
	}

	return c.JSON(201, nil)
}

func (h *Handler) Login(c echo.Context) error {
	defaultResponse := DefaultResponse{}
	request := LoginPayload{}
	response := LoginResponse{}
	err := c.Bind(&request)
	if err != nil {
		defaultResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, nil)
	}

	userEntity, err := h.service.Login(request.Email, request.Password)
	if err != nil {
		defaultResponse.Message = err.Error()
		return c.JSON(400, defaultResponse)
	}
	response.Token = userEntity.Token
	defaultResponse.Message = "success"
	defaultResponse.Data = response

	return c.JSON(200, defaultResponse)
}
