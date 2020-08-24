package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oceanmico/proto/rand"
	"oceanmico/proto/user"
)

type APIHandler struct {
	randClient rand.RandService
	userClient user.UserService
}

func GetAPIHandler(randService rand.RandService, userService user.UserService) *APIHandler {
	return &APIHandler{
		randClient: randService,
		userClient: userService,
	}
}

//Get请求
func (s *APIHandler) Rand(ctx *gin.Context) {
	request := rand.RandRequest{}
	//请求参数绑定到RandRequest
	ctx.ShouldBindQuery(&request)
	//使用客户端的GetRand方法处理
	response, err := s.randClient.GetRand(context.Background(), &request)
	if err != nil {
		log.Println(err)
	}
	username, _ := ctx.Get("user")
	//构造响应返回web
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"result":   response.GetResult(),
	})
}

//Post请求
func (s *APIHandler) RegistryUser(ctx *gin.Context) {
	request := user.UserRegistryRequest{}

	ctx.ShouldBindJSON(&request)

	responst, _ := s.userClient.UserRegistry(context.Background(), &request)

	ctx.JSON(http.StatusOK, gin.H{
		"message": responst.GetMessage(),
		"token":   responst.GetToken(),
	})
}

func (s *APIHandler) LoginUser(ctx *gin.Context) {
	request := user.UserLoginRequest{}

	_ = ctx.ShouldBindJSON(&request)

	response, _ := s.userClient.UserLogin(context.Background(), &request)

	ctx.JSON(http.StatusOK, gin.H{
		"message": response.GetMessage(),
		"token":   response.GetToken(),
	})

}
