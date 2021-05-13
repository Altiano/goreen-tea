package assistanceCoordinator

import (
	"context"

	"gitlab.com/altiano/goreen-tea/src/domain/customerOrder"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"gitlab.com/altiano/goreen-tea/src/domain/waiter"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
)

type (
	Domain interface {
		Assist(ctx context.Context, co coModels.CustomerOrder) error
	}

	domain struct {
		Tracer        trace.Tracer
		Waiter        waiter.Domain
		CustomerOrder customerOrder.Domain
	}
)

func NewDomain(
	trace trace.Tracer,
	waiter waiter.Domain,
	customerOrder customerOrder.Domain,
) Domain {
	return domain{
		Tracer:        trace,
		Waiter:        waiter,
		CustomerOrder: customerOrder,
	}
}
