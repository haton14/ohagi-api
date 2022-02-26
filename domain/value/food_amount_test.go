package value

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodAmount(t *testing.T) {
	cases := []struct {
		Name    string
		arg     float64
		want    float64
		wantErr error
	}{
		{"[正常]1", 1, 1, nil},
		{"[正常]float64最大値", 1.2, 1.2, nil},
		{"[正常]float64の0に最も近い小数", math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, nil},
		{"[正常]0", 0, 0, nil},
		{"[異常]-1", -1, -1, ErrMinRange},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			have, err := NewFoodAmount(c.arg)
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
