package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodUnit(t *testing.T) {
	string50 := "123456789|123456789|123456789|123456789|123456789|"
	cases := []struct {
		Name    string
		arg     string
		want    string
		wantErr error
	}{
		{"[正常]kg", "kg", "kg", nil},
		{"[正常]50文字", string50, string50, nil},
		{"[異常]空文字", "", "-", ErrMinLength},
		{"[異常]51文字", string50 + "a", "-", ErrMaxLength},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			have, err := NewFoodUnit(c.arg)
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
