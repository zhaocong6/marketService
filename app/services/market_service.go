package services

import (
	"errors"
	"github.com/zhaocong6/market"
	"time"
	"ws/marketApi/models"
)

type MarketerStore struct {
	Organize   string `json:"organize" binding:"required"`
	MarketType int8   `json:"market_type" binding:"required,min=1,max=4"`
	Symbol     string `json:"symbol" binding:"required"`
}

type MarketService struct{}

func (m *MarketService) AddAndSub(store *MarketerStore) error {

	marketModel := &models.Market{
		Organize: store.Organize,
		Symbol:   store.Symbol,
		Type:     store.MarketType,
		Expire:   nil,
		Status:   1,
	}

	transaction := models.NewTransaction()
	defer transaction.Rollback()

	if err := models.AddMarket(marketModel, transaction.Tx); err != nil {
		return err
	}

	h := &market.Subscriber{
		Symbol:     store.Symbol,
		MarketType: market.MarketType(store.MarketType),
		Organize:   market.Organize(store.Organize),
	}

	select {
	case <-time.After(time.Second * 10):
		return errors.New("sub timeout")
	case market.WriteSubscribing <- h:
	}

	transaction.Commit()

	return nil
}
