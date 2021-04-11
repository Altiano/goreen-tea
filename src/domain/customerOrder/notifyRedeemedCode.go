package customerOrder

import (
	"context"

	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/frameworks/email"
)

func (d domain) NotifyRedeemedCode(ctx context.Context, newCo coModels.CustomerOrder) error {
	ctx, span := d.Tracer.Start(ctx, "CustomerOrderDomain.NotifyRedeemedCode")
	defer span.End()

	// Get the creator of the referral code
	refererCo, err := d.Repo.GetReferralCodeCreator(ctx, newCo.UsedReferralCode)

	if err != nil {
		return err
	}

	// Build email to congratulate of this redeemed code
	emailObj := d.buildCongratzEmail(refererCo, newCo)

	// d.Trace.Log("sendingTheEmail")

	// Send the email
	// and Return
	return d.Email.Send(emailObj)
}

func (d domain) buildCongratzEmail(refererCo, newCo coModels.CustomerOrder) email.Email {
	// d.Trace.Log("buildingEmailObj")

	return email.Email{
		From: "mycompany.com",
		To:   refererCo.Email,
		Body: "Congratz " + refererCo.Email + ", your code used by " + newCo.Email,
	}
}
