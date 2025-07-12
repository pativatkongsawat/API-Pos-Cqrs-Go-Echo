package user

import (
	"echo_cqrs/internal/database"
	"echo_cqrs/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ControllerGetAllUser(ctx echo.Context) error {

	userService := DatabaseRequest{DB: database.DBMYSQL}

	user, err := userService.GetAllUser()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.ResponseMessage{
			Status:  500,
			Message: "Error Get All User",
			Result:  err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, utils.ResponseMessage{
		Status:  200,
		Message: "Get All User Succes",
		Result:  user,
	})

}
