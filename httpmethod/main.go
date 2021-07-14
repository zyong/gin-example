package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号
}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "Http get method")
}

func posting(c *gin.Context) {
	c.String(http.StatusOK, "Http post method")
}

func deleting(c *gin.Context) {
	c.String(http.StatusOK, "Http delete method")
}

func putting(c *gin.Context) {
	c.String(http.StatusOK, "Http put method")
}
