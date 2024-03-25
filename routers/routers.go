package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {

	//全局检查
	// r.Use()

	//初始化课程的路由
	initApi(r)

	//初始化API的路由
	initCourse(r)

}
