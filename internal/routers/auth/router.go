package auth

import (
	"github.com/gin-gonic/gin"
	"ops/internal/controllers/auth"
)

func Routers(r *gin.Engine) {
	routeGroup := r.Group("/auth")
	{
		routeGroup.POST("/login", auth.AuthController{}.Auth)
	}
}
