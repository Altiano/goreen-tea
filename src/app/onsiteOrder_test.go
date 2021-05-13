package app

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func TestOnsiteOrder(t *testing.T) {
	testCases := map[string]struct {
		req       OnsiteOrderReq
		createRes coModels.CustomerOrder
		createErr error
		assistErr error
		want      OnsiteOrderRes
		wantErr   error
	}{
		"success": {
			req: OnsiteOrderReq{
				Email:        "altiano",
				Orders:       []string{"a"},
				ReferralCode: 12,
			},
			createRes: coModels.CustomerOrder{
				ID: mocks.ObjectIDDummy,
			},
			createErr: nil,
			assistErr: nil,
			want: OnsiteOrderRes{
				ID: mocks.ObjectIDDummy,
			},
			wantErr: nil,
		},
		"error, on create": {
			req: OnsiteOrderReq{
				Email:        "altiano",
				Orders:       []string{"a"},
				ReferralCode: 12,
			},
			createRes: coModels.CustomerOrder{},
			createErr: mocks.ErrDummy,
			assistErr: nil,
			want:      OnsiteOrderRes{},
			wantErr:   mocks.ErrDummy,
		},
		"error, bad request on create, refereralCode already redeemed": {
			req: OnsiteOrderReq{
				Email:        "altiano",
				Orders:       []string{"a"},
				ReferralCode: 12,
			},
			createRes: coModels.CustomerOrder{},
			createErr: coModels.ErrReferralCodeAlreadyRedeemed,
			assistErr: nil,
			want:      OnsiteOrderRes{},
			wantErr:   coModels.ErrReferralCodeAlreadyRedeemed,
		},
		"error, on assist": {
			req: OnsiteOrderReq{
				Email:        "altiano",
				Orders:       []string{"a"},
				ReferralCode: 12,
			},
			createRes: coModels.CustomerOrder{},
			createErr: nil,
			assistErr: mocks.ErrDummy,
			want:      OnsiteOrderRes{},
			wantErr:   mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupApp(t)
			var wg sync.WaitGroup

			defer func() {
				// assert
				result, err := m.sut.OnsiteOrder(m.Ctx, tc.req)
				wg.Wait()

				// act
				assert.Equal(t, tc.want, result)
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "App.OnsiteOrder")

			//
			m.Domain.CustomerOrder.EXPECT().Create(m.Ctx, coModels.CreateReq{
				Email:        tc.req.Email,
				Orders:       tc.req.Orders,
				ReferralCode: tc.req.ReferralCode,
			}).Return(tc.createRes, tc.createErr)

			if tc.createErr != nil {
				return
			}

			//
			m.Domain.AssistanceCoordinator.EXPECT().Assist(m.Ctx, tc.createRes).Return(tc.assistErr)

			if tc.assistErr != nil {
				return
			}

			//
			wg.Add(1)
			m.Domain.VisitorCounter.EXPECT().IncreaseVisits(m.Ctx).DoAndReturn(func(ctx context.Context) error {
				wg.Done()
				return nil
			})

			//
			wg.Add(1)
			m.Domain.CustomerOrder.EXPECT().NotifyRedeemedCode(m.Ctx, tc.createRes).DoAndReturn(func(ctx context.Context, newco coModels.CustomerOrder) error {
				wg.Done()
				return nil
			})
		})
	}
}
