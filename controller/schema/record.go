package schema

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/labstack/echo/v4"
)

type Record struct {
	ID            int    `json:"id,omitempty"`
	Foods         []Food `json:"foods"`
	LastUpdatedAt int64  `json:"last_updated_at,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
}

type RecordResponse struct {
	c echo.Context
	Record
}

type RecordRequestIF interface {
	GetFoods() []Food
}
type RecordResponseIF interface {
	JSON(code int) error
	RecordRequestIF
}

func NewRecordRequest(c echo.Context) (RecordRequestIF, error) {
	s := Record{}
	if err := c.Bind(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func NewRecordResponse(c echo.Context, record entity.Record) RecordResponseIF {
	foods := make([]Food, 0, len(record.Foods()))
	for _, f := range record.Foods() {
		id := f.ID()
		food := Food{ID: &id, Name: f.Name(), Amount: f.Amount(), Unit: f.Unit()}
		foods = append(foods, food)
	}
	response := Record{ID: record.ID(), Foods: foods, LastUpdatedAt: record.LastUpdatedAt(), CreatedAt: record.CreatedAt()}
	return RecordResponse{c, response}
}

func (s Record) GetFoods() []Food {
	return s.Foods
}

func (s RecordResponse) JSON(code int) error {
	return s.c.JSON(code, s.Record)
}
