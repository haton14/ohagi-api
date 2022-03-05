package entity

import "github.com/haton14/ohagi-api/domain/value"

type FoodContent struct {
	foodID value.ID
	food   value.Food
	amount value.FoodAmount
}

func NewFoodContent(food Foodv3, amount value.FoodAmount) *FoodContent {
	return &FoodContent{foodID: food.ID(), food: food.Value(), amount: amount}
}

func (f *FoodContent) ID() int {
	return f.foodID.Value()
}

func (f *FoodContent) Name() string {
	return f.food.Name()
}

func (f *FoodContent) Amont() float64 {
	return f.Amont()
}

func (f *FoodContent) Unit() string {
	return f.food.Unit()
}
