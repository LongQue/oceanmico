package client

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"oceanmico/proto/user"
)

var (
	userClient user.UserService
)

func GetUserClient() user.UserService {
	if userClient == nil {
		consulRegistry := consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"),
		)
		service := micro.NewService(
			micro.Registry(consulRegistry),
		)

		//获取名为go.micro.ocean.rand的客户端
		userClient =user.NewUserService("go.micro.ocean.user",service.Client())
	}
	return userClient
}