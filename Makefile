GOHOSTOS:=$(shell go env GOHOSTOS)

ifeq ($(GOHOSTOS), windows)
	API_PROTO_FILES=$(shell bash.exe -c "find api -name *.proto")
else
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

gen: gen_router

gen_router:
	protoc --proto_path=. \
			--proto_path=third_party \
			--goc_out=paths=source_relative:. \
			--go-router_out=paths=source_relative:. \
			$(API_PROTO_FILES)