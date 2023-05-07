package routers

import (
	"github.com/gin-gonic/gin"
	"ops/controllers/goods"
)

func LoadGoodsRouters(r *gin.Engine) {
	r1 := r.Group("/goods")
	{
		r1.GET("/", goods.GoodsController{}.GoodsHandler)
		r1.GET("/goods-list", goods.GoodsController{}.GoodsListHandler)
		r1.GET("/updaters", goods.GoodsController{}.UpdateGoodsHandler)
	}
}
