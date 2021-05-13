package visitorCounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/altiano/goreen-tea/src/mocks"
)

func TestIncreaseVisits(t *testing.T) {
	testCases := map[string]struct {
		incrementKey    string
		incrementResult int
		incrementErr    error
		want            string
		wantErr         error
	}{
		"success": {
			incrementResult: 3,
			incrementErr:    nil,
			want:            "",
			wantErr:         nil,
		},
		"error": {
			incrementResult: 0,
			incrementErr:    mocks.ErrDummy,
			want:            "",
			wantErr:         mocks.ErrDummy,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			m := setupDomain(t)

			defer func() {
				// assert
				err := m.sut.IncreaseVisits(m.Ctx)

				// act
				assert.Equal(t, tc.wantErr, err)
			}()

			// arrange
			m.ExpectSpan(m.Ctx, "VisitorCounterDomain.IncreaseVisits")

			//
			m.Memcache.Manager.EXPECT().Increment("visitorCounter").
				Return(tc.incrementResult, tc.incrementErr)

		})
	}
}
