package customerOrder

import (
	"context"
	"errors"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d domain) Create(ctx context.Context, req coModels.CreateReq) (coModels.CustomerOrder, error) {
	ctx, s := d.Tracer.Start(ctx, "CustomerOrderDomain.Create")
	defer s.End()

	//
	if req.ReferralCode > 0 {
		// Validate whether the code exists
		_, err := d.Repo.GetReferralCodeCreator(ctx, req.ReferralCode)

		if errors.Is(err, mongo.ErrNoDocuments) {
			return coModels.CustomerOrder{}, coModels.ErrReferralCodeNotExists
		}

		if err != nil {
			return coModels.CustomerOrder{}, err
		}

		// Validate whether the code already used
		redeemer, err := d.Repo.GetReferralCodeRedeemer(ctx, req.ReferralCode)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			return coModels.CustomerOrder{}, err
		}

		if redeemer.UsedReferralCode == req.ReferralCode {
			return coModels.CustomerOrder{}, coModels.ErrReferralCodeAlreadyRedeemed
		}
	}

	// Generate new CO referral code
	lastCode, err := d.Repo.GetLastCode(ctx)
	if err != nil {
		return coModels.CustomerOrder{}, err
	}

	newCode := lastCode + 1
	// d.Trace.Log("generateReferralCode: " + strconv.Itoa(newCode))

	// Save and Return response
	newCo := coModels.CustomerOrder{
		Email:            req.Email,
		OwnReferralCode:  newCode,
		UsedReferralCode: req.ReferralCode,
		Orders:           req.Orders,
	}
	return d.Repo.Save(ctx, newCo)
}
