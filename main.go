package main

import (
	"gin-test/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	r.Run(":9090")
}
