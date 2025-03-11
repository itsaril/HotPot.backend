package main

import (
	"hotpot/internal/core/cfg"
	"hotpot/internal/core/utils/logger"
	"hotpot/internal/core/utils/servers"
	"hotpot/internal/core/utils/servers/http"
	"hotpot/internal/pkg"

	"github.com/gofiber/fiber/v2"
)

func init() {
	// init logic here
}

func main() {
	appCfg := cfg.Inst()

	appLogger := logger.New(logger.DefaultConfig())

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	appRouter := pkg.NewRouter(appLogger)

	servMan := servers.NewServerManager()

	httpServ := http.NewFiber(appCfg.HttpPort, app, appRouter.Init)

	servMan.AddServer(httpServ)

	servMan.StartAll()

	defer servMan.StopAll()
}
