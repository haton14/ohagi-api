package repository

import (
	"context"
	"fmt"
	"sort"

	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/domain/value"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/food"
)

type FoodIF interface {
	Save(food *entity.Food) error
	SaveV2(food value.Food) (*entity.Foodv3, error)
	FindByNameUnit(name, unit string) (*entity.Foodv2, error)
	FindByNameUnitV2(name value.FoodName, unit value.FoodUnit) ([]entity.Foodv3, error)
	List() ([]entity.Foodv2, error)
	ListV2() ([]entity.Foodv3, error)
	FindByID(id int) (*entity.Food, error)
	FindByIDV2(id value.ID) (*entity.Foodv3, error)
	UpdateNameUnitFindByID(name, unit string, id int) (*entity.Food, error)
	UpdateNameUnitFindByIDV2(name value.FoodName, unit value.FoodUnit, id value.ID) (*entity.Foodv3, error)
}

type Food struct {
	dbClient *ent.Client
}

func NewFood(dbClinet *ent.Client) FoodIF {
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
func (r Food) SaveV2(food value.Food) (*entity.Foodv3, error) {
	data, err := r.dbClient.Food.Create().SetName(food.Name()).SetUnit(food.Unit()).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foods検索時, detail:%s", ErrOthers, err)
	}
	id, err := value.NewID(data.ID)
	if err != nil {
		return nil, fmt.Errorf("[%w]ID生成時, detail:%s", ErrDomainGenerate, err)
	}
	return entity.NewFoodv3(*id, food), nil
}

func (r Food) FindByNameUnit(name, unit string) (*entity.Foodv2, error) {
	db, err := r.dbClient.Food.Query().Where(food.Name(name), food.Unit(unit)).All(context.Background())
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, nil
	}
	food, _ := entity.NewFoodv2(db[0].ID, db[0].Name, db[0].Unit)
	return food, nil
}

func (r Food) FindByNameUnitV2(name value.FoodName, unit value.FoodUnit) ([]entity.Foodv3, error) {
	datas, err := r.dbClient.Food.Query().Where(food.Name(name.Value()), food.Unit(unit.Value())).All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foods検索時, detail:%s", ErrOthers, err)
	} else if len(datas) < 1 {
		return nil, fmt.Errorf("[%w]foodsに該当レコードなし", ErrNotFoundRecord)
	}

	foods := make([]entity.Foodv3, 0, len(datas))
	for _, data := range datas {
		id, err := value.NewID(data.ID)
		if err != nil {
			return nil, fmt.Errorf("[%w]ID生成時, detail:%s", ErrDomainGenerate, err)
		}
		food := entity.NewFoodv3(*id, *value.NewFood(name, unit))
		foods = append(foods, *food)
	}
	return foods, nil
}

func (r Food) List() ([]entity.Foodv2, error) {
	db, err := r.dbClient.Food.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foodsの検索時", ErrOthers)
	}
	if db == nil {
		return nil, fmt.Errorf("[%w]foodsの検索結果0件", ErrNotFoundRecord)
	}
	foods := make([]entity.Foodv2, 0, len(db))
	for _, f := range db {
		food, err := entity.NewFoodv2(f.ID, f.Name, f.Unit)
		if err != nil {
			return nil, fmt.Errorf("[%w]Foodv2生成時", ErrDomainGenerate)
		}
		foods = append(foods, *food)
	}
	sort.SliceStable(foods, func(i, j int) bool { return foods[i].Unit() < foods[j].Unit() })
	sort.SliceStable(foods, func(i, j int) bool { return foods[i].Name() < foods[j].Name() })
	return foods, nil
}

func (r Food) ListV2() ([]entity.Foodv3, error) {
	datas, err := r.dbClient.Food.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foods検索時, detail:%s", ErrOthers, err)
	} else if len(datas) < 1 {
		return nil, fmt.Errorf("[%w]foodsに該当レコードなし", ErrNotFoundRecord)
	}
	foods := make([]entity.Foodv3, 0, len(datas))
	for _, data := range datas {
		id, err := value.NewID(data.ID)
		if err != nil {
			return nil, fmt.Errorf("[%w]ID生成時, detail:%s", ErrDomainGenerate, err)
		}
		name, err := value.NewFoodName(data.Name)
		if err != nil {
			return nil, fmt.Errorf("[%w]FoodName生成時, detail:%s", ErrDomainGenerate, err)
		}
		unit, err := value.NewFoodUnit(data.Unit)
		if err != nil {
			return nil, fmt.Errorf("[%w]FoodUnit生成時, detail:%s", ErrDomainGenerate, err)
		}
		food := entity.NewFoodv3(*id, *value.NewFood(*name, *unit))
		foods = append(foods, *food)
	}
	sort.SliceStable(foods, func(i, j int) bool { return foods[i].Unit() < foods[j].Unit() })
	sort.SliceStable(foods, func(i, j int) bool { return foods[i].Name() < foods[j].Name() })
	return foods, nil
}
func (r Food) FindByID(id int) (*entity.Food, error) {
	db, err := r.dbClient.Food.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	food, _ := entity.NewFood(db.ID, db.Name, 0, db.Unit)
	return &food, nil
}
func (r Food) FindByIDV2(id value.ID) (*entity.Foodv3, error) {
	data, err := r.dbClient.Food.Get(context.Background(), id.Value())
	if err != nil {
		return nil, fmt.Errorf("[%w]foodsに該当レコードなし", ErrNotFoundRecord)
	}
	name, err := value.NewFoodName(data.Name)
	if err != nil {
		return nil, fmt.Errorf("[%w]FoodName生成時, detail:%s", ErrDomainGenerate, err)
	}
	unit, err := value.NewFoodUnit(data.Unit)
	if err != nil {
		return nil, fmt.Errorf("[%w]FoodUnit生成時, detail:%s", ErrDomainGenerate, err)
	}
	return entity.NewFoodv3(id, *value.NewFood(*name, *unit)), nil
}

func (r Food) UpdateNameUnitFindByID(name, unit string, id int) (*entity.Food, error) {
	db, err := r.dbClient.Food.UpdateOneID(id).SetName(name).SetUnit(unit).Save(context.Background())
	if err != nil {
		return nil, err
	}
	food, _ := entity.NewFood(db.ID, db.Name, 0, db.Unit)
	return &food, nil
}

func (r Food) UpdateNameUnitFindByIDV2(name value.FoodName, unit value.FoodUnit, id value.ID) (*entity.Foodv3, error) {
	_, err := r.dbClient.Food.UpdateOneID(id.Value()).SetName(name.Value()).SetUnit(unit.Value()).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foods更新時, detail:%s", ErrOthers, err)
	}
	return entity.NewFoodv3(id, *value.NewFood(name, unit)), nil
}
