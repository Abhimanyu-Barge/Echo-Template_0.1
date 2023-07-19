package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitEchoServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo server Up!!! with auto reload")
	})
	e.Logger.Fatal(e.Start(":8081"))
}
