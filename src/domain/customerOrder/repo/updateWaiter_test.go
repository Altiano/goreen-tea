package repo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUpdateWaiter(t *testing.T) {
	testCases := map[string]struct {
		co           coModels.CustomerOrder
		waiterName   string
		insertOneErr error
		want         coModels.CustomerOrder
		wantErr      error
	}{
		"success": {
			co: coModels.CustomerOrder{
				ID: primitive.NewObjectID(),
			},
			waiterName: "john",
			wantErr:    nil,
		},
		"error": {
			co: coModels.CustomerOrder{
				ID: primitive.NewObjectID(),
			},
			waiterName: "john",
			wantErr:    errors.New("an_error"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupRepo(t)
			defer func() {
				// act .
				err := m.sut.UpdateWaiter(m.Ctx, tc.co, tc.waiterName)

				// assert
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.UpdateWaiter")

			//
			m.Coll.EXPECT().UpdateOne(m.Ctx, bson.M{
				"_id": tc.co.ID,
			}, bson.M{
				"$set": bson.M{
					"WaiterName": tc.waiterName,
				},
			}).Return(&mongo.UpdateResult{}, tc.wantErr)

		})
	}
}
