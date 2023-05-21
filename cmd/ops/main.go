package main

import "ops/internal/routers"

func main() {
	r := routers.InitRouter()
	r.Run(":8000")

}
