package logic

import (
	"errors"
	"github.com/zhaocong6/market"
	"time"
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
	defer transaction.Rollback()

	if err := models.AddMarket(marketModel, transaction.Tx); err != nil {
		return err
	}

	h := &market.Subscriber{
		Symbol:     s.Symbol,
		MarketType: market.MarketType(s.MarketType),
		Organize:   market.Organize(s.Organize),
	}
	select {
	case <-time.After(time.Second * 10):
		return errors.New("币对增加失败")
	case market.WriteSubscribing <- h:
	}

	transaction.Commit()
	return nil
}
