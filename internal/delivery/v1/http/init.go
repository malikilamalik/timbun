package V1http

import (
	usersRouter "timbun/internal/delivery/v1/http/users"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	v1 := app.Group("/v1")
	usersRouter.InitUserRouter(v1)
}
