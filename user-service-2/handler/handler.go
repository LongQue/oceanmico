package handler

import (
	"context"
	"oceanmico/proto/user"
	"oceanmico/user-service-2/service"
)

type handler struct {

}

func (h handler) UserRegistry(ctx context.Context, request *user.UserRegistryRequest, response *user.UserLoginResponse) error {
	*response = service.RegisterUser(request.GetUsername(), request.GetPassword(), request.GetEmail())
	return nil
}

func (h handler) UserLogin(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	panic("implement me")
}

func Handler() user.UserHandler {
	return handler{}
}

