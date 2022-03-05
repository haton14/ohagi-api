package value

type FoodContent struct {
	food   Food
	amount FoodAmount
}

func NewFoodContent(food Food, amount FoodAmount) *FoodContent {
	return &FoodContent{food: food, amount: amount}
}

func (f FoodContent) Amont() FoodAmount {
	return f.amount
}

func (f FoodContent) Food() Food {
	return f.food
}
