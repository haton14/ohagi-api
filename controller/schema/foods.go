package schema

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/labstack/echo/v4"
)

type foods struct {
	Foods []food `json:"foods"`
}

type FoodsResponseIF interface {
	JSON(code int) error
}
type FoodsResponse struct {
	c echo.Context
	foods
}

func NewFoodsResponse(c echo.Context, f []entity.Food) FoodsResponseIF {
	response := make([]food, 0, len(f))
	for _, v := range f {
		id := v.ID()
		food := food{ID: &id, Name: v.Name(), Unit: v.Unit()}
		response = append(response, food)
	}
	responses := foods{response}
	return FoodsResponse{c: c, foods: responses}
}

func (s FoodsResponse) JSON(code int) error {
	return s.c.JSON(code, s.foods)
}
