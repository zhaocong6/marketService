package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaocong6/market"
	"ws/marketApi/app/api/request"
	"ws/marketApi/app/services"
)

type Marketer struct {
	Base
}

func (m *Marketer) Index(c *gin.Context) {
	data := market.Find(c.Query("organize"), c.Query("symbol"))
	m.SuccessResponse(c, data)
}

func (m *Marketer) Store(c *gin.Context) {
	var (
		req     = request.MarketRequest{}
		service = &services.MarketService{}
	)

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
