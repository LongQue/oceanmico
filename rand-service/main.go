package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"oceanmico/proto/rand"
	"oceanmico/rand-service/handler"
	"oceanmico/rand-service/middleware"
)

func main() {
	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	//初始化限流器
	middleware.InitSentinel()
	service := micro.NewService(
		micro.Name("go.micro.ocean.rand"),
		micro.Registry(consulRegistry),
		//添加限流器
		micro.WrapHandler(middleware.LimitingWrapper()),
	)

	service.Init()

	// 暴露接口到对应服务

	_ = rand.RegisterRandHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}

}
