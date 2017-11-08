package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zaru/go-echo-api-test-sample/models"
)

type resultLists struct {
	Users []user.User `json:"users"`
}

type handler struct {
	UserModel user.UserModel
}

func (h *handler) GetIndex(c echo.Context) error {
	lists := h.UserModel.All()
	u := &resultLists{
		Users: lists,
	}
	return c.JSON(http.StatusOK, u)
}

func (h *handler) GetDetail(c echo.Context) error {
	id := c.Param("id")
	u := h.UserModel.GetDetail(id)
	return c.JSON(http.StatusOK, u)
}
