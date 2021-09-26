
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif


APP_TAGS ?= ""
BUILD_GOOS ?= linux
BUILD_GOARCH ?= amd64
BUILD_CGO_ENABLED ?= 0
DOCKER_BUILDKIT ?= 1
DOCKER_CONTAINER_IMAGE="demdxx/udetect-server-example:latest"

COMMIT_NUMBER ?= $(shell git rev-parse --short HEAD)
BUILD_VERSION ?= $(shell git describe --exact-match --tags $(git log -n1 --pretty='%h'))

OS_LIST = linux darwin
ARCH_LIST = amd64 arm64

ifeq ($(BUILD_VERSION),)
	BUILD_VERSION := commit-$(COMMIT_NUMBER)
endif


.PHONY: lint
lint: golint ## Run linter checks


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: test
test: ## Run unit tests
	go test -v -race ./...


.PHONY: golint
golint: $(GOLANGLINTCI)
	# golint -set_exit_status ./...
	golangci-lint run -v ./...


.PHONY: fmt
fmt: ## Run formatting code
	@echo "Fix formatting"
	@gofmt -w ${GO_FMT_FLAGS} $$(go list -f "{{ .Dir }}" ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi


.PHONY: build-proto
build-proto: ## Build protocol objects from protobuf defenition
	# go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	# go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	# go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	# go install github.com/golang/protobuf/protoc-gen-go
	# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I/usr/local/include -I. -I$(GOPATH)/src \
		-I$(GOPATH)/pkg/mod/github.com/googleapis/api-common-protos@v0.1.0 \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt logtostderr=true \
		--swagger_out=logtostderr=true:. \
		protocol/*.proto


.PHONY: help
help: ## Print help description
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: build-example
build-example: ## Build example server
	@echo "Build example server"
	@rm -rf .build/server
	GOOS=${BUILD_GOOS} GOARCH=${BUILD_GOARCH} CGO_ENABLED=${BUILD_CGO_ENABLED} \
		go build -ldflags "-X main.buildDate=`date -u +%Y%m%d.%H%M%S` -X main.buildCommit=${COMMIT_NUMBER} -X main.appVersion=${BUILD_VERSION}" \
			-tags ${APP_TAGS} -o ".build/server" examples/server/main.go


.PHONY: run-server
run-server: build-example ## Run example server
	DOCKER_BUILDKIT=${DOCKER_BUILDKIT} docker compose -f docker/develop/docker-compose.yml run --rm server


.PHONY: build-prod-example
build-prod-example:
	@echo "Build server application"
	@rm -rf .build/
	cd examples/server/; \
	for os in $(OS_LIST); do \
		for arch in $(ARCH_LIST); do \
			echo "Build $$os/$$arch"; \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=${BUILD_CGO_ENABLED} go build \
				-ldflags "-X main.buildDate=`date -u +%Y%m%d.%H%M%S` -X main.buildCommit=${COMMIT_NUMBER} -X main.appVersion=${BUILD_VERSION}" \
				-o ../../.build/$$os/$$arch/server ./main.go; \
		done \
	done


.PHONY: docker-build
docker-build: build-prod-example ## Build docker example server
	echo "Build server docker image"
	DOCKER_BUILDKIT=${DOCKER_BUILDKIT} docker buildx build \
		--push --platform linux/amd64,linux/arm64,darwin/amd64,darwin/arm64 \
		-t ${DOCKER_CONTAINER_IMAGE} -f docker/server-example.dockerfile .


.DEFAULT_GOAL := help
