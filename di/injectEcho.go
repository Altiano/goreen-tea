//+build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.com/altiano/goreen-tea/src/app"
	"gitlab.com/altiano/goreen-tea/src/frameworks"
	"gitlab.com/altiano/goreen-tea/src/frameworks/rest"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

func InjectEcho(config shared.Config, app app.App) frameworks.Server {
	panic(
		wire.Build(
			trace.NewOtelManager,
			rest.NewEcho,
		),
	)
}
