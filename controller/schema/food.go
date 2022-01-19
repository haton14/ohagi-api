package schema

import (
	"github.com/labstack/echo/v4"
)

type food struct {
	ID            *int    `json:"id,omitempty"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount,omitempty"`
	Unit          string  `json:"unit"`
	LastUpdatedAt *int64  `json:"last_updated_at,omitempty"`
}

type FoodRequestIF interface {
	GetName() string
}

func NewFoodRequest(c echo.Context) (FoodRequestIF, error) {
	s := food{}
	if err := c.Bind(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func (s food) GetName() string {
	return s.Name
}
