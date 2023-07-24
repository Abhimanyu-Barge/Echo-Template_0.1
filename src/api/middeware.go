package api

import (
	"fmt"
	"log"
	"net/http"
	"server/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadMiddleware() {
	log.Println("Load Middleware")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(CoustomJWTToken)
}

func CoustomJWTToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		fmt.Println(r.RequestURI)
		if utils.IsOpenUrl(r.RequestURI) {
			return next(c)
		}
		Jwttocken := r.Header["Authorization"]
		if len(Jwttocken) >= 1 {
			fmt.Println(Jwttocken[0])
			validateJwt, err := utils.ValidateToken(Jwttocken[0])
			if validateJwt {
				return next(c)
			}
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Authorization header")
	}
}
