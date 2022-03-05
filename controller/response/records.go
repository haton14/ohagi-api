package response

import "github.com/haton14/ohagi-api/domain/entity"

type RecordGetResponse struct {
	Records []Record `json:"records"`
}

type RecordsPost Record
type Record struct {
	ID            int           `json:"id"`
	FoodContents  []FoodContent `json:"foods"`
	LastUpdatedAt int64         `json:"last_updated_at"`
	CreatedAt     int64         `json:"created_at"`
}

type FoodContent struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Unit   string  `json:"unit"`
	Amount float64 `json:"amount"`
}

func NewRecordGetResponse(r []entity.Recordv2) (*RecordGetResponse, error) {
	response := make([]Record, 0, len(r))
	for _, rr := range r {
		foods := make([]FoodContent, 0, len(rr.RecordFoods()))
		for _, f := range rr.RecordFoods() {
			food := FoodContent{f.Food().ID(), f.Food().Name(), f.Food().Unit(), f.Amount()}
			foods = append(foods, food)
		}
		record := Record{rr.ID(), foods, rr.LastUpdatedAt(), rr.CreatedAt()}
		response = append(response, record)
	}
	return &RecordGetResponse{response}, nil
}

func NewRecordsPost(r entity.Recordv3) (*RecordsPost, error) {
	foodContents := make([]FoodContent, 0, r.LenFoodContent())
	for _, f := range r.FoodContents() {
		foodContent := FoodContent{f.ID(), f.Name(), f.Unit(), f.Amont()}
		foodContents = append(foodContents, foodContent)
	}
	return &RecordsPost{
		ID:            r.ID(),
		FoodContents:  foodContents,
		LastUpdatedAt: r.LastUpdatedAt(),
		CreatedAt:     r.CreatedAt(),
	}, nil
}
