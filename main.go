package main

import (
	"github.com/Splucheviy/akhilsharmaEchoLesson/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health-check", func(c echo.Context) error {
		return handlers.HealthCheckHandler(c)
	})
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id",handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
