package main

import (
	"service/config"
	"service/repository"
	"service/subscribes"

	"github.com/labstack/echo"
)

func main() {
	var e = echo.New()
	var dbConn = config.GetConnection()
	var repository = repository.NewRoomRepository(dbConn)
	subscribes.Listen(repository)

	e.Logger.Fatal(e.Start(":3003"))

}
