package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodName(t *testing.T) {
	string100 := "123456789|123456789|123456789|123456789|123456789|123456789|123456789|123456789|123456789|123456789|"
	cases := []struct {
		Name    string
		arg     string
		want    string
		wantErr error
	}{
		{"[正常]ミルワーム", "ミルワーム", "ミルワーム", nil},
		{"[正常]100文字", string100, string100, nil},
		{"[異常]空文字", "", "-", ErrMinLength},
		{"[異常]101文字", string100 + "a", "-", ErrMaxLength},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			have, err := NewFoodName(c.arg)
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
