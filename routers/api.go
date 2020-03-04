package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Api 路由注册方法
func ApiKernel(router *gin.RouterGroup) {
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello work")
	})
}
