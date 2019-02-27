package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kimihito/tohatebu/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/", handlers.Root)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.Start(":" + port)
}
