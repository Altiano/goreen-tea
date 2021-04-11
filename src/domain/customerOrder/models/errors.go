package coModels

import (
	"errors"
	"fmt"

	"gitlab.com/altiano/golang-boilerplate/src/shared"
)

var (
	ErrReferralCodeNotExists       = fmt.Errorf("%v: %w", shared.ErrBadRequest, errors.New("referral code not exists"))
	ErrReferralCodeAlreadyRedeemed = fmt.Errorf("%v: %w", shared.ErrBadRequest, errors.New("referral code already redeemed"))
)
