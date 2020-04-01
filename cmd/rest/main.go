package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zhaocong6/market"
	"log"
	"net/http"
	"net/url"
	"time"
	"ws/marketApi/models"
	"ws/marketApi/pkg/setting"
	"ws/marketApi/routes"
)

func main() {
	defer models.CloseDB()
	marketListen()
	serveListen()
}

func marketListen() {
	uProxy, _ := url.Parse("http://127.0.0.1:8888")
	market.DefaultDialer = &websocket.Dialer{
		Proxy:            http.ProxyURL(uProxy),
		HandshakeTimeout: 10 * time.Second,
	}
	market.Run()
}

func serveListen() {
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

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}
