package controller

import (
	"context"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/recordfood"
	"github.com/haton14/ohagi-api/usecase"
	"github.com/labstack/echo/v4"
)

type RecordIF interface {
	List(c echo.Context) error
	Create(c echo.Context) error
}

type Record struct {
	dbClient *ent.Client
	usecase  usecase.CreateRecordIF
}

func NewRecord(dbClient *ent.Client, usecase usecase.CreateRecordIF) RecordIF {
	return &Record{dbClient: dbClient, usecase: usecase}
}

func (r *Record) List(c echo.Context) error {
	rq := r.dbClient.Record.Query().Limit(50)
	recordsEnt, err := rq.All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}
	records := []schema.Record{}
	ids := []int{}
	for _, r := range recordsEnt {
		records = append(records, schema.Record{
			ID:            r.ID,
			Foods:         []schema.Food{},
			CreatedAt:     r.CreatedAt.Unix(),
			LastUpdatedAt: r.LastUpdatedAt.Unix(),
		})
		ids = append(ids, r.ID)
	}
	recordFoodsEnt, err := r.dbClient.RecordFood.Query().
		Where(func(s *sql.Selector) { sql.InInts(recordfood.FieldRecordID, ids...) }).
		All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}
	foodsEnt, err := r.dbClient.Food.Query().All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}

	for i, rc := range records {
		foods := []schema.Food{}
		for _, rf := range recordFoodsEnt {
			if rc.ID == rf.RecordID {
				for _, f := range foodsEnt {
					if rf.FoodID == f.ID {
						food := schema.Food{
							Name:   f.Name,
							Amount: rf.Amount,
							Unit:   f.Unit,
						}
						foods = append(foods, food)
						continue
					}
				}
			}
		}
		rc.Foods = append(rc.Foods, foods...)
		records[i] = rc
	}

	return c.JSON(http.StatusOK, &schema.Records{Records: records})
}

func (r *Record) Create(c echo.Context) error {
	// リクエストをもとにAPIで定義したリクエストスキーマに変換
	request, err := schema.NewRecord(c)
	if err != nil {
		c.Logger().Error("request parse: ", err)
		return c.String(http.StatusBadRequest, "request parse: "+err.Error())
	}

	// リクエストスキーマをusecaseに渡し、ドメインモデルをusecaseから受け取る
	record, err := r.usecase.Create(request, c.Logger())
	if err != nil {
		c.String(http.StatusInternalServerError, "All: "+err.Error())
	}

	// ドメインモデルをレスポンススキーマに変換する
	responseFoods := make([]schema.Food, 0, len(record.Foods()))
	for _, food := range record.Foods() {
		id := food.ID()
		f := schema.Food{ID: &id, Name: food.Name(), Amount: food.Amount(), Unit: food.Unit()}
		responseFoods = append(responseFoods, f)
	}

	response := schema.Record{ID: record.ID(), Foods: responseFoods, LastUpdatedAt: record.LastUpdatedAt(), CreatedAt: record.CreatedAt()}

	return c.JSON(http.StatusCreated, response)
}
