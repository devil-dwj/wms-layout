gen: gen_router

gen_router:
	protoc --proto_path=. \
			--proto_path=third_party \
			--goc_out=paths=source_relative:. \
			--go-routernew_out=paths=source_relative:. \
			api/helloworld/*.proto \