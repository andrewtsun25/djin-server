package main

import (
	"context"
	"fmt"
	rpc "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/rpc/v1"
	service "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/grpc"
	"log"
)

func main() {
	if err := clientRun(); err != nil {
		log.Fatal(err)
	}
}

func clientRun() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to DjinService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	djinServiceClient := service.NewDjinServiceClient(conn)
	resp, err := djinServiceClient.GetOrganizationById(context.Background(), &rpc.GetOrganizationByIdRequest{
		Id: "amazon",
	})
	if err != nil {
		return fmt.Errorf("failed to GetOrganizationById: %w", err)
	}

	log.Printf("Successfully Got Organization: %v", resp)
	return nil
}
