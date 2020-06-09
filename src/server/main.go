package server

import (
	"context"
	"fmt"
	"net"

	proto "github.com/vivekmarakana/k8s-grpc-gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			fmt.Printf("Failed to close %s %s: %v\n", network, address, err)
		}
	}()

	s := grpc.NewServer()
	proto.RegisterEchoServiceServer(s, newEchoServer())
	reflection.Register(s)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}
