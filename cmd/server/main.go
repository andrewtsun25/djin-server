package main

import (
	"context"
	"djin-server/internal/server"
	"fmt"
	servicev1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const ServiceAccountJson = "configs/firebase/djin-dev-003b126063d6.json"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}
	ctx := context.Background()
	serviceAccount := option.WithCredentialsFile(ServiceAccountJson)
	if _, err := os.ReadFile(ServiceAccountJson); err != nil {
		log.Printf("Cannot read from account: %s", ServiceAccountJson)
	}
	djinServiceServer := server.NewDjinServiceServer(ctx, serviceAccount)
	grpcServer := grpc.NewServer()
	servicev1.RegisterDjinServiceServer(grpcServer, djinServiceServer)
	log.Println("Listening on", listenOn)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}
