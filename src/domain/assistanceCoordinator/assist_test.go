package assistanceCoordinator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	waiterModels "gitlab.com/altiano/goreen-tea/src/domain/waiter/models"
	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func TestAssist(t *testing.T) {
	testCases := map[string]struct {
		co                  coModels.CustomerOrder
		pickWaiterRes       waiterModels.Waiter
		pickNormalWaiterErr error
		pickTopWaiterErr    error

		notifyWaiterNameErr   error
		increaseTotalServeErr error

		wantErr error
	}{
		"success, low priority": {
			co: coModels.CustomerOrder{},

			pickWaiterRes: waiterModels.Waiter{
				Name: "jeana",
			},
			pickNormalWaiterErr: nil,
			notifyWaiterNameErr: nil,

			increaseTotalServeErr: nil,

			wantErr: nil,
		},

		"success, high priority": {
			co: coModels.CustomerOrder{
				Orders: []string{"A", "B", "C", "D"},
			},

			pickWaiterRes: waiterModels.Waiter{
				Name: "jeana",
			},
			pickNormalWaiterErr: nil,
			notifyWaiterNameErr: nil,

			increaseTotalServeErr: nil,

			wantErr: nil,
		},

		"error, on pickNormalWaiter": {
			co:                  coModels.CustomerOrder{},
			pickWaiterRes:       waiterModels.Waiter{},
			pickNormalWaiterErr: mocks.ErrDummy,
			wantErr:             mocks.ErrDummy,
		},

		"error, on pickTopWaiter": {
			co: coModels.CustomerOrder{
				Orders: []string{"a", "b", "c", "d"},
			},
			pickWaiterRes:    waiterModels.Waiter{},
			pickTopWaiterErr: mocks.ErrDummy,
			wantErr:          mocks.ErrDummy,
		},

		"error, on notifyWaiterName": {
			co: coModels.CustomerOrder{},

			pickWaiterRes: waiterModels.Waiter{
				Name: "jeana",
			},
			pickNormalWaiterErr:   nil,
			notifyWaiterNameErr:   mocks.ErrDummy,
			increaseTotalServeErr: nil,

			wantErr: mocks.ErrDummy,
		},

		"error, on increaseTotalServeErr": {
			co: coModels.CustomerOrder{},

			pickWaiterRes: waiterModels.Waiter{
				Name: "jeana",
			},
			pickNormalWaiterErr:   nil,
			notifyWaiterNameErr:   nil,
			increaseTotalServeErr: mocks.ErrDummy,

			wantErr: mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				err := m.sut.Assist(m.Ctx, tc.co)

				// act
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "AssistanceCoordinator.Assist")

			//
			if tc.co.IsHighPriority() {
				m.Domain.Waiter.EXPECT().PickTopWaiter(m.Ctx).Return(tc.pickWaiterRes, tc.pickTopWaiterErr)
				if tc.pickTopWaiterErr != nil {
					return
				}
			} else {
				m.Domain.Waiter.EXPECT().PickNormalWaiter(m.Ctx).Return(tc.pickWaiterRes, tc.pickNormalWaiterErr)
				if tc.pickNormalWaiterErr != nil {
					return
				}
			}

			//
			m.Domain.CustomerOrder.EXPECT().NotifyWaiterName(m.Ctx, tc.co, tc.pickWaiterRes.Name).Return(tc.notifyWaiterNameErr)
			if tc.notifyWaiterNameErr != nil {
				return
			}

			//
			m.Domain.Waiter.EXPECT().IncreaseTotalServe(m.Ctx, tc.pickWaiterRes).Return(tc.increaseTotalServeErr)
		})
	}
}
