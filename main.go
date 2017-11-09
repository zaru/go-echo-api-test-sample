package main

import (
	"github.com/labstack/echo"
	"github.com/zaru/go-echo-api-test-sample/handlers"
	"github.com/zaru/go-echo-api-test-sample/models"
)

func main() {
	e := echo.New()

	h := users.NewHandler(user.NewUserModel())

	e.GET("/users", h.GetIndex)
	e.GET("/users/:id", h.GetDetail)

	e.Logger.Fatal(e.Start(":1324"))
}
