package repo

import (
	"context"

	waiterModels "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r repo) IncreaseTotalServe(ctx context.Context, waiter waiterModels.Waiter) error {
	ctx, span := r.Tracer.Start(ctx, "Repo.IncreaseTotalServe")
	defer span.End()

	_, err := r.Coll.UpdateOne(ctx, bson.M{
		"_id": waiter.ID,
	}, bson.M{
		"$inc": bson.M{
			"TotalServe": 1,
		},
	})

	return err
}
