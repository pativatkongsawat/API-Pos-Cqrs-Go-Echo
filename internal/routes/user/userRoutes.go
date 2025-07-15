package routes

import (
	user "echo_cqrs/internal/src/user/query"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	api := e.Group("/api")

	api.GET("/user", user.ControllerGetAllUser)

}
