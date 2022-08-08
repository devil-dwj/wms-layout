GOHOSTOS:=$(shell go env GOHOSTOS)

ifeq ($(GOHOSTOS), windows)
	API_PROTOS=$(shell bash.exe -c "find api -name *.proto")
else
	API_PROTOS=$(shell find api -name *.proto)
endif

.PHONY: gen
gen: gen_router

.PHONY: gen_router
gen_router:
	protoc --proto_path=. \
			--proto_path=third_party \
			--goc_out=paths=source_relative:. \
			--go-router_out=paths=source_relative:. \
			--go-errors_out=paths=source_relative:. \
			$(API_PROTOS)