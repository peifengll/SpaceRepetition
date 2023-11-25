//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/peifengll/SpaceRepetition/internal/handler"
	"github.com/peifengll/SpaceRepetition/internal/repository"
	"github.com/peifengll/SpaceRepetition/internal/server"
	"github.com/peifengll/SpaceRepetition/internal/service"
	"github.com/peifengll/SpaceRepetition/pkg/app"
	"github.com/peifengll/SpaceRepetition/pkg/helper/sid"
	"github.com/peifengll/SpaceRepetition/pkg/jwt"
	"github.com/peifengll/SpaceRepetition/pkg/log"
	"github.com/peifengll/SpaceRepetition/pkg/server/http"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
