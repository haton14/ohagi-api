package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFood(t *testing.T) {
	name, _ := NewFoodName("ハム太郎")
	unit, _ := NewFoodUnit("kg")
	type args struct {
		name FoodName
		unit FoodUnit
	}
	type want struct {
		name string
		unit string
	}
	cases := []struct {
		Name string
		args args
		want want
	}{
		{"[正常]ハム太郎 kg", args{*name, *unit}, want{"ハム太郎", "kg"}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			have := NewFood(c.args.name, c.args.unit)
			// check
			assert.NotNil(t, have)
			assert.Equal(t, c.want.name, have.name.Value())
			assert.Equal(t, c.want.unit, have.unit.Value())

		})
	}
}
