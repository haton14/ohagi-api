package controller

import (
	"net/http"

	"github.com/haton14/ohagi-api/controller/response"
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

type RecordIF interface {
	List(c echo.Context) error
	Create(c echo.Context) error
}

type Record struct {
	dbClient *ent.Client
	usecase  usecase.Record
}

func NewRecord(dbClient *ent.Client, usecase usecase.Record) RecordIF {
	return &Record{dbClient: dbClient, usecase: usecase}
}

func (r *Record) List(c echo.Context) error {
	// ドメインモデルをusecaseから受け取る
	records, errResp := r.usecase.List(c.Logger())
	if errResp != nil {
		return c.JSON(errResp.HttpStatus, errResp)
	}
	// ドメインモデルをレスポンススキーマに変換する
	resp, err := response.NewRecordGetResponse(records)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "予期しないエラー"})
	}
	return c.JSON(http.StatusOK, resp)
}

func (r *Record) Create(c echo.Context) error {
	// リクエストをもとにAPIで定義したリクエストスキーマに変換
	request, err := schema.NewRecordRequest(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}

	// リクエストスキーマをusecaseに渡し、ドメインモデルをusecaseから受け取る
	record, err := r.usecase.Create(request, c.Logger())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	// ドメインモデルをレスポンススキーマに変換する
	response := schema.NewRecordResponse(c, record)
	return response.JSON(http.StatusCreated)
}
