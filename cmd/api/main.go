package main

import (
	"time"

	"hackerNewsApi/internal/components"

	route "hackerNewsApi/internal/delivery/http/route"
)

func main() {

	app := components.AppConfig()

	env := app.Config

	timeout := time.Duration(env.ContextTimeout) * time.Second

	route.Setup(env, timeout, app.Server)

	app.Server.Run()
}
