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
	initPubsub := InitPubsub(initMySQL, initPostgres, initScribble)
	initServer := InitServer(initMySQL, initPostgres, initScribble)

	graceful.GracefulShutdown(
		context.TODO(),
		config.Get().Application.Graceful.MaxSecond,
		map[string]graceful.Operation{
			"scheduler": func(ctx context.Context) error {
				return initScheduler.Shutdown(ctx)
			},
			"rabbit": func(ctx context.Context) error {
				return initPubsub.Shutdown(ctx)
			},
			"http": func(ctx context.Context) error {
				return initServer.Shutdown(ctx)
			},
		},
	)

	initScheduler.Start()
	initPubsub.Start()
	initServer.Listen()
}
