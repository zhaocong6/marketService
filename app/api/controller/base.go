package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Base struct{}

//返回success response
func (b *Base) SuccessResponse(c *gin.Context, data interface{}) {
	if reflect.ValueOf(data).IsNil() {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	}
	c.JSON(http.StatusOK, data)
}

//返回验证错误信息
func (b *Base) ValidateResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}
