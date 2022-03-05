package request

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/domain/value"
	"github.com/labstack/echo/v4"
)

type recordsPost struct {
	Foods []recordPostFood `json:"foods"`
}

type recordPostFood struct {
	ID     *int    `json:"id"`
	Name   string  `json:"name"`
	Unit   string  `json:"unit"`
	Amount float64 `json:"amount"`
}

func NewRecordsPost(c echo.Context) ([]entity.FoodContent, []value.FoodContent, error) {
	request := &recordsPost{}
	if err := c.Bind(request); err != nil {
		return nil, nil, err
	}
	foodContentsEntity := []entity.FoodContent{}
	foodContentsValue := []value.FoodContent{}
	for _, food := range request.Foods {
		name, err := value.NewFoodName(food.Name)
		if err != nil {
			return nil, nil, err
		}
		unit, err := value.NewFoodUnit(food.Unit)
		if err != nil {
			return nil, nil, err
		}
		amount, err := value.NewFoodAmount(food.Amount)
		if err != nil {
			return nil, nil, err
		}
		if food.ID != nil {
			id, err := value.NewID(*food.ID)
			if err != nil {
				return nil, nil, err
			}
			food := entity.NewFoodContent(*entity.NewFoodv3(*id, *value.NewFood(*name, *unit)), *amount)
			foodContentsEntity = append(foodContentsEntity, *food)
		} else {
			food := value.NewFoodContent(*value.NewFood(*name, *unit), *amount)
			foodContentsValue = append(foodContentsValue, *food)
		}
	}
	return foodContentsEntity, foodContentsValue, nil
}
