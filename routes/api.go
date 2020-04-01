package routes

import (
	"github.com/gin-gonic/gin"
	"ws/marketApi/app/api/controller/market"
)

func InitApi(r *gin.Engine) {
	g := r.Group("/api/v1")

	marketC := &market.Marketer{}
	g.GET("/market", marketC.Index)
	g.POST("/market", marketC.Store)
}
