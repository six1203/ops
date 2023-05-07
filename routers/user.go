package routers

import (
	"github.com/gin-gonic/gin"
	"ops/controllers/user"
)

func LoadUserRouters(r *gin.Engine) {
	r1 := r.Group("/user")
	{
		r1.GET("/", user.UserController{}.UserHandler)
		r1.GET("/user-list", user.UserController{}.UserListHandler)
		r1.GET("/updater", user.UserController{}.UpdateUserHandler)
	}
}
