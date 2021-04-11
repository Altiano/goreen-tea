package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	waiterModels "gitlab.com/altiano/golang-boilerplate/src/domain/waiter/models"
	"gitlab.com/altiano/golang-boilerplate/src/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetByFreetime(t *testing.T) {
	testCases := map[string]struct {
		want    []waiterModels.Waiter
		findErr error
		wantErr error
	}{
		"success": {
			want: []waiterModels.Waiter{
				{
					Name:       "d",
					Rated:      0,
					TotalServe: 0,
				},
				{
					Name:       "b",
					Rated:      7,
					TotalServe: 3,
				},
				{
					Name:       "a",
					Rated:      8,
					TotalServe: 5,
				},
				{
					Name:       "c",
					Rated:      6,
					TotalServe: 9,
				},
			},
			findErr: nil,
			wantErr: nil,
		},
		"error, on find": {
			want:    []waiterModels.Waiter{},
			findErr: mocks.ErrDummy,
			wantErr: mocks.ErrDummy,
		},
		"error, on cursor.All": {
			want:    []waiterModels.Waiter{},
			wantErr: mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupRepo(t)

			defer func() {
				// assert
				result, err := m.sut.GetByFreetime(m.Ctx)

				// act
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "Repo.GetByFreetime")

			//
			m.Database.Coll.EXPECT().Find(m.Ctx, bson.M{}, &options.FindOptions{
				Sort: bson.M{
					"TotalServe": 1,
				},
			}).Return(m.Database.Cursor, tc.findErr)

			if tc.findErr != nil {
				return
			}

			//
			m.Database.Cursor.EXPECT().All(m.Ctx, &[]waiterModels.Waiter{}).
				DoAndReturn(func(ctx context.Context, waiters *[]waiterModels.Waiter) error {

					if tc.wantErr != nil {
						return tc.wantErr
					}

					*waiters = tc.want
					return nil
				})
		})
	}
}
