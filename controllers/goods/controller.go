package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个goodsController结构体
type GoodsController struct {
	BaseController
}

// 把如下方法绑定给goodsController结构体
func (g GoodsController) GoodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

func (g GoodsController) GoodsListHandler(c *gin.Context) {
	//c.String(http.StatusOK, "goodsListHandler")
	g.Success(c)
}
func (g GoodsController) UpdateGoodsHandler(c *gin.Context) {
	//c.String(http.StatusOK, "updateGoodsHandler")
	g.Error(c)
}
