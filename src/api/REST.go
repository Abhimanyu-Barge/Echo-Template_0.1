package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadREST() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo server Up!!! with auto reload ")

	})
}
