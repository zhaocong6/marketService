package controller

import (
	"context"
	"fmt"
	"github.com/zhaocong6/goUtils/goroutinepool"
	"google.golang.org/grpc"
	"sync"
	"testing"
	pd "marketApi/pd/market"
)

func TestGet(t *testing.T) {
	var wg sync.WaitGroup
	num := 1
	wg.Add(num)

	conn, _ := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure())
	w := goroutinepool.NewPool(goroutinepool.Options{
		Capacity:  3000,
		JobBuffer: 1000,
	})

	j := &Job{
		c:  pd.NewMarketClient(conn),
		wg: &wg,
	}
	for i := 0; i < num; i++ {
		w.Put(j)
	}
	wg.Wait()

	fmt.Println(j.data)
}

type Job struct {
	c    pd.MarketClient
	wg   *sync.WaitGroup
	data *pd.MarketResponse
}

func (j *Job) Handle() error {
	defer j.wg.Done()

	keys := make([]string, 2)
	keys[0] = "buy_first"
	keys[1] = "sell_first"

	req := &pd.MarketRequest{
		Organize: "okex",
		Symbol:   "EOS-USDT",
		Keys:     keys,
	}

	j.data, _ = j.c.GetMarket(context.Background(), req)

	return nil
}
