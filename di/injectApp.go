//+build wireinject

package di

import (
	"github.com/google/wire"
	application "gitlab.com/altiano/goreen-tea/src/app"
	"gitlab.com/altiano/goreen-tea/src/domain/assistanceCoordinator"
	"gitlab.com/altiano/goreen-tea/src/domain/customerOrder"
	coRepo "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/repo"
	"gitlab.com/altiano/goreen-tea/src/domain/visitorCounter"
	"gitlab.com/altiano/goreen-tea/src/domain/waiter"
	waiterRepo "gitlab.com/altiano/goreen-tea/src/domain/waiter/repo"
	"gitlab.com/altiano/goreen-tea/src/frameworks/database"
	"gitlab.com/altiano/goreen-tea/src/frameworks/email"
	"gitlab.com/altiano/goreen-tea/src/frameworks/memcache"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
	"gitlab.com/altiano/goreen-tea/src/shared"
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
