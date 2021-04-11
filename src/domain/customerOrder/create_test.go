package customerOrder

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/mocks"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreate(t *testing.T) {
	testCases := map[string]struct {
		req coModels.CreateReq

		getReferralCodeCreatorRes coModels.CustomerOrder
		getReferralCodeCreatorErr error

		getReferralCodeRedeemerRes coModels.CustomerOrder
		getReferralCodeRedeemerErr error

		getLastCodeRes int
		getLastCodeErr error

		saveRes coModels.CustomerOrder
		saveErr error

		want    coModels.CustomerOrder
		wantErr error
	}{
		"success, with referralCode": {
			req: coModels.CreateReq{
				Email:        "al",
				ReferralCode: 2,
			},
			getReferralCodeCreatorRes: coModels.CustomerOrder{},
			getReferralCodeCreatorErr: nil,

			getReferralCodeRedeemerRes: coModels.CustomerOrder{},
			getReferralCodeRedeemerErr: mongo.ErrNoDocuments,

			getLastCodeRes: 4,
			getLastCodeErr: nil,

			saveRes: coModels.CustomerOrder{
				Email:            "al",
				OwnReferralCode:  5,
				UsedReferralCode: 2,
			},
			saveErr: nil,

			want: coModels.CustomerOrder{
				Email:            "al",
				OwnReferralCode:  5,
				UsedReferralCode: 2,
			},
			wantErr: nil,
		},
		"success, no referralCode": {
			req: coModels.CreateReq{
				Email: "al",
			},

			getLastCodeRes: 4,
			getLastCodeErr: nil,

			saveRes: coModels.CustomerOrder{
				Email:            "al",
				OwnReferralCode:  5,
				UsedReferralCode: 2,
			},
			saveErr: nil,

			want: coModels.CustomerOrder{
				Email:            "al",
				OwnReferralCode:  5,
				UsedReferralCode: 2,
			},
			wantErr: nil,
		},

		"error, on getReferralCodeCreator notFound, use user friendly text": {
			req: coModels.CreateReq{
				ReferralCode: 3,
			},

			getReferralCodeCreatorRes: coModels.CustomerOrder{},
			getReferralCodeCreatorErr: mongo.ErrNoDocuments,

			want:    coModels.CustomerOrder{},
			wantErr: coModels.ErrReferralCodeNotExists,
		},

		"error, on getReferralCodeCreator throws": {
			req: coModels.CreateReq{
				ReferralCode: 3,
			},

			getReferralCodeCreatorRes: coModels.CustomerOrder{},
			getReferralCodeCreatorErr: mocks.ErrDummy,

			want:    coModels.CustomerOrder{},
			wantErr: mocks.ErrDummy,
		},

		"error, on getReferralCodeRedeemer": {
			req: coModels.CreateReq{
				ReferralCode: 3,
			},

			getReferralCodeRedeemerRes: coModels.CustomerOrder{},
			getReferralCodeRedeemerErr: mocks.ErrDummy,

			want:    coModels.CustomerOrder{},
			wantErr: mocks.ErrDummy,
		},
		"error, on getReferralCodeRedeemer, alreadyUsed": {
			req: coModels.CreateReq{
				ReferralCode: 23,
			},

			getReferralCodeRedeemerRes: coModels.CustomerOrder{
				UsedReferralCode: 23,
			},
			getReferralCodeRedeemerErr: nil,

			want:    coModels.CustomerOrder{},
			wantErr: coModels.ErrReferralCodeAlreadyRedeemed,
		},
		"error, on getLastCode": {
			req: coModels.CreateReq{
				ReferralCode: 2,
			},

			getReferralCodeRedeemerRes: coModels.CustomerOrder{},
			getReferralCodeRedeemerErr: nil,

			getLastCodeRes: 0,
			getLastCodeErr: mocks.ErrDummy,

			want:    coModels.CustomerOrder{},
			wantErr: mocks.ErrDummy,
		},

		"error, on saave": {
			req: coModels.CreateReq{
				ReferralCode: 2,
			},

			getReferralCodeRedeemerRes: coModels.CustomerOrder{},
			getReferralCodeRedeemerErr: nil,

			getLastCodeRes: 0,
			getLastCodeErr: nil,

			saveRes: coModels.CustomerOrder{},
			saveErr: mocks.ErrDummy,

			want:    coModels.CustomerOrder{},
			wantErr: mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				result, err := m.sut.Create(m.Ctx, tc.req)

				// act
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "CustomerOrderDomain.Create")

			//
			if tc.req.ReferralCode > 0 {
				//
				m.Repo.CustomerOrder.EXPECT().GetReferralCodeCreator(m.Ctx, tc.req.ReferralCode).
					Return(tc.getReferralCodeCreatorRes, tc.getReferralCodeCreatorErr)

				if tc.getReferralCodeCreatorErr != nil {
					return
				}

				//
				m.Repo.CustomerOrder.EXPECT().GetReferralCodeRedeemer(m.Ctx, tc.req.ReferralCode).
					Return(tc.getReferralCodeRedeemerRes, tc.getReferralCodeRedeemerErr)

				if tc.getReferralCodeRedeemerErr != nil && !errors.Is(tc.getReferralCodeRedeemerErr, mongo.ErrNoDocuments) {
					return
				}

				if tc.req.ReferralCode == tc.getReferralCodeRedeemerRes.UsedReferralCode {
					return
				}
			}

			//
			m.Repo.CustomerOrder.EXPECT().GetLastCode(m.Ctx).Return(tc.getLastCodeRes, tc.getLastCodeErr)
			if tc.getLastCodeErr != nil {
				return
			}

			//
			code := tc.getLastCodeRes + 1
			// nextCode := fmt.Sprintf("%v", code)
			// m.Trace.Tracer.EXPECT().Log("generateReferralCode: " + nextCode).Return()

			//
			newco := coModels.CustomerOrder{
				Email:            tc.req.Email,
				OwnReferralCode:  code,
				UsedReferralCode: tc.req.ReferralCode,
				Orders:           tc.req.Orders,
			}
			m.Repo.CustomerOrder.EXPECT().Save(m.Ctx, newco).
				Return(tc.saveRes, tc.saveErr)
		})
	}
}
