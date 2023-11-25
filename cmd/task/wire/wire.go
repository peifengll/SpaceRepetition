//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/peifengll/SpaceRepetition/internal/server"
	"github.com/peifengll/SpaceRepetition/pkg/app"
	"github.com/peifengll/SpaceRepetition/pkg/log"
	"github.com/spf13/viper"
)

var taskSet = wire.NewSet(server.NewTask)

// build App
func newApp(task *server.Task) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		taskSet,
		newApp,
	))
}
