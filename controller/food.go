package controller

import (
	"net/http"

	"github.com/haton14/ohagi-api/controller/request"
	"github.com/haton14/ohagi-api/controller/response"
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
	req, err := request.NewFoodsPost(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}
	// リクエストスキーマをusecaseに渡し、ドメインモデルをusecaseから受け取る
	food, errResp := f.usecase.Create(*req, c.Logger())
	if errResp != nil {
		return c.JSON(errResp.HttpStatus, errResp)
	}

	// ドメインモデルをレスポンススキーマに変換する
	resp := response.NewFoodsPost(*food)
	return c.JSON(http.StatusCreated, resp)
}

func (f *Food) List(c echo.Context) error {
	// ドメインモデルをusecaseから受け取る
	foods, errResp := f.usecase.List(c.Logger())
	if errResp != nil {
		return c.JSON(errResp.HttpStatus, errResp)
	}

	// ドメインモデルをレスポンススキーマに変換する
	resp, err := response.NewFoodsGet(foods)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "予期しないエラー"})
	}
	return c.JSON(http.StatusOK, resp)
}

func (f *Food) Update(c echo.Context) error {
	// リクエストをもとにAPIで定義したリクエストスキーマに変換
	req, err := request.NewFoodsPatch(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}
	// リクエストスキーマをusecaseに渡し、ドメインモデルをusecaseから受け取る
	food, errResp := f.usecase.Update(*req, c.Logger())
	if errResp != nil {
		return c.JSON(errResp.HttpStatus, errResp)
	}
	// ドメインモデルをレスポンススキーマに変換する
	resp := response.NewFoodsPatch(*food)
	return c.JSON(http.StatusCreated, resp)
}
