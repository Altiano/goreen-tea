package repo

import (
	"context"

	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetReferralCodeCreator(ctx context.Context, referralCode int) (coModels.CustomerOrder, error) {
	ctx, span := r.Tracer.Start(ctx, "Repo.GetReferralCodeCreator")
	defer span.End()

	co := coModels.CustomerOrder{}

	err := r.Coll.FindOne(ctx, bson.M{
		"OwnReferralCode": referralCode,
	}).Decode(&co)

	if err != nil {
		return coModels.CustomerOrder{}, err
	}

	return co, nil
}
