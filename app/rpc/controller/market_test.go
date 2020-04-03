package controller

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
	pd "ws/marketApi/pd/market"
)

func TestGet(t *testing.T) {
	conn, _ := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure())
	c := pd.NewMarketClient(conn)

	keys := make([]string, 2)
	keys[0] = "buy_first"
	keys[1] = "sell_first"

	req := &pd.MarketRequest{
		Organize: "huobi",
		Symbol:   "btcusdt",
		Keys:     keys,
	}

	_, err := c.GetMarket(context.Background(), req)

	s, _ := status.FromError(err)
	fmt.Println(s.Code(), s.Message(), s.Err())
}
