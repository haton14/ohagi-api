package repository

import (
	"context"

	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/food"
)

type FoodIF interface {
	Save(food *entity.Food) error
	FindByNameUnit(name, unit string) (*entity.Food, error)
}

type Food struct {
	dbClient *ent.Client
}

func NewFood(dbClinet *ent.Client) Food {
	return Food{dbClient: dbClinet}
}

func (r Food) Save(food *entity.Food) error {
	db, err := r.dbClient.Food.Create().SetName(food.Name()).SetUnit(food.Unit()).Save(context.Background())
	if err != nil {
		return err
	}
	food.SetID(db.ID)
	return nil
}

func (r Food) FindByNameUnit(name, unit string) (*entity.Food, error) {
	db, err := r.dbClient.Food.Query().Where(food.Name(name), food.Unit(unit)).All(context.Background())
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, nil
	}
	food, _ := entity.NewFood(db[0].ID, db[0].Name, 0, db[0].Unit)
	return &food, nil
}
