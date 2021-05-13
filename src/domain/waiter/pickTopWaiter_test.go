package waiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func TestPickTopWaiter(t *testing.T) {
	// T = TopRated; N = NormalRated
	testCases := map[string]struct {
		getByFreeTime []interface{}
		want          waiterModels.Waiter
		wantErr       error
	}{
		"success, NTT, pick the second": {
			getByFreeTime: []interface{}{
				[]waiterModels.Waiter{{
					Name:  "Monday",
					Rated: 5,
				}, {
					Name:  "Tuesday",
					Rated: 9,
				}, {
					Name:  "Wednesday",
					Rated: 8,
				}},
				nil,
			},
			want: waiterModels.Waiter{
				Name:  "Tuesday",
				Rated: 9,
			},
			wantErr: nil,
		},
		"success, NNT, pick the last": {
			getByFreeTime: []interface{}{
				[]waiterModels.Waiter{{
					Name:  "Monday",
					Rated: 4,
				}, {
					Name:  "Tuesday",
					Rated: 5,
				}, {
					Name:  "Wednesday",
					Rated: 9,
				}},
				nil,
			},
			want: waiterModels.Waiter{
				Name:  "Wednesday",
				Rated: 9,
			},
			wantErr: nil,
		},

		"success, NNN, nothing match so pick the first": {
			getByFreeTime: []interface{}{
				[]waiterModels.Waiter{{
					Name:  "Monday",
					Rated: 7,
				}, {
					Name:  "Tuesday",
					Rated: 7,
				}, {
					Name:  "Wednesday",
					Rated: 7,
				}},
				nil,
			},
			want: waiterModels.Waiter{
				Name:  "Monday",
				Rated: 7,
			},
			wantErr: nil,
		},
		"error, empty getByFreetime": {
			getByFreeTime: []interface{}{
				[]waiterModels.Waiter{},
				nil,
			},
			want:    waiterModels.Waiter{},
			wantErr: waiterModels.ErrEmptyList,
		},
		"error, getByFreetime throws": {
			getByFreeTime: []interface{}{
				[]waiterModels.Waiter{},
				mocks.ErrDummy,
			},
			want:    waiterModels.Waiter{},
			wantErr: mocks.ErrDummy,
		},
	}

	//
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				result, err := m.sut.PickTopWaiter(m.Ctx)

				// act
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "WaiterDomain.PickTopWaiter")

			//
			m.Repo.Waiter.EXPECT().GetByFreetime(m.Ctx).
				Return(tc.getByFreeTime[0], tc.getByFreeTime[1])
		})
	}
}
