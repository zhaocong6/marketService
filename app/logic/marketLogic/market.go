package marketLogic

import (
	"github.com/zhaocong6/market"
	"log"
	"ws/marketApi/models"
)

type Store struct {
	Organize   string `json:"organize"`
	MarketType int8   `json:"market_type"`
	Symbol     string `json:"symbol"`
}

func (s *Store) Store() error {
	marketModel := &models.Market{
		Organize: s.Organize,
		Symbol:   s.Symbol,
		Type:     s.MarketType,
		Expire:   nil,
		Status:   1,
	}

	transaction := models.NewTransaction()
	defer transaction.Commit()
	defer func() {
		if err := recover(); err != nil {
			transaction.Rollback()
			log.Println(err)
		}
	}()

	if err := models.AddMarket(marketModel, transaction.Tx); err != nil {
		transaction.Rollback()
		return err
	}

	var marketType = market.SpotMarket
	switch s.MarketType {
	case int8(market.SpotMarket):
		marketType = market.SpotMarket
	case int8(market.FuturesMarket):
		marketType = market.FuturesMarket
	case int8(market.WapMarket):
		marketType = market.WapMarket
	case int8(market.OptionMarket):
		marketType = market.OptionMarket
	}
	h := &market.Subscriber{
		Symbol:     s.Symbol,
		MarketType: marketType,
		Organize:   market.Organize(s.Organize),
	}
	market.WriteSubscribing <- h

	return nil
}
