package routers

import (
	"gin-test/routers/web"

	"github.com/gin-gonic/gin"
)

func initApi(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping) //这里web.ping是传入一个方法，不是调用一个方法
	v1.POST("/login", web.Login)
	v1.POST("/registe", web.Registe)

}
