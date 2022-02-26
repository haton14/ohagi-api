package value

import "fmt"

type FoodAmount struct {
	value float64
}

func NewFoodAmount(v float64) (*FoodAmount, error) {
	if v < 0.0 {
		return nil, fmt.Errorf("[%w]食事量が負数", ErrMinRange)
	}
	return &FoodAmount{v}, nil
}

func (v FoodAmount) Value() float64 {
	return v.value
}
