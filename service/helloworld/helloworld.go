package helloworld

import (
	"context"

	"github.com/devil-dwj/wms-template/api/helloworld"
	"github.com/devil-dwj/wms-template/service"
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
	return &helloworld.HelloReply{
		Message: "hello " + req.GetName(),
	}, nil
}
