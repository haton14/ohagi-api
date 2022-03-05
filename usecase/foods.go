package usecase

import (
	"errors"
	"net/http"

	"github.com/haton14/ohagi-api/controller/response"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/domain/value"
	"github.com/haton14/ohagi-api/repository"
	"github.com/labstack/echo/v4"
)

type CreateFoodIF interface {
	Create(food value.Food, logger echo.Logger) (*entity.Foodv3, *response.ErrorResponse)
}

type ListFoodIF interface {
	List(logger echo.Logger) ([]entity.Foodv3, *response.ErrorResponse)
}

type UpdateFoodIF interface {
	Update(food entity.Foodv3, logger echo.Logger) (*entity.Foodv3, *response.ErrorResponse)
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

func (u CreateFood) Create(f value.Food, logger echo.Logger) (*entity.Foodv3, *response.ErrorResponse) {
	conflict, err := u.foodRepo.FindByNameUnitV2(f)
	if len(conflict) > 0 {
		return nil, &response.ErrorResponse{Message: "登録しようとした食事は既に存在", HttpStatus: http.StatusConflict}
	}
	food, err := u.foodRepo.SaveV2(f)
	if err != nil {
		logger.Error("%w;foodRepo.SaveV2()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}
	return food, nil
}

func (u ListFood) List(logger echo.Logger) ([]entity.Foodv3, *response.ErrorResponse) {
	foods, err := u.foodRepo.ListV2()
	if errors.Is(err, repository.ErrNotFoundRecord) {
		logger.Warn("%w;foodRepo.List()でエラー", err)
		return nil, &response.ErrorResponse{Message: "データが存在しない", HttpStatus: http.StatusNotFound}
	} else if err != nil {
		logger.Error("%w;foodRepo.List()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}
	return foods, nil
}

func (u UpdateFood) Update(food entity.Foodv3, logger echo.Logger) (*entity.Foodv3, *response.ErrorResponse) {
	_, err := u.foodRepo.FindByIDV2(food.ID())
	if errors.Is(err, repository.ErrNotFoundRecord) {
		logger.Error("%w;foodRepo.FindByID()でエラー", err)
		return nil, &response.ErrorResponse{Message: "データが存在しない", HttpStatus: http.StatusNotFound}
	} else if err != nil {
		logger.Error("%w;foodRepo.FindByID()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}
	updateFood, err := u.foodRepo.UpdateNameUnitFindByIDV2(food)
	if err != nil {
		logger.Error("%w;foodRepo.UpdateNameUnitFindByIDV2()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}
	return updateFood, nil
}
