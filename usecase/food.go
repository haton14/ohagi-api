package usecase

import (
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/labstack/echo/v4"
)

type CreateFoodIF interface {
	Create(request schema.FoodRequestIF, logger echo.Logger) error
}

type Food struct {
	CreateFoodIF
}

type CreateFood struct {
}

func NewFood() Food {
	return Food{
		CreateFood{},
	}
}

func (u CreateFood) Create(request schema.FoodRequestIF, logger echo.Logger) error {
	_, _ = entity.NewFood(0, request.GetName(), 0, request.GetUnit())
	return nil
}
