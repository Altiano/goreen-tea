package app

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	OnsiteOrderReq struct {
		Email        string   `validate:"required"`
		Orders       []string `validate:"required,dive,required"`
		ReferralCode int      `validate:"gte=0"`
	}

	OnsiteOrderRes struct {
		ID               primitive.ObjectID
		YourReferralCode int
	}
)

func (a app) OnsiteOrder(ctx context.Context, req OnsiteOrderReq) (OnsiteOrderRes, error) {
	ctx, s := a.Tracer.Start(ctx, "App.OnsiteOrder")
	defer s.End()

	// Validate request
	if err := shared.Validate(req); err != nil {
		return OnsiteOrderRes{}, err
	}

	// Create customer order
	newCo, err := a.CustomerOrder.Create(ctx, coModels.CreateReq{
		Email:        req.Email,
		Orders:       req.Orders,
		ReferralCode: req.ReferralCode,
	})

	if err != nil {
		return OnsiteOrderRes{}, err
	}

	// Assign a customer assistant
	// and handle all the necessary side effects
	err = a.AssistanceCoordinator.Assist(ctx, newCo)

	if err != nil {
		return OnsiteOrderRes{}, err
	}

	go func() {
		// Increase current vistor count
		a.VisitorCounter.IncreaseVisits(ctx)

		// Send redeemed code congratz
		a.CustomerOrder.NotifyRedeemedCode(ctx, newCo)
	}()

	// Return response
	return OnsiteOrderRes{
		ID:               newCo.ID,
		YourReferralCode: newCo.OwnReferralCode,
	}, nil
}
