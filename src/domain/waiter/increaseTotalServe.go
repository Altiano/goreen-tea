package waiter

import (
	"context"

	waiterModels "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/models"
)

func (d domain) IncreaseTotalServe(ctx context.Context, waiter waiterModels.Waiter) error {
	ctx, span := d.Tracer.Start(ctx, "WaiterDomain.IncreaseTotalServe")
	defer span.End()

	return d.Repo.IncreaseTotalServe(ctx, waiter)
}
