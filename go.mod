module github.com/devil-dwj/wms-template

go 1.16

require (
	github.com/devil-dwj/wms v0.0.0-20221020070228-f717c402423d
	github.com/devil-dwj/wms-bee/log/zap v0.0.0-20220801064816-d42c108ca68f
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	google.golang.org/genproto v0.0.0-20221027153422-115e99e71e1c
	google.golang.org/protobuf v1.28.1
)

replace github.com/devil-dwj/wms => ../wms
