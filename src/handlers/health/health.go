package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthHanler(c echo.Context) error {
	return c.String(http.StatusOK, "Echo server Up!!! with auto reload ")
}
