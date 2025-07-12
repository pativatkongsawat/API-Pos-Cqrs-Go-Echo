package main

import (
	"echo_cqrs/internal/config"
	"echo_cqrs/internal/database"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	config.Init()
	database.ConnectMySQL()

	e.Logger.Fatal(e.Start(":1323"))

}
