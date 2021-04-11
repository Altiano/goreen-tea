//+build wireinject

package di

import (
	"github.com/google/wire"
	"gitlab.com/altiano/golang-boilerplate/src/app"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/rest"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
	"gitlab.com/altiano/golang-boilerplate/src/shared"
)

func InjectIris(config shared.Config, app app.App) frameworks.Server {
	panic(
		wire.Build(
			trace.NewOtelManager,
			rest.NewIris,
		),
	)
}
