package main

import (
	"github.com/gin-gonic/gin"
	"oceanmico/gin-api-gateway/client"
	"oceanmico/gin-api-gateway/handler"
	"oceanmico/gin-api-gateway/middleware"
)

func main() {
	//1、web http进来
	r := gin.Default()
	v1 := r.Group("/v1")
	//2、获取被调用的客户端
	randClient := client.GetRandClient()
	userClient := client.GetUserClient()
	//3、封装一层
	apiHandler := handler.GetAPIHandler(randClient, userClient)

	//4、执行处理方法，类似于springboot  service

	user := v1.Group("/user")
	user.POST("/registry", apiHandler.RegistryUser)
	user.POST("/login", apiHandler.LoginUser)
	user.GET("/rand", middleware.AuthMiddleware(), apiHandler.Rand)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
