package entity

import (
	"github.com/haton14/ohagi-api/domain/value"
)

type Food struct {
	id    value.ID
	value value.Food
}

func NewFoodv3(id value.ID, value value.Food) *Food {
	return &Food{id, value}
}

func (v Food) ID() value.ID {
	return v.id
}

func (v Food) Value() value.Food {
	return v.value
}
