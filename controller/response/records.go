package response

import "github.com/haton14/ohagi-api/domain/entity"

type RecordsGet struct {
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

func NewRecordsGet(r []entity.Record) (*RecordsGet, error) {
	response := make([]Record, 0, len(r))
	for _, rr := range r {
		foodContents := make([]FoodContent, 0, rr.LenFoodContent())
		for _, f := range rr.FoodContents() {
			foodContent := FoodContent{f.ID(), f.Name(), f.Unit(), f.Amont()}
			foodContents = append(foodContents, foodContent)
		}
		record := Record{rr.ID(), foodContents, rr.LastUpdatedAt(), rr.CreatedAt()}
		response = append(response, record)
	}
	return &RecordsGet{response}, nil
}

func NewRecordsPost(r entity.Record) (*RecordsPost, error) {
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
