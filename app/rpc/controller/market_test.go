package controller

import (
	"context"
	"encoding/json"
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
		Organize: "okex",
		Symbol:   "EOS-USDT",
		Keys:     keys,
	}

	v, err := c.GetMarket(context.Background(), req)
	s, _ := status.FromError(err)

	j, _ := json.Marshal(v)

	fmt.Println(s.Code(), s.Message(), s.Err())
	fmt.Println(string(j))
}
