package entity

import (
	"errors"

	"github.com/haton14/ohagi-api/domain/value"
)

type Food struct {
	id     int
	name   string
	amount float64
	unit   string
}

func NewFood(id int, name string, amount float64, unit string) (Food, error) {
	if id < 0 {
		return Food{}, errors.New("id isn't negative number")
	}
	if name == "" {
		return Food{}, errors.New("name isn't empty")
	}
	if amount < 0 {
		return Food{}, errors.New("amount isn't negative number")
	}
	if unit == "" {
		return Food{}, errors.New("unit isn't empty")
	}
	return Food{id, name, amount, unit}, nil
}

func (e Food) ID() int {
	return e.id
}

func (e *Food) SetID(id int) {
	e.id = id
}

func (e Food) Name() string {
	return e.name
}

func (e Food) Amount() float64 {
	return e.amount
}

func (e Food) Unit() string {
	return e.unit
}

func (e Food) IsNullID() bool {
	return e.ID() == 0
}

func (e Food) IsNullAmount() bool {
	return e.Amount() == 0
}

type Foodv2 struct {
	id   value.ID
	name value.FoodName
	unit value.FoodUnit
}

func NewFoodv2(id int, name, unit string) (*Foodv2, error) {
	vid, err := value.NewID(id)
	if err != nil {
		return nil, err
	}
	vname, err := value.NewFoodName(name)
	if err != nil {
		return nil, err
	}
	vunit, err := value.NewFoodUnit(unit)
	if err != nil {
		return nil, err
	}
	return &Foodv2{*vid, *vname, *vunit}, nil
}
func (v Foodv2) ID() int {
	return v.id.Value()
}

func (v Foodv2) Name() string {
	return v.name.Value()
}

func (v Foodv2) Unit() string {
	return v.unit.Value()
}

type Foodv3 struct {
	id    value.ID
	value value.Food
}

func NewFoodv3(id value.ID, value value.Food) *Foodv3 {
	return &Foodv3{id, value}
}

func (v Foodv3) ID() value.ID {
	return v.id
}

func (v Foodv3) Value() value.Food {
	return v.value
}
