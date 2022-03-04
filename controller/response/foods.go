package response

import "github.com/haton14/ohagi-api/domain/entity"

type FoodsGet struct {
	Foods []Food `json:"foods"`
}

type FoodsPost Food

type FoodsPatch Food
type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func NewFoodsGet(ff []entity.Foodv3) (*FoodsGet, error) {
	foods := make([]Food, 0, len(ff))
	for _, f := range ff {
		food := Food{
			f.ID().Value(),
			f.Value().Name(),
			f.Value().Unit(),
		}
		foods = append(foods, food)
	}
	return &FoodsGet{
		Foods: foods,
	}, nil
}

func NewFoodsPost(food entity.Foodv3) *FoodsPost {
	return &FoodsPost{
		ID:   food.ID().Value(),
		Name: food.Value().Name(),
		Unit: food.Value().Unit(),
	}
}

func NewFoodsPatch(food entity.Foodv3) *FoodsPatch {
	return &FoodsPatch{
		ID:   food.ID().Value(),
		Name: food.Value().Name(),
		Unit: food.Value().Unit(),
	}
}
