package server

import (
	"timbun/config"
	V1http "timbun/internal/delivery/v1/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	//Init Config
	config.InitConfig()

	//Init Fiber
	app := fiber.New()

	//Init Router
	V1http.InitRouter(app)

	app.Listen(":3000")
}
