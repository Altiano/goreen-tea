package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"gitlab.com/altiano/goreen-tea/src/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetLastCode(t *testing.T) {
	testCases := map[string]struct {
		want       int
		findOneRes coModels.CustomerOrder
		findOneErr error
		wantErr    error
	}{
		"success": {
			findOneRes: coModels.CustomerOrder{
				OwnReferralCode: 2,
			},
			findOneErr: nil,
			want:       2,
			wantErr:    nil,
		},
		"success, accept noDocuments found": {
			findOneRes: coModels.CustomerOrder{},
			findOneErr: mongo.ErrNoDocuments,
			want:       0,
			// wantErr:    nil,
		},
		"error, Decode throws": {
			findOneRes: coModels.CustomerOrder{},
			findOneErr: mocks.ErrDummy,
			want:       0,
			wantErr:    mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupRepo(t)

			defer func() {
				// act
				result, err := m.sut.GetLastCode(m.Ctx)

				// assert
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.GetLastCode")

			//
			m.Coll.EXPECT().FindOne(m.Ctx, bson.M{}, &options.FindOneOptions{
				Sort: bson.D{
					{
						Key:   "OwnReferralCode",
						Value: -1,
					},
				},
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
