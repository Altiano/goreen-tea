package customerOrder

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
)

func (d domain) NotifyWaiterName(ctx context.Context, newCo coModels.CustomerOrder, waiterName string) error {
	ctx, span := d.Tracer.Start(ctx, "CustomerOrderDomain.NotifyWaiterName")
	defer span.End()

	return d.Repo.UpdateWaiter(ctx, newCo, waiterName)
}
