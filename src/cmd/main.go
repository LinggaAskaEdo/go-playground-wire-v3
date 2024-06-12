package main

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/graceful"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/logger"
)

func main() {
	logger.InitLogger()

	initMySQL := InitMySQL()
	initPostgres := InitPostgres()
	initScribble := InitScribble()
	initScheduler := InitScheduler(initMySQL, initPostgres, initScribble)
	initServer := InitServer(initMySQL, initPostgres, initScribble)

	graceful.GracefulShutdown(
		context.TODO(),
		config.Get().Application.Graceful.MaxSecond,
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return initServer.Shutdown(ctx)
			},
			"scheduler": func(ctx context.Context) error {
				return initScheduler.Shutdown(ctx)
			},
		},
	)

	initScheduler.Start()
	initServer.Listen()
}
