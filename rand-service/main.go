package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"oceanmico/proto/rand"
	"oceanmico/rand-service/handler"
)

func main() {
	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	service := micro.NewService(
		micro.Name("go.micro.ocean.rand"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	// 暴露接口到对应服务

	_ = rand.RegisterRandHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}

}
