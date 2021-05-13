package repo

import (
	"context"

	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
)

func (r *repo) Save(ctx context.Context, co coModels.CustomerOrder) (coModels.CustomerOrder, error) {
	ctx, span := r.Tracer.Start(ctx, "Repo.Save")
	defer span.End()

	ir, err := r.Coll.InsertOne(ctx, co)

	if err != nil {
		return coModels.CustomerOrder{}, err
	}

	if err = r.Coll.FindByID(ctx, ir.InsertedID, &co); err != nil {
		return coModels.CustomerOrder{}, err
	}

	return co, nil
}
