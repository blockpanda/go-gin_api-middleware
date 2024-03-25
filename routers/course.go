package routers

import (
	"gin-test/routers/middleware"
	"gin-test/routers/web"

	"github.com/gin-gonic/gin"
)

func initCourse(r *gin.Engine) {

	v1 := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)

	v1.POST("/course", web.Create)
	v1.GET("/course", web.Get)
	v1.PUT("/course", web.Edit)
	v1.DELETE("/course", web.Delete)

	v2 := r.Group("/v2")
	v2.POST("/course", web.Create2) //版本升级后的v2

}
