//+build wireinject

package di

import (
	"github.com/google/wire"
	application "gitlab.com/altiano/golang-boilerplate/src/app"
	"gitlab.com/altiano/golang-boilerplate/src/domain/assistanceCoordinator"
	"gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder"
	coRepo "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/repo"
	"gitlab.com/altiano/golang-boilerplate/src/domain/visitorCounter"
	"gitlab.com/altiano/golang-boilerplate/src/domain/waiter"
	waiterRepo "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/repo"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/database"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/email"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/memcache"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
	"gitlab.com/altiano/golang-boilerplate/src/shared"
)

func InjectApp(config shared.Config) application.App {
	panic(
		wire.Build(
			database.NewMongoDB,
			trace.NewOtelManager,
			email.NewManager,
			memcache.NewManager,

			coRepo.NewRepo,
			customerOrder.NewDomain,

			assistanceCoordinator.NewDomain,

			waiterRepo.NewRepo,
			waiter.NewDomain,

			visitorCounter.NewDomain,
			application.NewApp,
		),
	)
}
