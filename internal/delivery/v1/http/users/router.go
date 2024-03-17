package usersRouter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserRouter(router *echo.Group) {
	usersRouter := router.Group("/users")
	usersRouter.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/1/files/*")
	})
}
