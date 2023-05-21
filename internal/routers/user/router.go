package user

import (
	"github.com/gin-gonic/gin"
	"ops/internal/controllers/user"
)

func Routers(r *gin.Engine) {
	routeGroup := r.Group("/user")
	{
		routeGroup.GET("/list", user.UserController{}.QueryUserList)
		routeGroup.GET("/info", user.UserController{}.GetUserById)
	}
}
