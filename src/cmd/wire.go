//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	jwtauth "github.com/linggaaskaedo/go-playground-wire-v3/lib/auth"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/database"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/http"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/http/router"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/rabbit"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/rabbit/consumer"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/scheduler"
	"github.com/linggaaskaedo/go-playground-wire-v3/lib/scheduler/task"
	pubsubrabbithandler "github.com/linggaaskaedo/go-playground-wire-v3/src/handler/pubsub/rabbit"
	resthandler "github.com/linggaaskaedo/go-playground-wire-v3/src/handler/rest"
	schedulerhandler "github.com/linggaaskaedo/go-playground-wire-v3/src/handler/scheduler"
	newsrepo "github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/repository"
	newssvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/service"
	productrepo "github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/repository"
	productsvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/service"
	userrepo "github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/repository"
	usersvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/user/service"
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

// product
var productRepo = wire.NewSet(
	productrepo.NewProductRepository,
	wire.Bind(
		new(productrepo.ProductRepository),
		new(*productrepo.ProductRepositoryImpl),
	),
)

var productSvc = wire.NewSet(
	productsvc.NewProductService,
	wire.Bind(
		new(productsvc.ProductService),
		new(*productsvc.ProductServiceImpl),
	),
)

// user
var userRepo = wire.NewSet(
	userrepo.NewUserRepository,
	wire.Bind(
		new(userrepo.UserRepository),
		new(*userrepo.UserRepositoryImpl),
	),
)

var userSvc = wire.NewSet(
	usersvc.NewUserService,
	wire.Bind(
		new(usersvc.UserService),
		new(*usersvc.UserServiceImpl),
	),
)

// Wiring for http protocol
var restHandler = wire.NewSet(
	resthandler.NewRestHandler,
)

var schedulerHandler = wire.NewSet(
	schedulerhandler.NewSchedulerHandler,
)

var pubsubRabbitHandler = wire.NewSet(
	pubsubrabbithandler.NewRabbitHandler,
)

var schedulerTask = wire.NewSet(
	task.NewSchedulerTask,
)

var pubsubConsumer = wire.NewSet(
	consumer.NewPubsubConsumer,
)

// Wiring protocol routing
var httpRouter = wire.NewSet(
	router.NewHttpRouter,
)

func InitMySQL() *database.MysqlImpl {
	wire.Build(
		database.NewMysqlClient,
	)

	return nil
}

func InitPostgres() *database.PostgresImpl {
	wire.Build(
		database.NewPostgresClient,
	)

	return nil
}

func InitScribble() *database.ScribleImpl {
	wire.Build(
		database.NewScribleClient,
	)

	return nil
}

func InitScheduler(a *database.MysqlImpl, b *database.PostgresImpl, c *database.ScribleImpl) *scheduler.SchedulerImpl {
	wire.Build(
		newsRepo,
		jwtAuth,
		newsSvc,
		schedulerHandler,
		schedulerTask,
		scheduler.NewScheduler,
	)

	return nil
}

func InitPubsub(a *database.MysqlImpl, b *database.PostgresImpl, c *database.ScribleImpl) *rabbit.RabbitImpl {
	wire.Build(
		userRepo,
		jwtAuth,
		userSvc,
		pubsubRabbitHandler,
		pubsubConsumer,
		rabbit.NewRabbit,
	)

	return nil
}

func InitServer(a *database.MysqlImpl, b *database.PostgresImpl, c *database.ScribleImpl) *http.HttpImpl {
	wire.Build(
		newsRepo,
		productRepo,
		jwtAuth,
		newsSvc,
		productSvc,
		restHandler,
		httpRouter,
		http.NewHttpProtocol,
	)

	return nil
}
