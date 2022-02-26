package value

import "fmt"

type FoodName struct {
	value string
}

func NewFoodName(v string) (*FoodName, error) {
	if len(v) < 1 {
		return nil, fmt.Errorf("[%w]食事名が空文字列", ErrMinLength)
	}
	if len(v) > 100 {
		return nil, fmt.Errorf("[%w]食事名が100文字より長い", ErrMaxLength)
	}
	return &FoodName{v}, nil
}

func (v FoodName) Value() string {
	return v.value
}
