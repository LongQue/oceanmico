package handler

import (
	"context"
	"oceanmico/proto/rand"
	"oceanmico/rand-service/service"
)

type handler struct {

}

func (h handler) GetRand(ctx context.Context, request *rand.RandRequest, response *rand.RandResponse) error {
	max:=request.GetMax()
	response.Result=service.GetRand(max)
	return nil
}
//生成实例(?) 供外部调用上面的方法
func Handler() rand.RandHandler{
	return handler{}
}


