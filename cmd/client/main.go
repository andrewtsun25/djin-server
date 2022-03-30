package main

import (
	"context"
	"fmt"
	rpcv1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/rpc/v1"
	"log"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	servicev1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := client_run(); err != nil {
		log.Fatal(err)
	}
}

func client_run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to DjinService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	djinClient := servicev1.NewDjinServiceClient(conn)
	resp, err := djinClient.GetOrganizationById(context.Background(), &rpcv1.GetOrganizationByIdRequest{
		Id: "amazon",
	})
	if err != nil {
		return fmt.Errorf("failed to GetOrganizationById: %w", err)
	}

	log.Printf("Successfully Got Organization: %v", resp)
	return nil
}
