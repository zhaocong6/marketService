package market

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaocong6/market"
	"net/http"
)

func Index(c *gin.Context) {
	m := market.Find(c.Query("organize"), c.Query("symbol"))
	c.JSON(http.StatusOK, m)
}