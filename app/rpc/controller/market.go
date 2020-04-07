package controller

import (
	"context"
	"marketApi/app/rpc/services"
	pd "marketApi/pd/market"
)

type Market Base

var marketService = &services.MarketService{}

func (m *Market) GetMarket(ctx context.Context, in *pd.MarketRequest) (*pd.MarketResponse, error) {
	resp := marketService.GetMarketData(in)
	return resp, nil
}
