package value

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	cases := []struct {
		Name    string
		arg     int
		want    int
		wantErr error
	}{
		{"[正常]1", 1, 1, nil},
		{"[正常]int最大値", math.MaxInt, math.MaxInt, nil},
		{"[正常]0", 0, 0, nil},
		{"[異常]-1", -1, -1, ErrMinRange},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			have, err := NewID(c.arg)
			// check
			if c.wantErr != nil {
				assert.Error(t, err)
				assert.Nil(t, have)
				assert.ErrorIs(t, err, c.wantErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, have)
				assert.Equal(t, c.want, have.Value())
			}
		})
	}
}
