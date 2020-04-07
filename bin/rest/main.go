package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhaocong6/market"
	"log"
	"marketApi/models"
	"marketApi/pkg/setting"
	"marketApi/routes"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer models.CloseDB()

	//启动行情监听服务
	if setting.WsProxy.Port != 0 {
		uProxy, _ := url.Parse(fmt.Sprintf("http://%s:%d", setting.WsProxy.Host, setting.WsProxy.Port))
		market.DefaultDialer = &websocket.Dialer{
			Proxy:            http.ProxyURL(uProxy),
			HandshakeTimeout: 10 * time.Second,
		}
	}
	market.Run()
	defer market.Close()

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

	//启动gin
	gin.SetMode(setting.RunMode)
	router := gin.New()
	routes.InitApi(router)
	server := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf(":%d", setting.HTTP.Port),
		IdleTimeout:       20 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	//平滑重启
	ch := make(chan os.Signal)
	//监听信号
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)
	sig := <-ch
	log.Println("exit signal:", sig)

	//设置一个关闭最大超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	//向服务发送关闭信号
	go server.Shutdown(ctx)

	<-ctx.Done()
	log.Println("shutting down")
}
