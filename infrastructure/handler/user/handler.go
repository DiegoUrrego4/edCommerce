package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/DiegoUrrego4/edCommerce/domain/user"
	"github.com/DiegoUrrego4/edCommerce/model"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{}
}

func (h *handler) Create(c echo.Context) error {
	newUser := model.User{}

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.Create(&newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newUser)

}

func (h *handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)

}
