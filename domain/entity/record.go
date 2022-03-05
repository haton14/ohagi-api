package entity

import (
	"github.com/haton14/ohagi-api/domain/value"
)

type Record struct {
	id            value.ID
	lastUpdatedAt int64
	createdAt     int64
	foodContents  []FoodContent
}

func NewRecordv3(id value.ID, last, create int64) *Record {
	foodContents := []FoodContent{}
	return &Record{
		id:            id,
		lastUpdatedAt: last,
		createdAt:     create,
		foodContents:  foodContents,
	}
}

func (r Record) ID() int {
	return r.id.Value()
}

func (r Record) LastUpdatedAt() int64 {
	return r.lastUpdatedAt
}

func (r Record) CreatedAt() int64 {
	return r.createdAt
}

func (r *Record) AddFoodContent(food Food, amount value.FoodAmount) {
	foodContent := NewFoodContent(food, amount)
	r.foodContents = append(r.foodContents, *foodContent)
}

func (r Record) LenFoodContent() int {
	return len(r.foodContents)
}

func (r Record) FoodContents() []FoodContent {
	return r.foodContents
}
