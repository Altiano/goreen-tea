package app

import (
	"context"

	"gitlab.com/altiano/golang-boilerplate/src/domain/assistanceCoordinator"
	"gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder"
	"gitlab.com/altiano/golang-boilerplate/src/domain/visitorCounter"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/trace"
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
