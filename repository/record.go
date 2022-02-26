package repository

import (
	"context"
	"fmt"
	"sort"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/recordfood"
)

type RecordIF interface {
	List() ([]entity.Recordv2, error)
}

type Record struct {
	dbClient *ent.Client
}

func NewRecord(dbClinet *ent.Client) RecordIF {
	return Record{dbClient: dbClinet}
}

func (r Record) List() ([]entity.Recordv2, error) {
	rq := r.dbClient.Record.Query().Limit(50)
	recordsDB, err := rq.All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]recordの検索時", ErrOthers)
	}
	records := make([]entity.Recordv2, 0, len(recordsDB))
	ids := make([]int, 0, len(recordsDB))
	for _, r := range recordsDB {
		ids = append(ids, r.ID)
	}
	recordFoodsDB, err := r.dbClient.RecordFood.Query().
		Where(func(s *sql.Selector) { sql.InInts(recordfood.FieldRecordID, ids...) }).
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]record_foodsの検索時", ErrOthers)
	}
	foodsDB, err := r.dbClient.Food.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foodsの検索時", ErrOthers)
	}

	for _, r := range recordsDB {
		foods := make([]entity.RecordFood, 0, len(recordFoodsDB))
		for _, rf := range recordFoodsDB {
			if r.ID == rf.RecordID {
				for _, f := range foodsDB {
					if rf.FoodID == f.ID {
						food, err := entity.NewRecordFood(f.ID, f.Name, f.Unit, rf.Amount)
						if err != nil {
							return nil, fmt.Errorf("[%w]RecordFood生成時", ErrDomainGenerate)
						}
						foods = append(foods, *food)
						break
					}
				}
			}
		}
		record, _ := entity.NewRecordv2(r.ID, foods, r.LastUpdatedAt.Unix(), r.CreatedAt.Unix())
		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool { return records[i].CreatedAt() < records[j].CreatedAt() })
	return records, nil
}
