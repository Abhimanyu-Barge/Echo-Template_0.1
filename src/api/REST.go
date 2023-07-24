package api

import (
	"server/src/handlers/health"
	"server/src/handlers/login"
)

// {{   /o/ :- is for open   }}
// {{  /r/:-  or nothing is resticted without tocken you canot access api  }}
func LoadREST() {
	e.GET("/o/", health.HealthHanler)
	e.GET("/o/login", login.LoginHandler)
}
