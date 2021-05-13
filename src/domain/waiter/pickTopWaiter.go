package waiter

import (
	"context"

	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
)

func (d domain) PickTopWaiter(ctx context.Context) (waiterModels.Waiter, error) {
	ctx, span := d.Tracer.Start(ctx, "WaiterDomain.PickTopWaiter")
	defer span.End()

	// Get waiter list ordered by lowest total serve
	list, err := d.Repo.GetByFreetime(ctx)

	if err != nil {
		return waiterModels.Waiter{}, err
	}

	if len(list) == 0 {
		return waiterModels.Waiter{}, waiterModels.ErrEmptyList
	}

	// Pick one from the list
	// default to the first waiter if no one is match
	picked := list[0]
	for _, waiter := range list {
		if waiter.Rated >= waiterModels.TopRated {
			picked = waiter
			break
		}
	}

	// Return
	return picked, nil
}
