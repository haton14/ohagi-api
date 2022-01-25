package usecase

import (
	"fmt"

	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/repository"
	"github.com/labstack/echo/v4"
)

type CreateFoodIF interface {
	Create(request schema.FoodRequestIF, logger echo.Logger) (entity.Food, error)
}

type ListFoodIF interface {
	List(logger echo.Logger) ([]entity.Food, error)
}

type UpdateFoodIF interface {
	Update(request schema.FoodRequestIF, logger echo.Logger) (entity.Food, error)
}
type Food struct {
	CreateFoodIF
	ListFoodIF
	UpdateFoodIF
}

type CreateFood struct {
	foodRepo repository.FoodIF
}
type ListFood struct {
	foodRepo repository.FoodIF
}

type UpdateFood struct {
	foodRepo repository.FoodIF
}

func NewFood(foodRepo repository.FoodIF) Food {
	return Food{
		CreateFood{foodRepo: foodRepo},
		ListFood{foodRepo: foodRepo},
		UpdateFood{foodRepo: foodRepo},
	}
}

func (u CreateFood) Create(request schema.FoodRequestIF, logger echo.Logger) (entity.Food, error) {
	food, err := entity.NewFood(0, request.GetName(), 0, request.GetUnit())
	if err != nil {
		return entity.Food{}, err
	}
	conflict, err := u.foodRepo.FindByNameUnit(food.Name(), food.Unit())
	if conflict != nil {
		return entity.Food{}, fmt.Errorf("conflict food, name: %s, unit: %s", food.Name(), food.Unit())
	}
	err = u.foodRepo.Save(&food)
	if err != nil {
		return entity.Food{}, fmt.Errorf("food save err: %s", err)
	}
	return food, nil
}

func (u ListFood) List(logger echo.Logger) ([]entity.Food, error) {
	foods, err := u.foodRepo.List()
	if err != nil {
		return nil, fmt.Errorf("foods list err: %s", err)
	}

	return foods, nil
}

func (u UpdateFood) Update(request schema.FoodRequestIF, logger echo.Logger) (entity.Food, error) {
	food, err := entity.NewFood(request.GetID(), request.GetName(), 0, request.GetUnit())
	if err != nil {
		return entity.Food{}, err
	}
	conflict, err := u.foodRepo.FindByID(food.ID())
	if conflict == nil {
		return entity.Food{}, fmt.Errorf("not exit food. id:%d, name: %s, unit: %s", food.ID(), food.Name(), food.Unit())
	}
	updateFood, err := u.foodRepo.UpdateNameUnitFindByID(food.Name(), food.Unit(), food.ID())
	if err != nil {
		return entity.Food{}, fmt.Errorf("food update err: %s", err)
	}
	return *updateFood, nil
}
