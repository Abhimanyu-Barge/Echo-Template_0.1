package login

import (
	"net/http"
	"server/src/utils"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	token, _ := utils.GenerateJWT("Abhimanyu Barge")
	return c.JSON(http.StatusOK, token)
}
