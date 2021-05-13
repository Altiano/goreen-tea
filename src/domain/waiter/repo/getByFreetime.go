package repo

import (
	"context"

	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) GetByFreetime(ctx context.Context) ([]waiterModels.Waiter, error) {
	ctx, span := r.Tracer.Start(ctx, "Repo.GetByFreetime")
	defer span.End()

	result := []waiterModels.Waiter{}
	cursor, err := r.Coll.Find(ctx, bson.M{}, &options.FindOptions{
		Sort: bson.M{
			"TotalServe": 1,
		},
	})

	if err != nil {
		return result, err
	}

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
