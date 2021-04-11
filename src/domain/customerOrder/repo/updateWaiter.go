package repo

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) UpdateWaiter(ctx context.Context, co coModels.CustomerOrder, waiterName string) error {
	ctx, span := r.Tracer.Start(ctx, "Repo.UpdateWaiter")
	defer span.End()

	_, err := r.Coll.UpdateOne(ctx, bson.M{
		"_id": co.ID,
	}, bson.M{
		"$set": bson.M{
			"WaiterName": waiterName,
		},
	})

	return err
}
