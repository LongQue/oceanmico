package handler

import (
	"context"
	"oceanmico/proto/user"
	"oceanmico/user-service/service"
)

type handler struct {
}

func (h handler) UserRegistry(ctx context.Context, request *user.UserRegistryRequest, response *user.UserLoginResponse) error {
	*response = service.RegisterUser(request.GetUsername(), request.GetPassword(), request.GetEmail())
	return nil
}

func (h handler) UserLogin(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	*response = service.UserLogin(request.GetUsername(), request.GetPassword())
	return nil
}

func Handler() user.UserHandler {
	return handler{}
}
