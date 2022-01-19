package controller

import (
	"net/http"

	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

type FoodIF interface {
	Create(c echo.Context) error
}

type Food struct {
	usecase usecase.Food
}

func NewFood(usecase usecase.Food) FoodIF {
	return &Food{usecase: usecase}
}
func (f *Food) Create(c echo.Context) error {
	// リクエストをもとにAPIで定義したリクエストスキーマに変換
	_, err := schema.NewFoodRequest(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}
	return c.NoContent(http.StatusOK)
}
