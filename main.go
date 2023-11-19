package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//创建一个/hello的请求路径方式，当客户端请求这个路径返回码200，会执行下面的函数
	//打印"message" : "Hello  world"
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello  world",
		})
	})
	//启动http服务，默认是在0.0.0.0:8080启动服务
	//下面是指定在9090端口号启动http服务
	r.Run(":9090")
}
