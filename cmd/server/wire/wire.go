//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/peifengll/SpaceRepetition/internal/handler"
	"github.com/peifengll/SpaceRepetition/internal/query"
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
	repository.NewFloderRepository,
	repository.NewDeckRepository,
	repository.NewKnowledgeRepository,
	repository.NewRecordRepository,
	repository.NewSectionRepository,
	repository.NewAnnouncementsRepository,
	repository.NewAnnouncementReadRecordRepository,
	repository.NewAdminRepository,
)

var querySet = wire.NewSet(
	query.NewQuery,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewFloderService,
	service.NewDeckService,
	service.NewKnowledgeService,
	service.NewRecordService,
	service.NewSectionService,
	service.NewAnnouncementsService,
	service.NewAdminService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewFloderHandler,
	handler.NewDeckHandler,
	handler.NewKnowledgeHandler,
	handler.NewRecordHandler,
	handler.NewSectionHandler,
	handler.NewAdminHandler,
	handler.NewAnnouncementsHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServerFont,
	server.NewHTTPServerAdmin,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.ServerFont, httpServer2 *http.ServerAdmin, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, httpServer2, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		querySet,
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
