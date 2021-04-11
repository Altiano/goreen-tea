package repo

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetReferralCodeRedeemer(ctx context.Context, referralCode int) (coModels.CustomerOrder, error) {
	ctx, span := r.Tracer.Start(ctx, "Repo.GetReferralCodeRedeemer")
	defer span.End()

	co := coModels.CustomerOrder{}

	// NOTE, don't supress errNoDocument, because it is easier that way
	if err := r.Coll.FindOne(ctx, bson.M{
		"UsedReferralCode": referralCode,
	}).Decode(&co); err != nil {
		return coModels.CustomerOrder{}, err
	}

	return co, nil
}
