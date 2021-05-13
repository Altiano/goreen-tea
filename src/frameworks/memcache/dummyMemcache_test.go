package memcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/altiano/goreen-tea/src/shared"
)

func TestMyFunction(t *testing.T) {
	t.Run("increment 4x", func(t *testing.T) {
		var (
			prefix  string = "altiano"
			want    int
			wantErr error
		)

		// arrange
		m := NewDummy(shared.Config{})

		for i := 0; i < 4; i++ {
			// assert
			result, err := m.Increment(prefix)

			// act
			want++
			assert.Equal(t, want, result)
			assert.Equal(t, wantErr, err)
		}
	})
}
