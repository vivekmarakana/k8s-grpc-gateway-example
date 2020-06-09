package server

import (
	"context"

	proto "github.com/vivekmarakana/k8s-grpc-gateway/proto"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() proto.EchoServiceServer {
	return new(echoServer)
}

func (s *echoServer) Echo(ctx context.Context, req *proto.RequestMessage) (*proto.ResponseMessage, error) {
	if req.Message == "" {
		return &proto.ResponseMessage{
			You:     "you",
			Me:      "me",
			Message: ":(",
		}, nil
	}

	return &proto.ResponseMessage{
		You:     "you",
		Me:      "me",
		Message: req.Message,
	}, nil
}
