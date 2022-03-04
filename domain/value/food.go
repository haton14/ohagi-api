package value

type Food struct {
	name FoodName
	unit FoodUnit
}

func NewFood(name FoodName, unit FoodUnit) *Food {
	return &Food{name, unit}
}

func (v Food) Name() string {
	return v.name.Value()
}

func (v Food) Unit() string {
	return v.unit.Value()
}
