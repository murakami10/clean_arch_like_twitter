package main

import (
	"github.com/labstack/echo"
)

func main() {
	//sqlHandler := infra.NewSqlHandler()

	//defer func() {
	//	err := sqlHandler.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	e := echo.New()
	//infra.NewUserHandler(e, sqlHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
