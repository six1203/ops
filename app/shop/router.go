package shop

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/shop", shopHandler)
	e.GET("/shopinfo", shopInfoHandler)
}
