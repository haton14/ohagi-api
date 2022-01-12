package schema

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/labstack/echo/v4"
)

type records struct {
	Records []record `json:"records"`
}

type RecordsResponseIF interface {
	JSON(code int) error
}
type RecordsResponse struct {
	c       echo.Context
	records []record
}

func NewRecordsResponse(c echo.Context, records []entity.Record) RecordsResponseIF {
	response := make([]record, 0, len(records))
	for _, r := range records {
		foods := make([]food, 0, len(r.Foods()))
		for _, f := range r.Foods() {
			id := f.ID()
			food := food{ID: &id, Name: f.Name(), Amount: f.Amount(), Unit: f.Unit()}
			foods = append(foods, food)
		}
		record := record{ID: r.ID(), Foods: foods, LastUpdatedAt: r.LastUpdatedAt(), CreatedAt: r.CreatedAt()}
		response = append(response, record)
	}
	return RecordsResponse{c, response}
}

func (s RecordsResponse) JSON(code int) error {
	return s.c.JSON(code, s.records)
}
