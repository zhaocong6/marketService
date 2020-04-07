package request

import (
	"gopkg.in/go-playground/validator.v9"
	"marketApi/models"
)

type MarketRequest struct {
	requester

	Organize   string `json:"organize" validate:"required,organizeMarketTypeSymbolUnique"`
	MarketType int8   `json:"market_type" validate:"required,min=1,max=4"`
	Symbol     string `json:"symbol" validate:"required"`
}

//唯一验证
//没有保障并发安全
func organizeMarketTypeSymbolUnique(fl validator.FieldLevel) bool {
	var req MarketRequest
	val, _, _, _ := fl.GetStructFieldOK2()
	req = val.Interface().(MarketRequest)

	data := &models.Market{
		Organize: req.Organize,
		Symbol:   req.Symbol,
		Type:     req.MarketType,
	}

	data.FirstByQuery(data.Query())

	if data.ID != 0 {
		return false
	}

	return true
}
