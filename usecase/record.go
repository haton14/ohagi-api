package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/ent"
	"github.com/labstack/echo/v4"
)

type CreateRecordIF interface {
	Create(request schema.RecordRequestIF, logger echo.Logger) (entity.Record, error)
}

type Record struct {
	CreateRecord
}

type CreateRecord struct {
	dbClient *ent.Client
}

func NewCreateRecord(dbClient *ent.Client) CreateRecordIF {
	return CreateRecord{dbClient: dbClient}
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
