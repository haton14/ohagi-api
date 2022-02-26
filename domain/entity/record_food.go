package entity

import "github.com/haton14/ohagi-api/domain/value"

type RecordFood struct {
	food   Foodv2
	amount value.FoodAmount
}

func NewRecordFood(id int, name, unit string, amount float64) (*RecordFood, error) {
	food, err := NewFoodv2(id, name, unit)
	if err != nil {
		return nil, err
	}
	vamount, err := value.NewFoodAmount(amount)
	if err != nil {
		return nil, err
	}
	return &RecordFood{*food, *vamount}, nil
}

func (v RecordFood) Food() Foodv2 {
	return v.food
}

func (v RecordFood) Amount() float64 {
	return v.amount.Value()
}
