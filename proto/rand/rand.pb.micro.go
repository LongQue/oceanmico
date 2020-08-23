// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/rand/rand.proto

package rand

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Rand service

func NewRandEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Rand service

type RandService interface {
	GetRand(ctx context.Context, in *RandRequest, opts ...client.CallOption) (*RandResponse, error)
}

type randService struct {
	c    client.Client
	name string
}

func NewRandService(name string, c client.Client) RandService {
	return &randService{
		c:    c,
		name: name,
	}
}

func (c *randService) GetRand(ctx context.Context, in *RandRequest, opts ...client.CallOption) (*RandResponse, error) {
	req := c.c.NewRequest(c.name, "Rand.GetRand", in)
	out := new(RandResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rand service

type RandHandler interface {
	GetRand(context.Context, *RandRequest, *RandResponse) error
}

func RegisterRandHandler(s server.Server, hdlr RandHandler, opts ...server.HandlerOption) error {
	type rand interface {
		GetRand(ctx context.Context, in *RandRequest, out *RandResponse) error
	}
	type Rand struct {
		rand
	}
	h := &randHandler{hdlr}
	return s.Handle(s.NewHandler(&Rand{h}, opts...))
}

type randHandler struct {
	RandHandler
}

func (h *randHandler) GetRand(ctx context.Context, in *RandRequest, out *RandResponse) error {
	return h.RandHandler.GetRand(ctx, in, out)
}
