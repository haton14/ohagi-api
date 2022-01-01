package controller

import (
	"context"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/recordfood"
	"github.com/labstack/echo/v4"
)

type RecordIF interface {
	List(c echo.Context) error
	Create(c echo.Context) error
}

type Record struct {
	dbClient *ent.Client
}

func NewRecord(dbClient *ent.Client) RecordIF {
	return &Record{dbClient: dbClient}
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
	record := schema.Record{}
	if err := c.Bind(&record); err != nil {
		return err
	}
	recordEnt, err := r.dbClient.Record.Create().SetCreatedAt(time.Now()).SetLastUpdatedAt(time.Now()).Save(context.Background())
	if err != nil {
		c.Logger().Error("Save: ", err)
		return c.String(http.StatusInternalServerError, "Save: "+err.Error())
	}
	foodsEnt, err := r.dbClient.Food.Query().All(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}
	recordFoodBulk := make([]*ent.RecordFoodCreate, len(record.Foods))
	for i, food := range record.Foods {
		match := false
		for _, foodEnt := range foodsEnt {
			if food.Name == foodEnt.Name && food.Unit == foodEnt.Unit {
				recordFoodBulk[i] = r.dbClient.RecordFood.Create().SetRecordID(recordEnt.ID).SetFoodID(foodEnt.ID).SetAmount(food.Amount)
				match = true
				break
			}
		}
		if !match {
			foodEnt, err := r.dbClient.Food.Create().SetName(food.Name).SetUnit(food.Unit).Save(context.Background())
			if err != nil {
				c.Logger().Error("All: ", err)
				return c.String(http.StatusInternalServerError, "All: "+err.Error())
			}
			recordFoodBulk[i] = r.dbClient.RecordFood.Create().SetRecordID(recordEnt.ID).SetFoodID(foodEnt.ID).SetAmount(food.Amount)
		}
	}
	_, err = r.dbClient.RecordFood.CreateBulk(recordFoodBulk...).Save(context.Background())
	if err != nil {
		c.Logger().Error("All: ", err)
		return c.String(http.StatusInternalServerError, "All: "+err.Error())
	}

	record.ID = recordEnt.ID
	record.CreatedAt = recordEnt.CreatedAt.Unix()
	record.LastUpdatedAt = recordEnt.LastUpdatedAt.Unix()

	return c.JSON(http.StatusCreated, record)
}