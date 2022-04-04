package main

import (
	"context"
	"djin-server/internal/server"
	"fmt"
	service "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"net"
)

const ServiceAccountJson = "configs/firebase/djin-dev-003b126063d6.json"
const PortAddress = "127.0.0.1:8080"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", PortAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", PortAddress, err)
	}
	ctx := context.Background()
	serviceAccount := option.WithCredentialsFile(ServiceAccountJson)
	djinServiceServer := server.NewDjinServiceServer(ctx, serviceAccount)
	if djinServiceServer == nil {
		return fmt.Errorf("failed to instantiate djinServiceServer")
	}
	defer func() {
		_ = djinServiceServer.Close()
	}()
	grpcServer := grpc.NewServer()
	service.RegisterDjinServiceServer(grpcServer, djinServiceServer)
	log.Println("Listening on", PortAddress)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}
