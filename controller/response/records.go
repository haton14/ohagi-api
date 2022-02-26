package response

import "github.com/haton14/ohagi-api/domain/entity"

type RecordGetResponse struct {
	Records []Record `json:"records"`
}

type Record struct {
	ID            int          `json:"id"`
	Foods         []RecordFood `json:"foods"`
	LastUpdatedAt int64        `json:"last_updated_at"`
	CreatedAt     int64        `json:"created_at"`
}

type RecordFood struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Unit   string  `json:"unit"`
	Amount float64 `json:"amount"`
}

func NewRecordGetResponse(r []entity.Recordv2) (*RecordGetResponse, error) {
	response := make([]Record, 0, len(r))
	for _, rr := range r {
		foods := make([]RecordFood, 0, len(rr.RecordFoods()))
		for _, f := range rr.RecordFoods() {
			food := RecordFood{f.Food().ID(), f.Food().Name(), f.Food().Unit(), f.Amount()}
			foods = append(foods, food)
		}
		record := Record{rr.ID(), foods, rr.LastUpdatedAt(), rr.CreatedAt()}
		response = append(response, record)
	}
	return &RecordGetResponse{response}, nil
}
