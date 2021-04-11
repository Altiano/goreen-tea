package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/mocks"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetReferralCodeCreator(t *testing.T) {
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
				OwnReferralCode: 2,
			},
			findOneErr: nil,
			want: coModels.CustomerOrder{
				OwnReferralCode: 2,
			},
			wantErr: nil,
		},
		// "success, accept errNoDocument": {
		// 	referralCode: 2,
		// 	findOneRes:  coModels.CustomerOrder{},
		// 	findOneErr:  mongo.ErrNoDocuments,
		// 	want:        coModels.CustomerOrder{},
		// 	wantErr:     nil,
		// },
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
				result, err := m.sut.GetReferralCodeCreator(m.Ctx, tc.referralCode)

				// assert
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.GetReferralCodeCreator")

			//
			m.Coll.EXPECT().FindOne(m.Ctx, bson.M{
				"OwnReferralCode": tc.referralCode,
			}).Return(m.SingleResult)

			//
			m.SingleResult.EXPECT().Decode(&coModels.CustomerOrder{}).DoAndReturn(func(m *coModels.CustomerOrder) error {
				if tc.findOneErr != nil {
					return tc.findOneErr
				}

				*m = tc.want
				return nil
			})

		})
	}
}
