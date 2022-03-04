package response

import "github.com/haton14/ohagi-api/domain/entity"

type FoodGetResponse struct {
	Foods []Food `json:"foods"`
}

type FoodsPost Food
type FoodsPatch Food
type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func NewFoodGetResponse(foods []entity.Foodv3) (*FoodGetResponse, error) {
	response := make([]Food, 0, len(foods))
	for _, food := range foods {
		resp := Food{food.ID().Value(), food.Value().Name(), food.Value().Unit()}
		response = append(response, resp)
	}
	return &FoodGetResponse{response}, nil
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
