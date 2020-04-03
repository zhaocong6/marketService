package controller

import (
	"context"
	"fmt"
	"github.com/zhaocong6/market"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pd "ws/marketApi/pd/market"
)

type Market Base

func (m *Market) GetMarket(ctx context.Context, in *pd.MarketRequest) (*pd.MarketResponse, error) {
	resp := &pd.MarketResponse{}

	return nil, status.Error(codes.InvalidArgument, "参数错误")
	fmt.Println(in.Keys)

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
