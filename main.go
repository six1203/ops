package main

import (
	"github.com/gin-gonic/gin"
	"ops/routers"
)

func main() {
	r := gin.Default()
	routers.LoadGoodsRouters(r)
	routers.LoadOrderRouters(r)
	routers.LoadUserRouters(r)
	r.Run()
}
