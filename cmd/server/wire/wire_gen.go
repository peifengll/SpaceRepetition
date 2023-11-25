// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/peifengll/SpaceRepetition/internal/handler"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"github.com/peifengll/SpaceRepetition/internal/server"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"github.com/peifengll/SpaceRepetition/pkg/app"
	"github.com/peifengll/SpaceRepetition/pkg/helper/sid"
	"github.com/peifengll/SpaceRepetition/pkg/jwt"
	"github.com/peifengll/SpaceRepetition/pkg/log"
	"github.com/peifengll/SpaceRepetition/pkg/server/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	db := repository.NewDB(viperViper, logger)
	client := repository.NewRedis(viperViper)
	repositoryRepository := repository.NewRepository(db, client, logger)
	transaction := repository.NewTransaction(repositoryRepository)
	sidSid := sid.NewSid()
	serviceService := service.NewService(transaction, logger, sidSid, jwtJWT)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	httpServer := server.NewHTTPServer(logger, viperViper, jwtJWT, userHandler)
	job := server.NewJob(logger)
	appApp := newApp(httpServer, job)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewTransaction, repository.NewUserRepository)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var serverSet = wire.NewSet(server.NewHTTPServer, server.NewJob, server.NewTask)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(app.WithServer(httpServer, job), app.WithName("demo-server"))
}
