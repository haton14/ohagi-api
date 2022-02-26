package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/haton14/ohagi-api/controller/response"
	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/repository"
	"github.com/labstack/echo/v4"
)

type CreateRecordIF interface {
	Create(request schema.RecordRequestIF, logger echo.Logger) (entity.Record, error)
}

type ListRecordIF interface {
	List(logger echo.Logger) ([]entity.Recordv2, *response.ErrorResponse)
}

type Record struct {
	CreateRecordIF
	ListRecordIF
}

type CreateRecord struct {
	dbClient *ent.Client
}

type ListRecord struct {
	recordRepo repository.RecordIF
}

func NewRecord(dbClient *ent.Client, recordRepo repository.RecordIF) Record {
	return Record{
		CreateRecord{dbClient: dbClient},
		ListRecord{recordRepo: recordRepo},
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

func (u ListRecord) List(logger echo.Logger) ([]entity.Recordv2, *response.ErrorResponse) {
	records, err := u.recordRepo.List()
	if errors.Is(err, repository.ErrNotFoundRecord) {
		logger.Warn("%w;recordRepo.List()でエラー", err)
		return nil, &response.ErrorResponse{Message: "データが存在しない", HttpStatus: http.StatusNotFound}
	} else if err != nil {
		logger.Error("%w;recordRepo.List()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー"}
	}
	return records, nil
}
