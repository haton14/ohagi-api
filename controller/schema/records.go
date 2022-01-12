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
	c echo.Context
	records
}

func NewRecordsResponse(c echo.Context, r []entity.Record) RecordsResponseIF {
	response := make([]record, 0, len(r))
	for _, v := range r {
		foods := make([]food, 0, len(v.Foods()))
		for _, f := range v.Foods() {
			id := f.ID()
			food := food{ID: &id, Name: f.Name(), Amount: f.Amount(), Unit: f.Unit()}
			foods = append(foods, food)
		}
		record := record{ID: v.ID(), Foods: foods, LastUpdatedAt: v.LastUpdatedAt(), CreatedAt: v.CreatedAt()}
		response = append(response, record)
	}
	responses := records{response}
	return RecordsResponse{c: c, records: responses}
}

func (s RecordsResponse) JSON(code int) error {
	return s.c.JSON(code, s.records)
}
