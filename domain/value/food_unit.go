package value

import "fmt"

type FoodUnit struct {
	value string
}

func NewFoodUnit(v string) (*FoodUnit, error) {
	if len(v) < 1 {
		return nil, fmt.Errorf("[%w]食事単位が空文字列", ErrMinLength)
	}
	if len(v) > 50 {
		return nil, fmt.Errorf("[%w]食事単位が50文字より長い", ErrMaxLength)
	}
	return &FoodUnit{v}, nil
}

func (v FoodUnit) Value() string {
	return v.value
}
