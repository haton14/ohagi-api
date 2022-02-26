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

func NewFoodGetResponse(f []entity.Foodv2) (*FoodGetResponse, error) {
	response := make([]Food, 0, len(f))
	for _, ff := range f {
		resp := Food{ff.ID(), ff.Name(), ff.Unit()}
		response = append(response, resp)
	}
	return &FoodGetResponse{response}, nil
}
