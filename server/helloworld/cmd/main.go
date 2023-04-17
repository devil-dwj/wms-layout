package main

import (
	"fmt"

	"github.com/devil-dwj/wms-bee/log/zap"
	apihelloworld "github.com/devil-dwj/wms-layout/api/helloworld"
	"github.com/devil-dwj/wms-layout/server/helloworld/conf"
	servicehelloworld "github.com/devil-dwj/wms-layout/server/helloworld/service/helloworld"
	"github.com/devil-dwj/wms/app"
	"github.com/devil-dwj/wms/config"
	"github.com/devil-dwj/wms/log"
	"github.com/devil-dwj/wms/middleware/logging"
	"github.com/devil-dwj/wms/middleware/recovery"
	"github.com/devil-dwj/wms/middleware/validate"
	"github.com/devil-dwj/wms/runtime/http"
)

func main() {
	c := config.New(
		config.WithSource("../conf/config.json"),
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
			validate.Validator(),
		),
	)

	helloworldServer := servicehelloworld.NewHelloWorldServer()
	apihelloworld.RegisterGreeterRouter(httpSvr, helloworldServer)

	app := app.New(app.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatalf("app run err : %+v", err)
	}
}
