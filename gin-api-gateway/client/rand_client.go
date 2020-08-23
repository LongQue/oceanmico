package client

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"oceanmico/proto/rand"
)

var (
	randClient rand.RandService
)

func GetRandClient() rand.RandService {
	if randClient == nil {
		consulRegistry := consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"),
		)
		service := micro.NewService(
			micro.Registry(consulRegistry),
		)

		//获取名为go.micro.ocean.rand的客户端
		randClient=rand.NewRandService("go.micro.ocean.rand",service.Client())
	}
	return randClient
}
