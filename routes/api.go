package routes

import (
	"github.com/gin-gonic/gin"
	"marketApi/app/api/controller"
)

func InitApi(r *gin.Engine) {
	g := r.Group("/api/v1")

	marketC := &controller.Marketer{}
	g.GET("/market", marketC.Index)
	g.POST("/market", marketC.Store)
}
