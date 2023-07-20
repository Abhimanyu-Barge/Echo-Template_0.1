package api

import (
	"log"

	"github.com/labstack/echo/v4/middleware"
)

func LoadMiddleware() {
	log.Println("Load Middleware")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
}
