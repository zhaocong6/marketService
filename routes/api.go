package routes

import (
	"github.com/gin-gonic/gin"
	"marketApi/app/api/controller"
)

func InitApi(r *gin.Engine) {
	g := r.Group("/api/v1")

	marketC := &controller.Marketer{}
	g.GET("/markets/:organize/:symbol", marketC.Index)
	g.GET("/markets", marketC.List)
	g.POST("/markets", marketC.Store)
}
