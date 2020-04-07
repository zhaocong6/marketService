package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zhaocong6/market"
	"google.golang.org/grpc"
	"log"
	"marketApi/models"
	"marketApi/pkg/setting"
	"marketApi/routes"
	"net"
	"net/http"
	"net/url"
	"time"
)

func main() {
	defer models.CloseDB()
	marketListen()
	serveListen()
}

func serveListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", setting.RPC.Port))
	if err != err {
		log.Panicln(err)
	}

	s := grpc.NewServer()
	routes.InitRpc(s)

	log.Println("服务已启动...")
	err = s.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}

func marketListen() {
	if setting.WsProxy.Port != 0 {
		uProxy, _ := url.Parse(fmt.Sprintf("http://%s:%d", setting.WsProxy.Host, setting.WsProxy.Port))
		market.DefaultDialer = &websocket.Dialer{
			Proxy:            http.ProxyURL(uProxy),
			HandshakeTimeout: 10 * time.Second,
		}
	}

	market.Run()

	go func() {
		mar := &models.Market{}

		mar.GetChunk(mar.Query(), func(markets []models.Market) {
			for _, m := range markets {
				h := &market.Subscriber{
					Symbol:     m.Symbol,
					MarketType: market.MarketType(m.Type),
					Organize:   market.Organize(m.Organize),
				}
				market.WriteSubscribing <- h
			}
		})
	}()
}
