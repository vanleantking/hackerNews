package main

import (
	"hackerNewsApi/internal/components"

	route "hackerNewsApi/internal/delivery/http/route"
)

func main() {

	app := components.AppConfig()

	routeCfg := route.NewRouteConfig(app.Server, app.Logger, app.DB)
	routeCfg.Setup()

	app.Server.Run()
}
