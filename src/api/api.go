package api

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func InitEchoServer() {
	e := echo.New()
	// loading env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo server Up!!! with auto reload")
	})
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
