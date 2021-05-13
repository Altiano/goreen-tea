package assistanceCoordinator

import (
	"context"

	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
)

func (d domain) Assist(ctx context.Context, co coModels.CustomerOrder) error {
	ctx, span := d.Tracer.Start(ctx, "AssistanceCoordinator.Assist")
	defer span.End()

	var (
		waiter waiterModels.Waiter
		err    error
	)

	// Decide which waiter type to pick
	// based on the priority score
	if co.IsHighPriority() {
		waiter, err = d.Waiter.PickTopWaiter(ctx)
	} else {
		waiter, err = d.Waiter.PickNormalWaiter(ctx)
	}

	if err != nil {
		return err
	}

	// Notify CO of the selected waiter name
	if err := d.CustomerOrder.NotifyWaiterName(ctx, co, waiter.Name); err != nil {
		return err
	}

	// Increase total serve of the waiter
	// and Return
	return d.Waiter.IncreaseTotalServe(ctx, waiter)
}
