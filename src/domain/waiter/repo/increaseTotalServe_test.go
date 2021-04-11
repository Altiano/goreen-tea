package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	waiterModels "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/models"
	"gitlab.com/altiano/golang-boilerplate/src/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestIncreaseTotalServe(t *testing.T) {
	testCases := map[string]struct {
		waiter       waiterModels.Waiter
		updateOneErr error
		wantErr      error
	}{
		"success": {
			waiter: waiterModels.Waiter{
				ID:   mocks.ObjectIDDummy,
				Name: "jeane",
			},

			updateOneErr: nil,
			wantErr:      nil,
		},

		"error": {
			waiter: waiterModels.Waiter{
				ID:   mocks.ObjectIDDummy,
				Name: "jeane",
			},

			updateOneErr: mocks.ErrDummy,
			wantErr:      mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupRepo(t)

			defer func() {
				// assert
				err := m.sut.IncreaseTotalServe(m.Ctx, tc.waiter)

				// act
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.IncreaseTotalServe")

			//
			m.Database.Coll.EXPECT().UpdateOne(m.Ctx, bson.M{
				"_id": tc.waiter.ID,
			}, bson.M{
				"$inc": bson.M{
					"TotalServe": 1,
				},
			}).Return(&mongo.UpdateResult{}, tc.updateOneErr)
		})
	}
}
