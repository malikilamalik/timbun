package usersRouter

import "github.com/gofiber/fiber/v2"

func InitUserRouter(router fiber.Router) {
	usersRouter := router.Group("/users")
	usersRouter.Get("", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})
}
