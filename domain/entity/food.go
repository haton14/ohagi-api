package entity

type Food struct {
	id     int
	name   string
	amount float64
	unit   string
}

func NewFood(id int, name string, amount float64, unit string) (Food, error) {
	return Food{id, name, amount, unit}, nil
}

func (e Food) ID() int {
	return e.id
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
