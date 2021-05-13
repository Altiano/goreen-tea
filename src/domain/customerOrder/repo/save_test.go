package repo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestSave(t *testing.T) {
	testCases := map[string]struct {
		inputCo      coModels.CustomerOrder
		insertOneErr error
		want         coModels.CustomerOrder
		wantErr      error
	}{
		"success": {
			inputCo: coModels.CustomerOrder{
				Email: "al",
			},
			want: coModels.CustomerOrder{
				ID:    primitive.NewObjectID(),
				Email: "al",
			},
			wantErr: nil,
		},
		"error, insertOne throws": {
			inputCo: coModels.CustomerOrder{
				Email: "al",
			},
			insertOneErr: errors.New("an_error"),
			want:         coModels.CustomerOrder{},
			wantErr:      errors.New("an_error"),
		},
		"error, findByID throws": {
			inputCo: coModels.CustomerOrder{
				Email: "al",
			},
			want:    coModels.CustomerOrder{},
			wantErr: errors.New("an_error"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupRepo(t)

			defer func() {
				// act
				result, err := m.sut.Save(m.Ctx, tc.inputCo)

				// assert
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.Save")

			//
			m.Coll.EXPECT().InsertOne(m.Ctx, tc.inputCo).Return(&mongo.InsertOneResult{
				InsertedID: tc.want.ID,
			}, tc.insertOneErr)

			//
			if tc.insertOneErr == nil {
				m.Coll.EXPECT().FindByID(m.Ctx, tc.want.ID, &tc.inputCo).DoAndReturn(func(ctx, id interface{}, m *coModels.CustomerOrder) error {
					*m = tc.want
					return tc.wantErr
				})
			}

		})
	}
}
