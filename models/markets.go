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
func (m Market) Query() map[string]interface{} {
	var query = make(map[string]interface{})

	if m.Organize != "" {
		query["organize"] = m.Organize
	}

	if m.Symbol != "" {
		query["symbol"] = m.Symbol
	}

	if m.Type != 0 {
		query["type"] = m.Type
	}

	return query
}

//使用雪花ID迭代数据
func (Market) GetChunk(query map[string]interface{}, callback func(markets []Market)) {
	var (
		markets []Market
	)

	db.Where(query).Where(query).Limit(100).Find(&markets)
	callback(markets)

	for len(markets) > 0 {
		lastId := markets[len(markets)-1].ID
		db := *db.Where("id > ?", lastId)
		db.Where(query).Limit(1).Find(&markets)

		if len(markets) > 0 {
			callback(markets)
		}
	}
}
