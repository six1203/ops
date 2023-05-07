package routers

import "ops/controllers/order"
import (
	"github.com/gin-gonic/gin"
)

func LoadOrderRouters(r *gin.Engine) {
	r1 := r.Group("/Oder")
	{
		r1.GET("/", order.OrderController{}.OrderHandler)
		r1.GET("/order-list", order.OrderController{}.OrderListHandler)
		r1.GET("/updater", order.OrderController{}.UpdateOrderHandler)
	}
}
