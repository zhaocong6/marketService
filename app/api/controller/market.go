package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaocong6/market"
	"marketApi/app/api/request"
	"marketApi/app/api/services"
)

type Marketer struct {
	Base
}

var service = &services.MarketService{}

func (m *Marketer) Index(c *gin.Context) {
	symbol := c.Param("symbol")
	organize := c.Param("organize")
	fields := c.QueryArray("fields[]")

	keys := make(map[string][]string)
	symbols := make([]string, 1)
	symbols[0] = symbol
	keys[organize] = symbols

	data := service.GetMarketData(&services.GetMarketParams{
		Keys:   keys,
		Fields: fields,
	})
	m.SuccessResponse(c, data[symbol])
}

func (m *Marketer) List(c *gin.Context) {
	fields := c.QueryArray("fields[]")
	keysMap := c.QueryMap("keys")
	keys := make(map[string][]string)
	data := make(map[string]map[string]*market.Marketer)

	for k := range keysMap {
		keys[k] = c.QueryArray(fmt.Sprintf("keys[%s][]", k))
		params := &services.GetMarketParams{
			Keys:   keys,
			Fields: fields,
		}

		if ok := service.GetMarketData(params); len(ok) != 0 {
			data[k] = ok
		}
	}

	m.SuccessResponse(c, data)
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
