package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	// Start as a web server
	if err := e.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
