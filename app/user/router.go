package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/user", userHandler)
	e.GET("/userinfo", userInfoHandler)
}
