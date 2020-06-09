package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	proto "github.com/vivekmarakana/k8s-grpc-gateway/proto"
	"google.golang.org/grpc"
)


// Run starts the example http gateway.
// "endpoint" is passed to register the gRPC server.
// "addr" is passed to net.Listen.
func Run(ctx context.Context, addr string, endpoint string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)
	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}
	
	err := proto.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, endpoint, grpcOpts)
	if err != nil {
		fmt.Printf("Failed to register endpoint server: %v", err)
		return err
	}

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Printf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Failed to listen and serve: %v", err)
		return err
	}

	return nil
}
