package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Market struct {
	Model

	Organize string     `json:"organize" gorm:"unique_index:organize_type_symbol_uindex"`
	Symbol   string     `json:"organize" gorm:"unique_index:organize_type_symbol_uindex"`
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
