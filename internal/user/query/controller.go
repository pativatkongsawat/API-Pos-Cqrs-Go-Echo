package user

import (
	"echo_cqrs/internal/database"

	"github.com/labstack/echo"
)

func GetAllUser(ctx echo.Context) error {

	userService := DatabaseRequest{DB: database.DBMYSQL}

	return nil

}
