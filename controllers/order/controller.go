package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个OrderController结构体
type OrderController struct {
}

// 把如下方法绑定给OrderController结构体
func (o OrderController) OrderHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "dylan",
	})
}

func (o OrderController) OrderListHandler(c *gin.Context) {
	c.String(http.StatusOK, "orderListHandler")
}
func (o OrderController) UpdateOrderHandler(c *gin.Context) {
	c.String(http.StatusOK, "updateOrderHandler")
}
