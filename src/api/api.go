package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func InitEchoServer() {
	// loading env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Loding middleware
	LoadMiddleware()

	//Loading REST api
	LoadREST()

	//Enable Port
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
