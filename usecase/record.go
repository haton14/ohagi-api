package usecase

import (
	"context"
	"fmt"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/recordfood"
	"github.com/labstack/echo/v4"
)

type CreateRecordIF interface {
	Create(request schema.RecordRequestIF, logger echo.Logger) (entity.Record, error)
}

type ListRecordIF interface {
	List(logger echo.Logger) ([]entity.Record, error)
}

type Record struct {
	CreateRecordIF
	ListRecordIF
}

type CreateRecord struct {
	dbClient *ent.Client
}

type ListRecord struct {
	dbClient *ent.Client
}

func NewRecord(dbClient *ent.Client) Record {
	return Record{
		CreateRecord{dbClient: dbClient},
		ListRecord{dbClient: dbClient},
	}
}

func (u CreateRecord) Create(request schema.RecordRequestIF, logger echo.Logger) (entity.Record, error) {
	recordEnt, err := u.dbClient.Record.Create().SetCreatedAt(time.Now()).SetLastUpdatedAt(time.Now()).Save(context.Background())
	if err != nil {
		logger.Error("Save: ", err)
		return entity.Record{}, fmt.Errorf("Save: %w", err)
	}
	foodsEnt, err := u.dbClient.Food.Query().All(context.Background())
	if err != nil {
		logger.Error("All: ", err)
		return entity.Record{}, fmt.Errorf("All: %w", err)
	}
	recordFoodBulk := make([]*ent.RecordFoodCreate, len(request.GetFoods()))
	foods := make([]entity.Food, 0, len(request.GetFoods()))
	for i, food := range request.GetFoods() {
		match := false
		for _, foodEnt := range foodsEnt {
			if food.Name == foodEnt.Name && food.Unit == foodEnt.Unit {
				recordFoodBulk[i] = u.dbClient.RecordFood.Create().SetRecordID(recordEnt.ID).SetFoodID(foodEnt.ID).SetAmount(food.Amount)
				foodE, _ := entity.NewFood(foodEnt.ID, foodEnt.Name, food.Amount, foodEnt.Unit)
				foods = append(foods, foodE)
				match = true
				break
			}
		}
		if !match {
			foodEnt, err := u.dbClient.Food.Create().SetName(food.Name).SetUnit(food.Unit).Save(context.Background())
			if err != nil {
				logger.Error("All: ", err)
				return entity.Record{}, fmt.Errorf("All: %w", err)
			}
			recordFoodBulk[i] = u.dbClient.RecordFood.Create().SetRecordID(recordEnt.ID).SetFoodID(foodEnt.ID).SetAmount(food.Amount)
			foodE, _ := entity.NewFood(foodEnt.ID, foodEnt.Name, food.Amount, foodEnt.Unit)
			foods = append(foods, foodE)
		}
	}
	_, err = u.dbClient.RecordFood.CreateBulk(recordFoodBulk...).Save(context.Background())
	if err != nil {
		logger.Error("All: ", err)
		return entity.Record{}, fmt.Errorf("All: %w", err)
	}

	record, _ := entity.NewRecord(recordEnt.ID, foods, recordEnt.LastUpdatedAt.Unix(), recordEnt.CreatedAt.Unix())
	return record, nil
}

func (u ListRecord) List(logger echo.Logger) ([]entity.Record, error) {
	rq := u.dbClient.Record.Query().Limit(50)
	recordsEnt, err := rq.All(context.Background())
	if err != nil {
		logger.Error("All: ", err)
		return nil, fmt.Errorf("All: %w", err)
	}
	records := make([]entity.Record, 0, len(recordsEnt))
	ids := make([]int, 0, len(recordsEnt))
	for _, r := range recordsEnt {
		ids = append(ids, r.ID)
	}
	recordFoodsEnt, err := u.dbClient.RecordFood.Query().
		Where(func(s *sql.Selector) { sql.InInts(recordfood.FieldRecordID, ids...) }).
		All(context.Background())
	if err != nil {
		logger.Error("All: ", err)
		return nil, fmt.Errorf("All: %w", err)
	}
	foodsEnt, err := u.dbClient.Food.Query().All(context.Background())
	if err != nil {
		logger.Error("All: ", err)
		return nil, fmt.Errorf("All: %w", err)
	}

	for _, r := range recordsEnt {
		foods := make([]entity.Food, 0, len(recordFoodsEnt))
		for _, rf := range recordFoodsEnt {
			if r.ID == rf.RecordID {
				for _, f := range foodsEnt {
					if rf.FoodID == f.ID {
						food, _ := entity.NewFood(f.ID, f.Name, rf.Amount, f.Unit)
						foods = append(foods, food)
						break
					}
				}
			}
		}
		record, _ := entity.NewRecord(r.ID, foods, r.LastUpdatedAt.Unix(), r.CreatedAt.Unix())
		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool { return records[i].CreatedAt() < records[j].CreatedAt() })
	return records, nil
}
