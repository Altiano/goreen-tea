package waiter

import (
	"context"

	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
	"gitlab.com/altiano/goreen-tea/src/domain/waiter/repo"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
)

type (
	Domain interface {
		PickTopWaiter(ctx context.Context) (waiterModels.Waiter, error)
		PickNormalWaiter(ctx context.Context) (waiterModels.Waiter, error)
		IncreaseTotalServe(ctx context.Context, waiter waiterModels.Waiter) error
	}

	domain struct {
		Tracer trace.Tracer
		Repo   repo.Repo
	}
)

func NewDomain(trace trace.Tracer, repo repo.Repo) Domain {
	return domain{
		Tracer: trace,
		Repo:   repo,
	}
}
