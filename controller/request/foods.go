package request

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/domain/value"
	"github.com/labstack/echo/v4"
)

type foodsPatch struct {
	ID   int    `param:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func NewFoodsPatch(c echo.Context) (*entity.Foodv3, error) {
	request := &foodsPatch{}
	if err := c.Bind(request); err != nil {
		return nil, err
	}
	id, err := value.NewID(request.ID)
	if err != nil {
		return nil, err
	}
	name, err := value.NewFoodName(request.Name)
	if err != nil {
		return nil, err
	}
	unit, err := value.NewFoodUnit(request.Name)
	if err != nil {
		return nil, err
	}
	return entity.NewFoodv3(*id, *value.NewFood(*name, *unit)), nil
}
