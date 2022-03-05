package entity

import "github.com/haton14/ohagi-api/domain/value"

type FoodContent struct {
	foodID value.ID
	food   value.Food
	amount value.FoodAmount
}

func NewFoodContent(food Food, amount value.FoodAmount) *FoodContent {
	return &FoodContent{foodID: food.ID(), food: food.Value(), amount: amount}
}

func (f FoodContent) ID() value.ID {
	return f.foodID
}

func (f FoodContent) Food() value.Food {
	return f.food
}

func (f FoodContent) Amont() value.FoodAmount {
	return f.amount
}
