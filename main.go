package main

import (
	"github.com/labstack/echo"
	"github.com/zaru/go-echo-api-test-sample/handlers"
)

func main() {
	e := echo.New()

	e.GET("/users", users.GetIndex)
	e.GET("/users/:id", users.GetDetail)

	e.Logger.Fatal(e.Start(":1324"))
}
