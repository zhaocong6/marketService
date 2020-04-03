package controller

import (
	"github.com/gin-gonic/gin"
	"ws/marketApi/app/api/request"
	"ws/marketApi/app/api/services"
)

type Marketer struct {
	Base
}

var service = &services.MarketService{}

func (m *Marketer) Index(c *gin.Context) {
	m.SuccessResponse(c, service.GetMarketData(c))
}

func (m *Marketer) Store(c *gin.Context) {
	var req = request.MarketRequest{}

	if err := req.ValidateZH(c, &req); err != nil {
		m.ValidateResponse(c, err.Error())
		return
	}

	if err := service.AddAndSub(&req); err != nil {
		m.ValidateResponse(c, err.Error())
		return
	}

	m.SuccessResponse(c, nil)
}
