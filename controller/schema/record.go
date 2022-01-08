package schema

import "github.com/labstack/echo/v4"

type Record struct {
	ID            int    `json:"id,omitempty"`
	Foods         []Food `json:"foods"`
	LastUpdatedAt int64  `json:"last_updated_at,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
}

type RecordIF interface {
	GetFoods() []Food
}

func NewRecord(c echo.Context) (RecordIF, error) {
	s := Record{}
	if err := c.Bind(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func (s Record) GetFoods() []Food {
	return s.Foods
}
