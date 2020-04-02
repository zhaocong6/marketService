package services

import (
	"errors"
	"github.com/zhaocong6/market"
	"time"
	"ws/marketApi/app/api/request"
	"ws/marketApi/models"
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
