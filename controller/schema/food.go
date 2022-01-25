package schema

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/utility/anycast"
	"github.com/labstack/echo/v4"
)

type food struct {
	ID            *int    `json:"id,omitempty" param:"id"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount,omitempty"`
	Unit          string  `json:"unit"`
	LastUpdatedAt *int64  `json:"last_updated_at,omitempty"`
}

type FoodResponse struct {
	c echo.Context
	food
}
type FoodRequestIF interface {
	GetID() int
	GetName() string
	GetUnit() string
}
type FoodResponseIF interface {
	JSON(code int) error
	FoodRequestIF
}

func NewFoodRequest(c echo.Context) (FoodRequestIF, error) {
	s := food{}
	if err := c.Bind(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func NewFoodResponse(c echo.Context, f entity.Food) FoodResponseIF {
	return FoodResponse{c, food{ID: anycast.ToIntP(f.ID()), Name: f.Name(), Unit: f.Unit()}}
}

func (s food) GetID() int {
	return *s.ID
}

func (s food) GetName() string {
	return s.Name
}

func (s food) GetUnit() string {
	return s.Unit
}
func (s FoodResponse) JSON(code int) error {
	return s.c.JSON(code, s.food)
}
