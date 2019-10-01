
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

build:
	# go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	# go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	# go get -v -u github.com/golang/protobuf/protoc-gen-go
	protoc -I/usr/local/include -I. -I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:. *.proto
