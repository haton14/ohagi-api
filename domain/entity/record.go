package entity

import "sort"

type Record struct {
	id            int
	foods         []Food
	lastUpdatedAt int64
	createdAt     int64
}

func NewRecord(id int, foods []Food, lastUpdatedAt int64, createdAt int64) (Record, error) {
	sort.Slice(foods, func(i, j int) bool { return foods[i].ID() < foods[j].ID() })
	return Record{id, foods, lastUpdatedAt, createdAt}, nil
}

func (e Record) ID() int {
	return e.id
}
func (e Record) Foods() []Food {
	return e.foods
}

func (e Record) LastUpdatedAt() int64 {
	return e.lastUpdatedAt
}

func (e Record) CreatedAt() int64 {
	return e.createdAt
}
