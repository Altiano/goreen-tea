package customerOrder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"gitlab.com/altiano/goreen-tea/src/frameworks/email"
	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func TestNotifyRedeemedCode(t *testing.T) {
	testCases := map[string]struct {
		co                        coModels.CustomerOrder
		getReferralCodeCreatorErr error
		wantErr                   error
	}{
		"success": {
			co: coModels.CustomerOrder{
				UsedReferralCode: 23,
			},
			wantErr: nil,
		},
		"error, getReferralCodeCreator": {
			co: coModels.CustomerOrder{
				UsedReferralCode: 23,
			},
			getReferralCodeCreatorErr: mocks.ErrDummy,
			wantErr:                   mocks.ErrDummy,
		},

		"error, send email": {
			co: coModels.CustomerOrder{
				UsedReferralCode: 23,
			},
			wantErr: mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				err := m.sut.NotifyRedeemedCode(m.Ctx, tc.co)

				// act
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "CustomerOrderDomain.NotifyRedeemedCode")

			//
			refererCo := coModels.CustomerOrder{
				Email:           "something",
				OwnReferralCode: tc.co.UsedReferralCode,
			}
			m.Repo.CustomerOrder.EXPECT().GetReferralCodeCreator(
				m.Ctx,
				tc.co.UsedReferralCode,
			).Return(refererCo, tc.getReferralCodeCreatorErr)

			if tc.getReferralCodeCreatorErr != nil {
				return
			}

			//
			// m.Trace.Tracer.EXPECT().Log("sendingTheEmail").Return()

			//
			// m.Trace.Tracer.EXPECT().Log("buildingEmailObj").Return()

			//
			m.Email.Manager.EXPECT().Send(email.Email{
				From: "mycompany.com",
				To:   refererCo.Email,
				Body: "Congratz " + refererCo.Email + ", your code used by " + tc.co.Email,
			}).Return(tc.wantErr)
		})
	}
}
