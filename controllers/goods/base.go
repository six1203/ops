package goods

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (b BaseController) Success(c *gin.Context) {
	c.String(200, "商品成功")
}

func (b BaseController) Error(c *gin.Context) {
	c.String(500, "商品失败")
}
