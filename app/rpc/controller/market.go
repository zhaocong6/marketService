package controller

import (
	"context"
	"github.com/zhaocong6/market"
	marketPd "ws/marketApi/proto/market"
)

type Market Base

func (m Market) GetMarket(ctx context.Context, in *marketPd.MarketRequest) (*marketPd.MarketResponse, error) {
	var resp = new(marketPd.MarketResponse)
	data, ok := market.Find(in.Organize, in.Symbol)[in.Symbol]
	if ok {
		resp.Symbol = data.Symbol
		resp.Organize = string(data.Organize)
		resp.Timestamp = int64(data.Timestamp)
		resp.BuyFirst = data.BuyFirst
		resp.SellFirst = data.SellFirst
	}
	return resp, nil
}
