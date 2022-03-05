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
	Create([]entity.FoodContent, []value.FoodContent, echo.Logger) (*entity.Recordv3, *response.ErrorResponse)
}

type ListRecordIF interface {
	List(logger echo.Logger) ([]entity.Recordv3, *response.ErrorResponse)
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
) (*entity.Recordv3, *response.ErrorResponse) {
	now := time.Now()
	record, err := u.recordRepo.Save(now.Unix(), now.Unix())
	if err != nil {
		logger.Error("%w;foodRepo.SaveV2()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}

	for _, f := range foodContentsValue {
		conflict, err := u.foodRepo.FindByNameUnitV2(f.Food())
		if len(conflict) > 0 {
			return nil, &response.ErrorResponse{Message: "登録しようとした食事は既に存在", HttpStatus: http.StatusConflict}
		}
		food, err := u.foodRepo.SaveV2(f.Food())
		if err != nil {
			logger.Error("%w;foodRepo.SaveV2()でエラー", err)
			return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
		}
		foodContentsEntity = append(foodContentsEntity, *entity.NewFoodContent(*food, f.Amont()))
	}
	err = u.recordRepo.SaveFoodContent(*record)
	if err != nil {
		logger.Error("%w;recordRepo.SaveFoodContent()でエラー", err)
		return nil, &response.ErrorResponse{Message: "予期しないエラー", HttpStatus: http.StatusInternalServerError}
	}
	return record, nil
}

func (u ListRecord) List(logger echo.Logger) ([]entity.Recordv3, *response.ErrorResponse) {
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
