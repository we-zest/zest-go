package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zest-go/app/helper"
	"zest-go/routers"
)

// 初始化路由方法
func setupRouter() {
	router := gin.Default()
	// 设置静态文件
	router.Static("/static", "./statics")
	// 设置 ico
	router.StaticFile("/favicon.ico", "./statics/favicon.ico")
	// 设置 FileSystem 可处理的文件目录
	router.StaticFS("/uploads", http.Dir(helper.RootPath()+"uploads/"))

	// Api 路由
	apiRouter := router.Group("/api")
	{
		routers.ApiKernel(apiRouter)
	}

	// Web 路由
	routers.WebKernel(router)

	_ = router.Run()
}