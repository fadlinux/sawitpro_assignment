package handler

import (
	"github.com/fadlinux/sawitpro_assignment/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e       *echo.Echo
	Service service.ServiceInterface
}

func NewServer(
	e *echo.Echo,
	service service.ServiceInterface,
) *Server {
	return &Server{
		e:       e,
		Service: service,
	}
}
