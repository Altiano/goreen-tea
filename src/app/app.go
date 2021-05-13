package app

import (
	"context"

	"gitlab.com/altiano/goreen-tea/src/domain/assistanceCoordinator"
	"gitlab.com/altiano/goreen-tea/src/domain/customerOrder"
	"gitlab.com/altiano/goreen-tea/src/domain/visitorCounter"
	"gitlab.com/altiano/goreen-tea/src/frameworks/trace"
)

type (
	App interface {
		OnsiteOrder(ctx context.Context, request OnsiteOrderReq) (OnsiteOrderRes, error)
	}

	app struct {
		Tracer                trace.Tracer
		CustomerOrder         customerOrder.Domain
		AssistanceCoordinator assistanceCoordinator.Domain
		VisitorCounter        visitorCounter.Domain
		// Save to redis for easy access
	}
)

func NewApp(
	tracer trace.Tracer,
	customerOrder customerOrder.Domain,
	assistanceCoordinator assistanceCoordinator.Domain,
	visitorCounter visitorCounter.Domain,
) App {
	return app{
		tracer,
		customerOrder,
		assistanceCoordinator,
		visitorCounter,
	}
}
