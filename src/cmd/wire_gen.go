// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/auth"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/http"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/http/router"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/handler/rest"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/repository"
	"github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/service"
)

// Injectors from wire.go:

func InitHttpProtocol() *http.HttpImpl {
	scribleImpl := database.NewScribleClient()
	jwtTokenImpl := auth.NewJwt(scribleImpl)
	mysqlImpl := database.NewMysqlClient()
	postgresImpl := database.NewPostgresClient()
	newsRepositoryImpl := repository.NewNewsRepository(mysqlImpl, postgresImpl)
	newsServiceImpl := service.NewNewsService(jwtTokenImpl, newsRepositoryImpl)
	restHandlerImpl := rest.NewRestHandler(newsServiceImpl)
	httpRouterImpl := router.NewHttpRouter(restHandlerImpl)
	httpImpl := http.NewHttpProtocol(httpRouterImpl)
	return httpImpl
}

// wire.go:

// wiring jwt auth
var jwtAuth = wire.NewSet(auth.NewJwt, wire.Bind(
	new(auth.JwtToken),
	new(*auth.JwtTokenImpl),
),
)

// Wiring for domain
// news
var newsRepo = wire.NewSet(repository.NewNewsRepository, wire.Bind(
	new(repository.NewsRepository),
	new(*repository.NewsRepositoryImpl),
),
)

var newsSvc = wire.NewSet(service.NewNewsService, wire.Bind(
	new(service.NewsService),
	new(*service.NewsServiceImpl),
),
)

// Wiring for http protocol
var restHandler = wire.NewSet(rest.NewRestHandler)

// Wiring protocol routing
var httpRouter = wire.NewSet(router.NewHttpRouter)