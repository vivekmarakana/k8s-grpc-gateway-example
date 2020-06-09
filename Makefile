build: dep compile fmt vet

dep:
	go mod tidy

compile:
	go build -o bin/server src/main.go 

fmt:
	go fmt src/main.go 

vet:
	go vet src/main.go 

gp:
	protoc -I. -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --go_out=plugins=grpc,paths=source_relative:. ./proto/echo_service.proto
	protoc -I. -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --grpc-gateway_out=logtostderr=true:. ./proto/echo_service.proto

dockerbuild: gp
	docker build . -t $(DOCKER_HUB_USERNAME)/k8s-grpc-gateway:latest
