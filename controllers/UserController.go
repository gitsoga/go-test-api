package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct{}

func NewUserController() *userController {
    return new(userController)
}

func (hc *userController) Create(c echo.Context) error {
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "OK",
    ))
}

func (hc *userController) Index(c echo.Context) error {
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "OK",
    ))
}

func (hc *userController) Update(c echo.Context) error {
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "OK",
    ))
}

func (hc *userController) Delete(c echo.Context) error {
    return c.JSON(http.StatusOK, newResponse(
        http.StatusOK,
        http.StatusText(http.StatusOK),
        "OK",
    ))
}