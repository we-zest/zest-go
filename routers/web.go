package routers

import (
	"github.com/gin-gonic/gin"
	"zest-go/app/controller/user"
)

// Web 路由注册方法
func WebKernel(router *gin.Engine) {
	// 用户创建
	router.POST("/", user.Create)
}
