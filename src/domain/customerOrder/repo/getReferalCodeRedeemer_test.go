package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"gitlab.com/altiano/goreen-tea/src/mocks"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetReferralCodeRedeemer(t *testing.T) {
	testCases := map[string]struct {
		referralCode int
		findOneRes   coModels.CustomerOrder
		findOneErr   error
		want         coModels.CustomerOrder
		wantErr      error
	}{
		"success": {
			referralCode: 2,
			findOneRes: coModels.CustomerOrder{
				UsedReferralCode: 2,
			},
			findOneErr: nil,
			want: coModels.CustomerOrder{
				UsedReferralCode: 2,
			},
			wantErr: nil,
		},

		"error": {
			referralCode: 2,
			findOneRes:   coModels.CustomerOrder{},
			findOneErr:   mocks.ErrDummy,
			want:         coModels.CustomerOrder{},
			wantErr:      mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupRepo(t)

			defer func() {
				// act .
				result, err := m.sut.GetReferralCodeRedeemer(m.Ctx, tc.referralCode)

				// assert
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.GetReferralCodeRedeemer")

			//
			m.Coll.EXPECT().FindOne(m.Ctx, bson.M{
				"UsedReferralCode": tc.referralCode,
			}).Return(m.SingleResult)

			//
			m.SingleResult.EXPECT().Decode(&coModels.CustomerOrder{}).DoAndReturn(func(m *coModels.CustomerOrder) error {
				if tc.findOneErr != nil {
					return tc.findOneErr
				}

				*m = tc.findOneRes
				return nil
			})

		})
	}
}
