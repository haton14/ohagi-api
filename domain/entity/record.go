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

type Recordv2 struct {
	id            int
	recordFoods   []RecordFood
	lastUpdatedAt int64
	createdAt     int64
}

func NewRecordv2(id int, recordFood []RecordFood, lastUpdatedAt int64, createdAt int64) (Recordv2, error) {
	sort.Slice(recordFood, func(i, j int) bool { return recordFood[i].Food().ID() < recordFood[j].Food().ID() })
	return Recordv2{id, recordFood, lastUpdatedAt, createdAt}, nil
}

func (e Recordv2) ID() int {
	return e.id
}
func (e Recordv2) RecordFoods() []RecordFood {
	return e.recordFoods
}

func (e Recordv2) LastUpdatedAt() int64 {
	return e.lastUpdatedAt
}

func (e Recordv2) CreatedAt() int64 {
	return e.createdAt
}
