package response

import "github.com/haton14/ohagi-api/domain/entity"

type FoodGetResponse struct {
	Foods []Food `json:"foods"`
}

type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func NewFoodGetResponse(foods []entity.Foodv3) (*FoodGetResponse, error) {
	response := make([]Food, 0, len(foods))
	for _, food := range foods {
		resp := Food{food.ID(), food.Name(), food.Unit()}
		response = append(response, resp)
	}
	return &FoodGetResponse{response}, nil
}
