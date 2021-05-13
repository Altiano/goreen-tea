package repo

import (
	"context"

	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetLastCode(ctx context.Context) (int, error) {
	ctx, span := r.Tracer.Start(ctx, "Repo.GetLastCode")
	defer span.End()

	co := coModels.CustomerOrder{}

	err := r.Coll.FindOne(ctx, bson.M{}, &options.FindOneOptions{
		Sort: bson.D{
			{
				Key:   "OwnReferralCode",
				Value: -1,
			},
		},
	}).Decode(&co)

	if err != nil && err != mongo.ErrNoDocuments {
		return 0, err
	}

	return co.OwnReferralCode, nil
}
