package server

import (
	"context"
	"os"

	proto "github.com/vivekmarakana/k8s-grpc-gateway/proto"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() proto.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, req *proto.RequestMessage) (*proto.ResponseMessage, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	if req.Message == "" {
		return &proto.ResponseMessage{
			Host:    hostname,
			Message: ":(",
		}, nil
	}

	return &proto.ResponseMessage{
		Host:    hostname,
		Message: req.Message,
	}, nil
}
