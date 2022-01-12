package schema

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/labstack/echo/v4"
)

type record struct {
	ID            int    `json:"id,omitempty"`
	Foods         []food `json:"foods"`
	LastUpdatedAt int64  `json:"last_updated_at,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
}

type RecordResponse struct {
	c echo.Context
	record
}

type RecordRequestIF interface {
	GetFoods() []food
}
type RecordResponseIF interface {
	JSON(code int) error
	RecordRequestIF
}

func NewRecordRequest(c echo.Context) (RecordRequestIF, error) {
	s := record{}
	if err := c.Bind(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func NewRecordResponse(c echo.Context, r entity.Record) RecordResponseIF {
	foods := make([]food, 0, len(r.Foods()))
	for _, f := range r.Foods() {
		id := f.ID()
		food := food{ID: &id, Name: f.Name(), Amount: f.Amount(), Unit: f.Unit()}
		foods = append(foods, food)
	}
	response := record{ID: r.ID(), Foods: foods, LastUpdatedAt: r.LastUpdatedAt(), CreatedAt: r.CreatedAt()}
	return RecordResponse{c, response}
}

func (s record) GetFoods() []food {
	return s.Foods
}

func (s RecordResponse) JSON(code int) error {
	return s.c.JSON(code, s.record)
}
