package waiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func TestIncreaseTotalServe(t *testing.T) {
	testCases := map[string]struct {
		waiter                waiterModels.Waiter
		increaseTotalServeErr error
		wantErr               error
	}{
		"success": {
			waiter: waiterModels.Waiter{
				Name: "al",
			},
			increaseTotalServeErr: nil,
			wantErr:               nil,
		},

		"error": {
			waiter:                waiterModels.Waiter{},
			increaseTotalServeErr: mocks.ErrDummy,
			wantErr:               mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				err := m.sut.IncreaseTotalServe(m.Ctx, tc.waiter)

				// act
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "WaiterDomain.IncreaseTotalServe")

			//
			m.Repo.Waiter.EXPECT().IncreaseTotalServe(m.Ctx, tc.waiter).Return(tc.increaseTotalServeErr)
		})
	}
}
