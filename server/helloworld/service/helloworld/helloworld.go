package helloworld

import (
	"context"

	"github.com/devil-dwj/wms-layout/api/helloworld"
	"github.com/devil-dwj/wms-layout/server/helloworld/service"
)

type HelloWorldServer struct {
	opt service.Options
}

func NewHelloWorldServer(opts ...service.Option) *HelloWorldServer {
	o := service.Options{}
	for _, opt := range opts {
		opt(&o)
	}
	return &HelloWorldServer{
		opt: o,
	}
}

func (s *HelloWorldServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if req.GetName() == "" {
		return nil, helloworld.ErrorNotNil("name cannot be nil")
	}

	return &helloworld.HelloReply{
		Message: "hello " + req.GetName(),
	}, nil
}
