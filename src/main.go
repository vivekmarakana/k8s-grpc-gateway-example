package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vivekmarakana/k8s-grpc-gateway/src/gateway"
	"github.com/vivekmarakana/k8s-grpc-gateway/src/server"
)

var (
	serverAddr  = flag.String("serverAddr", ":9090", "endpoint of the gRPC server")
	gatewayAddr = flag.String("gatewayAddr", ":8080", "endpoint of the gRPC gateway")
	network     = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		cancel()
		fmt.Println("Exiting server on ", sig)
		os.Exit(0)
	}()

	go func() {
		fmt.Println("Starting HTTP gateway on", *gatewayAddr)
		if err := gateway.Run(ctx, *gatewayAddr, "localhost"+*serverAddr); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Starting gRPC server on", *serverAddr)
	if err := server.Run(ctx, *network, *serverAddr); err != nil {
		log.Fatal(err)
	}
}
