package market

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaocong6/market"
	"ws/marketApi/app/api/controller"
	"ws/marketApi/app/logic/marketLogic"
)

type Marketer struct {
	controller.Base
}

func (m *Marketer) Index(c *gin.Context) {
	data := market.Find(c.Query("organize"), c.Query("symbol"))
	m.SuccessResponse(c, data)
}

//新增行情结构体
type Store struct {
	Organize   string `json:"organize" binding:"required"`
	MarketType int8   `json:"market_type" binding:"required,min=1,max=4"`
	Symbol     string `json:"symbol" binding:"required"`
}

func (m *Marketer) Store(c *gin.Context) {
	var store Store

	if err := c.ShouldBindJSON(&store); err != nil {
		m.ValidateResponse(c, err.Error())
		return
	}

	logic := &marketLogic.Store{
		Organize:   store.Organize,
		MarketType: store.MarketType,
		Symbol:     store.Symbol,
	}

	if err := logic.Store(); err != nil {
		m.ValidateResponse(c, err.Error())
		return
	}

	m.SuccessResponse(c, nil)
}
