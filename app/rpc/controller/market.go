package controller

import (
	"context"
	"ws/marketApi/app/rpc/services"
	pd "ws/marketApi/pd/market"
)

type Market Base

var marketService = &services.MarketService{}

func (m *Market) GetMarket(ctx context.Context, in *pd.MarketRequest) (*pd.MarketResponse, error) {
	resp := marketService.GetMarketData(in)
	return resp, nil
}
