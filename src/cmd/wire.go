//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	jwtauth "github.com/linggaaskaedo/go-playground-wire-v3/lib/auth"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/http"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/http/router"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/handler/rest"
	newsrepo "github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/repository"
	newssvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/service"
)

// wiring jwt auth
var jwtAuth = wire.NewSet(
	jwtauth.NewJwt,
	wire.Bind(
		new(jwtauth.JwtToken),
		new(*jwtauth.JwtTokenImpl),
	),
)

// Wiring for domain
// news
var newsRepo = wire.NewSet(
	newsrepo.NewNewsRepository,
	wire.Bind(
		new(newsrepo.NewsRepository),
		new(*newsrepo.NewsRepositoryImpl),
	),
)

var newsSvc = wire.NewSet(
	newssvc.NewNewsService,
	wire.Bind(
		new(newssvc.NewsService),
		new(*newssvc.NewsServiceImpl),
	),
)

// Wiring for http protocol
var restHandler = wire.NewSet(
	rest.NewRestHandler,
)

// Wiring protocol routing
var httpRouter = wire.NewSet(
	router.NewHttpRouter,
)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(
		database.NewMysqlClient,
		database.NewPostgresClient,
		database.NewScribleClient,
		// transaction.NewTransaction,
		newsRepo,
		jwtAuth,
		newsSvc,
		restHandler,
		httpRouter,
		http.NewHttpProtocol,
	)

	return nil
}
