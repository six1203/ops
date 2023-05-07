package main

import (
	"ops/app/shop"
	"ops/app/user"
	"ops/routers"
)

func main() {
	routers.Include(shop.Routers, user.Routers)
	r := routers.Init()
	r.Run()
}
