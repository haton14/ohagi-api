package controller

import (
	"net/http"

	"github.com/haton14/ohagi-api/controller/response"
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

type FoodIF interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
}

type Food struct {
	usecase usecase.Food
}

func NewFood(usecase usecase.Food) FoodIF {
	return &Food{usecase: usecase}
}
func (f *Food) Create(c echo.Context) error {
	// リクエストをもとにAPIで定義したリクエストスキーマに変換
	request, err := schema.NewFoodRequest(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}
	// リクエストスキーマをusecaseに渡し、ドメインモデルをusecaseから受け取る
	food, err := f.usecase.Create(request, c.Logger())

	// ドメインモデルをレスポンススキーマに変換する
	response := schema.NewFoodResponse(c, food)
	return response.JSON(http.StatusCreated)
}

func (f *Food) List(c echo.Context) error {
	// ドメインモデルをusecaseから受け取る
	foods, errResp := f.usecase.List(c.Logger())
	if errResp != nil {
		return c.JSON(errResp.HttpStatus, errResp)
	}

	// ドメインモデルをレスポンススキーマに変換する
	resp, err := response.NewFoodGetResponse(foods)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "予期しないエラー"})
	}
	return c.JSON(http.StatusOK, resp)
}

func (f *Food) Update(c echo.Context) error {
	// リクエストをもとにAPIで定義したリクエストスキーマに変換
	request, err := schema.NewFoodRequest(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}
	// リクエストスキーマをusecaseに渡し、ドメインモデルをusecaseから受け取る
	food, err := f.usecase.Update(request, c.Logger())

	// ドメインモデルをレスポンススキーマに変換する
	response := schema.NewFoodResponse(c, food)
	return response.JSON(http.StatusCreated)
}
