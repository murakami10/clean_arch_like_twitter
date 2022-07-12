package infra

import (
	"clean_arch/domain"
	"clean_arch/interfaces/database"
	"clean_arch/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	Interactor usecase.UserInteractor
}

func NewUserHandler(e *echo.Echo, s database.SqlHandler) {
	handler := &UserHandler{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: s,
			},
		},
	}

	e.GET("/users", handler.Index)
	e.GET("/users/:id", handler.Show)
	e.POST("/users", handler.Create)

}

func (h *UserHandler) Index(c echo.Context) error {
	users, err := h.Interactor.Users()
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, users)
}

func (h *UserHandler) Create(c echo.Context) error {
	u := domain.User{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(500, err)
	}
	err := h.Interactor.Add(u)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, nil)
}

func (h *UserHandler) Show(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(500, err)
	}
	user, err := h.Interactor.UserById(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
}
