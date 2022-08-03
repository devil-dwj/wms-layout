package main

import (
	"fmt"

	"github.com/devil-dwj/wms-bee/log/zap"
	apihelloworld "github.com/devil-dwj/wms-template/api/helloworld"
	"github.com/devil-dwj/wms-template/conf"
	servicehelloworld "github.com/devil-dwj/wms-template/service/helloworld"
	"github.com/devil-dwj/wms/app"
	"github.com/devil-dwj/wms/config"
	"github.com/devil-dwj/wms/log"
	"github.com/devil-dwj/wms/middleware/logging"
	"github.com/devil-dwj/wms/middleware/recovery"
	"github.com/devil-dwj/wms/runtime/http"
)

func main() {
	c := config.New(
		config.WithSource("config.json"),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var cc conf.Config
	if err := c.Scan(&cc); err != nil {
		panic(err)
	}

	zlog := zap.NewLogger()
	log.SetLogger(zlog)

	httpSvr := http.NewServer(fmt.Sprintf(":%d", cc.Port),
		http.Middleware(
			recovery.Recovery(),
			logging.Logging,
		),
	)

	helloworldServer := servicehelloworld.NewHelloWorldServer()
	apihelloworld.RegisterGreeterRouter(httpSvr, helloworldServer)

	app := app.New(app.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatalf("app run err : %+v", err)
	}
}
