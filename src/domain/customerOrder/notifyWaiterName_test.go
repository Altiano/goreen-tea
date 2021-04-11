package customerOrder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/golang-boilerplate/src/domain/customerOrder/models"
	"gitlab.com/altiano/golang-boilerplate/src/mocks"
)

func TestNotifyWaiterName(t *testing.T) {
	testCases := map[string]struct {
		co         coModels.CustomerOrder
		waiterName string
		wantErr    error
	}{
		"success": {
			co: coModels.CustomerOrder{
				UsedReferralCode: 23,
			},
			waiterName: "john",
			wantErr:    nil,
		},
		"error": {
			co: coModels.CustomerOrder{
				UsedReferralCode: 23,
			},
			waiterName: "john",
			wantErr:    mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				err := m.sut.NotifyWaiterName(m.Ctx, tc.co, tc.waiterName)

				// act
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "CustomerOrderDomain.NotifyWaiterName")

			//
			m.Repo.CustomerOrder.EXPECT().UpdateWaiter(m.Ctx, tc.co, tc.waiterName).Return(tc.wantErr)
		})
	}
}
