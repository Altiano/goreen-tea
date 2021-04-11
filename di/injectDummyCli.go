//+build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.com/altiano/golang-boilerplate/src/app"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/cli"
	"gitlab.com/altiano/golang-boilerplate/src/shared"
)

func InjectDummyCLI(config shared.Config, app app.App) frameworks.Server {
	panic(
		wire.Build(
			cli.NewDummyCLI,
		),
	)
}
