package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zhaocong6/market"
	"google.golang.org/grpc"
	"log"
	marketPd "market.pd"
	"net"
	"net/http"
	"net/url"
	"time"
	rpcMarket "ws/marketApi/app/rpc/controller/market"
	"ws/marketApi/pkg/setting"
)

func main() {
	marketRun()
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", setting.RPC.Port))
	if err != err {
		log.Panicln(err)
	}

	s := grpc.NewServer()
	marketPd.RegisterMarketServer(s, &rpcMarket.Market{})

	log.Println("服务已启动...")
	err = s.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}

func marketRun() {
	uProxy, _ := url.Parse("http://127.0.0.1:8888")
	market.DefaultDialer = &websocket.Dialer{
		Proxy:            http.ProxyURL(uProxy),
		HandshakeTimeout: 10 * time.Second,
	}
	market.Run()

	for k := range map[string]string{
		"ETH-USDT": "",
		"BTC-USDT": "",
		"EOS-USDT": "",
		"TRX-USDT": "",
		"BCH-USDT": "",
		"OKB-USDT": "",
		"XRP-USDT": "",
		"BSV-USDT": "",
		"LTC-USDT": "",
		"ADA-USDT": "",
	} {

		s := &market.Subscriber{
			Symbol:     k,
			MarketType: market.SpotMarket,
			Organize:   market.OkEx,
		}
		market.WriteSubscribing <- s
	}

	for k := range map[string]string{
		"ethusdt": "",
		"btcusdt": "",
		"eosusdt": "",
		"trxusdt": "",
		"bchusdt": "",
		"htusdt":  "",
		"xrpusdt": "",
		"bsvusdt": "",
		"ltcusdt": "",
		"adausdt": "",
	} {

		h := &market.Subscriber{
			Symbol:     k,
			MarketType: market.SpotMarket,
			Organize:   market.HuoBi,
		}

		market.WriteSubscribing <- h
	}
}
