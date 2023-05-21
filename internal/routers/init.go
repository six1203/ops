package routers

import (
	"github.com/gin-gonic/gin"
	"ops/internal/middleware"
	"ops/internal/routers/auth"
	"ops/internal/routers/user"
)

type Option func(r *gin.Engine)

var options []Option

func Include(opts ...Option) {
	options = append(options, opts...)
}

func InitRouter() *gin.Engine {
	r := gin.New()

	middleware.InitMiddleware(r)

	Include(auth.Routers, user.Routers)

	for _, opt := range options {
		opt(r)
	}
	return r
}
