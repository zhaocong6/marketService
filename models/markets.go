package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Market struct {
	Model

	Organize string     `json:"organize" gorm:"unique_index:organize_type_symbol_uindex"`
	Symbol   string     `json:"symbol" gorm:"unique_index:organize_type_symbol_uindex"`
	Type     int8       `json:"type" gorm:"unique_index:organize_type_symbol_uindex"`
	Expire   *time.Time `json:"expire"`
	Status   int8       `json:"status"`
}

func (Market) TableName() string {
	return "market"
}

func AddMarket(m *Market, tx *gorm.DB) error {
	if tx != nil {
		return tx.Create(m).Error
	}

	return db.Create(m).Error
}

//创建一个query map
func (m *Market) Query() *Query {
	query := &Query{
		Table: m.TableName(),
		Where: make(queryWhere),
	}

	if m.Organize != "" {
		query.Where["organize"] = m.Organize
	}

	if m.Symbol != "" {
		query.Where["symbol"] = m.Symbol
	}

	if m.Type != 0 {
		query.Where["type"] = m.Type
	}

	return query
}

func (m *Market) FirstByQuery(q *Query) {
	db.Table(q.Table).Where(q.Where).First(m)
}

//使用雪花ID迭代数据
func (*Market) GetChunk(query *Query, callback func(markets []Market)) {
	var (
		markets []Market
	)

	db = db.Where(query.Where).Limit(100).Order("id asc")
	db.Find(&markets)
	callback(markets)

	for len(markets) > 0 {
		lastId := markets[len(markets)-1].ID
		db := db.Where("id > ?", lastId)
		db.Find(&markets)

		if len(markets) > 0 {
			callback(markets)
		}
	}
}
