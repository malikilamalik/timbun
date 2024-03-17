package V1http

import (
	usersRouter "timbun/internal/delivery/v1/http/users"

	"github.com/labstack/echo/v4"
)

func InitRouter(app *echo.Group) {
	v1 := app.Group("/v1")
	usersRouter.InitUserRouter(v1)
}
