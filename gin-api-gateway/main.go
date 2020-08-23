package main

import (
	"github.com/gin-gonic/gin"
	"oceanmico/gin-api-gateway/client"
	"oceanmico/gin-api-gateway/handler"
)
func main() {
	//1、web http进来
	r:=gin.Default()
	v1 :=r.Group("/v1")
	//TODO 解决回调函数
	//2、获取被调用的客户端
	randClient:=client.GetRandClient()
	userClient:=client.GetUserClient()
	//3、封装一层
	apiHandler:=handler.GetAPIHandler(randClient,userClient)

	//4、执行处理方法，类似于springboot  service
	v1.GET("/v1",apiHandler.Rand)

	user := v1.Group("/user")
	user.POST("/registry",apiHandler.RegistryUser)

	if err := r.Run();err!= nil{
		panic(err)
	}
}
