package main

import (
	"context"

	"github.com/linggaaskaedo/go-playground-wire-v3/lib/config"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/graceful"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/logger"
)

func main() {
	logger.InitLogger()

	initProtocol := InitHttpProtocol()

	graceful.GracefulShutdown(
		context.TODO(),
		config.Get().Application.Graceful.MaxSecond,
		map[string]graceful.Operation{
			"http": func(ctx context.Context) error {
				return initProtocol.Shutdown(ctx)
			},
		},
	)

	initProtocol.Listen()
}
