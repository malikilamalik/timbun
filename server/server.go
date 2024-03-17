package server

import (
	"timbun/config"
	V1http "timbun/internal/delivery/v1/http"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	//Init Config
	config.InitConfig()

	//Init Echo
	e := echo.New()
	g := e.Group("")
	//Init Router
	V1http.InitRouter(g)

	e.Logger.Fatal(e.Start(":1323"))
}
