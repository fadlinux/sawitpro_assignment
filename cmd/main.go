package main

import (
	"os"

	"github.com/fadlinux/sawitpro_assignment/handler"
	"github.com/fadlinux/sawitpro_assignment/repository"
	"github.com/fadlinux/sawitpro_assignment/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	//set new service dependent repo
	service.NewService(repo)

	var service = service.NewService(repo)

	NewRouter(e, service)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":1323"))
}

func NewRouter(
	e *echo.Echo,
	service service.ServiceInterface,
) {
	handle := handler.NewServer(e, service)

	r := e.Group("")
	r.POST("/estate", handle.CreateEstate)
	r.POST("/estate/tree", handle.CreateTree)
	r.GET("/estate/:id/stats", handle.GetStats)
	r.GET("/estate/:id/drone-plan", handle.GetDronePlan)
}
