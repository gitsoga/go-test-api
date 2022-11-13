package server

import (
	"github.com/gitsoga/go-test-api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
    router := echo.New()
    router.Use(middleware.Logger())

	userController := controllers.NewUserController()
	router.POST("/users", userController.Create)
	router.GET("/users", userController.Index)
	router.PUT("/users", userController.Update)
	router.DELETE("/users", userController.Delete)

	return router, nil
}
