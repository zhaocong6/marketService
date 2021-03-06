package services

import (
	"errors"
	"github.com/zhaocong6/market"
	"marketApi/app/api/request"
	"marketApi/models"
	"time"
)

type MarketService struct{}

func (m *MarketService) AddAndSub(req *request.MarketRequest) error {

	marketModel := &models.Market{
		Organize: req.Organize,
		Symbol:   req.Symbol,
		Type:     req.MarketType,
		Expire:   nil,
		Status:   1,
	}

	transaction := models.NewTransaction()
	defer transaction.Rollback()

	if err := models.AddMarket(marketModel, transaction.Tx); err != nil {
		return err
	}

	h := &market.Subscriber{
		Symbol:     req.Symbol,
		MarketType: market.MarketType(req.MarketType),
		Organize:   market.Organize(req.Organize),
	}

	select {
	case <-time.After(time.Second * 10):
		return errors.New("sub timeout")
	case market.WriteSubscribing <- h:
	}

	transaction.Commit()

	return nil
}

type GetMarketParams struct {
	Keys   map[string][]string
	Fields []string
}

func (m *MarketService) GetMarketData(p *GetMarketParams) map[string]*market.Marketer {
	for organize, symbols := range p.Keys {
		data := market.Find(organize, symbols...)
		fields := m.keysMap(p.Fields)
		if len(fields) > 0 {
			data = m.marketFieldFilter(data, fields)
		}

		return data
	}
	return nil
}

//过滤字段
//此处可以修改为位图运算
func (m *MarketService) marketFieldFilter(data map[string]*market.Marketer, fieldsMap map[string]interface{}) map[string]*market.Marketer {
	newMarketData := make(map[string]*market.Marketer)

	for key, val := range data {

		m := &market.Marketer{
			Organize:  "",
			Symbol:    "",
			BuyFirst:  "",
			SellFirst: "",
			BuyDepth:  nil,
			SellDepth: nil,
			Timestamp: 0,
			Temporize: 0,
		}

		m.Organize = val.Organize
		m.Symbol = val.Symbol

		if _, ok := fieldsMap["buy_first"]; ok {
			m.BuyFirst = val.BuyFirst
		}

		if _, ok := fieldsMap["sell_first"]; ok {
			m.SellFirst = val.SellFirst
		}

		if _, ok := fieldsMap["buy_depth"]; ok {
			m.BuyDepth = val.BuyDepth
		}

		if _, ok := fieldsMap["sell_depth"]; ok {
			m.SellDepth = val.SellDepth
		}

		if _, ok := fieldsMap["timestamp"]; ok {
			m.Timestamp = val.Timestamp
		}

		if _, ok := fieldsMap["temporize"]; ok {
			m.Temporize = val.Temporize
		}

		newMarketData[key] = m
	}

	return newMarketData
}

func (m *MarketService) keysMap(keys []string) map[string]interface{} {
	keyMap := make(map[string]interface{})

	for _, v := range keys {
		keyMap[v] = nil
	}

	return keyMap
}
