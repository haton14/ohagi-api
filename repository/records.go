package repository

import (
	"context"
	"fmt"
	"sort"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/domain/value"
	"github.com/haton14/ohagi-api/ent"
	"github.com/haton14/ohagi-api/ent/record"
	"github.com/haton14/ohagi-api/ent/recordfood"
)

type RecordIF interface {
	List() ([]entity.Record, error)
	Save(lastUpdatedAt, createdAt int64) (*entity.Record, error)
	SaveFoodContent(record entity.Record) error
}

type Record struct {
	dbClient *ent.Client
}

func NewRecord(dbClinet *ent.Client) RecordIF {
	return Record{dbClient: dbClinet}
}

func (r Record) List() ([]entity.Record, error) {
	rq := r.dbClient.Record.Query().Order(ent.Asc(record.FieldID)).Limit(50)
	recordDatas, err := rq.All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]recordの検索時", ErrOthers)
	}
	records := make([]entity.Record, 0, len(recordDatas))
	ids := make([]int, 0, len(recordDatas))
	for _, r := range recordDatas {
		ids = append(ids, r.ID)
	}
	contentDatas, err := r.dbClient.RecordFood.Query().
		Where(func(s *sql.Selector) { sql.InInts(recordfood.FieldRecordID, ids...) }).
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]record_foodsの検索時", ErrOthers)
	}
	foodDatas, err := r.dbClient.Food.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]foodsの検索時", ErrOthers)
	}

	for _, recordData := range recordDatas {
		recordID, err := value.NewID(recordData.ID)
		if err != nil {
			return nil, err
		}
		record := entity.NewRecordv3(*recordID, recordData.LastUpdatedAt.Unix(), recordData.CreatedAt.Unix())
		for _, contentData := range contentDatas {
			if recordData.ID == contentData.RecordID {
				for _, foodData := range foodDatas {
					if contentData.FoodID == foodData.ID {
						foodID, err := value.NewID(recordData.ID)
						if err != nil {
							return nil, err
						}
						name, err := value.NewFoodName(foodData.Name)
						if err != nil {
							return nil, err
						}
						unit, err := value.NewFoodUnit(foodData.Unit)
						if err != nil {
							return nil, err
						}
						amount, err := value.NewFoodAmount(contentData.Amount)
						if err != nil {
							return nil, err
						}
						record.AddFoodContent(*entity.NewFoodv3(*foodID, *value.NewFood(*name, *unit)), *amount)
						break
					}
				}
			}
		}
		records = append(records, *record)
	}

	sort.Slice(records, func(i, j int) bool { return records[i].CreatedAt() < records[j].CreatedAt() })
	return records, nil
}

func (r Record) Save(lastUpdatedAt, createdAt int64) (*entity.Record, error) {
	data, err := r.dbClient.Record.Create().SetCreatedAt(time.Now()).SetLastUpdatedAt(time.Now()).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("[%w]records保存時, detail:%s", ErrOthers, err)
	}
	id, err := value.NewID(data.ID)
	if err != nil {
		return nil, fmt.Errorf("[%w]ID生成時, detail:%s", ErrDomainGenerate, err)
	}
	return entity.NewRecordv3(*id, lastUpdatedAt, createdAt), nil
}

func (r Record) SaveFoodContent(record entity.Record) error {
	bulk := make([]*ent.RecordFoodCreate, 0, record.LenFoodContent())
	for _, foodContent := range record.FoodContents() {
		b := r.dbClient.RecordFood.Create().
			SetRecordID(record.ID().Value()).
			SetFoodID(foodContent.ID().Value()).
			SetAmount(foodContent.Amont().Value())
		bulk = append(bulk, b)
	}
	_, err := r.dbClient.RecordFood.CreateBulk(bulk...).Save(context.Background())
	if err != nil {
		return fmt.Errorf("[%w]record_foods保存時, detail:%s", ErrOthers, err)
	}
	return nil
}
