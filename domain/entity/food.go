package entity

import "errors"

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
