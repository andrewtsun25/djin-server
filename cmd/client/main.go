package main

import (
	"context"
	"fmt"
	grpcEntity "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"
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
	ctx := context.Background()

	// Test ListEducationsByType

	// Coding
	listEdusResponse, err := djinServiceClient.ListEducationsByType(ctx, &rpc.ListEducationsByTypeRequest{
		EducationType: grpcEntity.EducationType_EDUCATION_TYPE_CODING,
	})
	if err != nil {
		return fmt.Errorf("[ListEducationsByType] Failed to list coding educations: %w\n\n", err)
	}
	log.Printf("[ListEducationsByType] Coding Educations: %v\n\n", listEdusResponse)

	// Music
	listEdusResponse, err = djinServiceClient.ListEducationsByType(ctx, &rpc.ListEducationsByTypeRequest{
		EducationType: grpcEntity.EducationType_EDUCATION_TYPE_MUSIC,
	})
	if err != nil {
		return fmt.Errorf("[ListEducationsByType] Failed to list music educations: %w\n\n", err)
	}
	log.Printf("[ListEducationsByType] Music Educations: %v\n\n", listEdusResponse)

	// Test ListEmployments
	listEmploymentsResponse, err := djinServiceClient.ListEmployments(ctx, &rpc.ListEmploymentsRequest{})
	if err != nil {
		return fmt.Errorf("[ListEmployments] Failed to list employments: %w\n\n", err)
	}
	log.Printf("[ListEmployments] Employments: %v\n\n", listEmploymentsResponse)

	// Test Get Organization
	getOrgResponse, err := djinServiceClient.GetOrganizationById(ctx, &rpc.GetOrganizationByIdRequest{
		Id: "amazon",
	})
	if err != nil {
		return fmt.Errorf("[GetOrganizationById] Failed to get organization 'amazon': %w", err)
	}
	log.Printf("[GetOrganizationById] Organization 'amazon': %v\n\n", getOrgResponse)

	return nil
}
