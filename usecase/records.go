package usecase

import (
	"errors"
	"net/http"
	"time"

	"github.com/haton14/ohagi-api/controller/response"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/domain/value"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/repository"
	"github.com/labstack/echo/v4"
)

type CreateRecordIF interface {
	Create([]entity.FoodContent, []value.FoodContent, echo.Logger) (*entity.Record, *response.ErrorResponse)
}

type ListRecordIF interface {
	List(logger echo.Logger) ([]entity.Record, *response.ErrorResponse)
}

type Record struct {
	CreateRecordIF
	ListRecordIF
}

type CreateRecord struct {
	dbClient   *ent.Client
	recordRepo repository.RecordIF
	foodRepo   repository.FoodIF
}

type ListRecord struct {
	recordRepo repository.RecordIF
}

func NewRecord(dbClient *ent.Client, recordRepo repository.RecordIF, foodRepo repository.FoodIF) Record {
	return Record{
		CreateRecord{dbClient: dbClient, recordRepo: recordRepo, foodRepo: foodRepo},
		ListRecord{recordRepo: recordRepo},
	}
}

func (u CreateRecord) Create(
	foodContentsEntity []entity.FoodContent,
	foodContentsValue []value.FoodContent,
	logger echo.Logger,
) (*entity.Record, *response.ErrorResponse) {
	now := time.Now()
	record, err := u.recordRepo.Save(now.Unix(), now.Unix())
	if err != nil {
		logger.Error("%w;foodRepo.Save()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}

	for _, f := range foodContentsValue {
		conflict, err := u.foodRepo.FindByNameUnit(f.Food())
		if len(conflict) > 0 {
			return nil, &response.ErrorResponse{Message: "登録しようとした食事は既に存在", HttpStatus: http.StatusConflict}
		}
		food, err := u.foodRepo.Save(f.Food())
		if err != nil {
			logger.Error("%w;foodRepo.Save()でエラー", err)
			return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
		}
		record.AddFoodContent(*food, f.Amont())
	}
	for _, f := range foodContentsEntity {
		_, err := u.foodRepo.FindByID(f.ID())
		if errors.Is(err, repository.ErrNotFoundRecord) {
			logger.Warn("%w;foodRepo.FindByID()でエラー", err)
			return nil, &response.ErrorResponse{Message: "データが存在しない", HttpStatus: http.StatusNotFound}
		} else if err != nil {
			logger.Error("%w;foodRepo.FindByID()でエラー", err)
			return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
		}
		record.AddFoodContent(*entity.NewFoodv3(f.ID(), f.Food()), f.Amont())
	}

	err = u.recordRepo.SaveFoodContent(*record)
	if err != nil {
		logger.Error("%w;recordRepo.SaveFoodContent()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}
	return record, nil
}

func (u ListRecord) List(logger echo.Logger) ([]entity.Record, *response.ErrorResponse) {
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
