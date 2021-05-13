//+build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.com/altiano/goreen-tea/src/app"
	"gitlab.com/altiano/goreen-tea/src/frameworks"
	"gitlab.com/altiano/goreen-tea/src/frameworks/cli"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

func InjectDummyCLI(config shared.Config, app app.App) frameworks.Server {
	panic(
		wire.Build(
			cli.NewDummyCLI,
		),
	)
}
