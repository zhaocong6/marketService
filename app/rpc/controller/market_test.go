package controller

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	marketPd "market.pd"
	"testing"
)

func TestGet(t *testing.T) {
	conn, _ := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure())
	c := marketPd.NewMarketClient(conn)
	req := &marketPd.MarketRequest{
		Organize: "huobi",
		Symbol:   "btcusdt",
	}

	res, _ := c.GetMarket(context.Background(), req)

	fmt.Println(res)
}
