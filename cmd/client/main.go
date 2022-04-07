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

var (
	martialArtsStyles = []string{"hdgd", "itfTkd", "wtTkd"}
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

	// Coding Educations
	listEdusResponse, err := djinServiceClient.ListEducationsByType(ctx, &rpc.ListEducationsByTypeRequest{
		EducationType: grpcEntity.EducationType_EDUCATION_TYPE_CODING,
	})
	if err != nil {
		return fmt.Errorf("[ListEducationsByType] Failed to list coding educations: %w\n\n", err)
	}
	log.Printf("[ListEducationsByType] Coding Educations: %v\n\n", listEdusResponse)

	// Music Educations
	listEdusResponse, err = djinServiceClient.ListEducationsByType(ctx, &rpc.ListEducationsByTypeRequest{
		EducationType: grpcEntity.EducationType_EDUCATION_TYPE_MUSIC,
	})
	if err != nil {
		return fmt.Errorf("[ListEducationsByType] Failed to list music educations: %w\n\n", err)
	}
	log.Printf("[ListEducationsByType] Music Educations: %v\n\n", listEdusResponse)

	// Employments
	listEmploymentsResponse, err := djinServiceClient.ListEmployments(ctx, &rpc.ListEmploymentsRequest{})
	if err != nil {
		return fmt.Errorf("[ListEmployments] Failed to list employments: %w\n\n", err)
	}
	log.Printf("[ListEmployments] Employments: %v\n\n", listEmploymentsResponse)

	// HbvResearchPapers
	listHbvResearchPapersResponse, err := djinServiceClient.ListHbvResearchPapers(ctx, &rpc.ListHbvResearchPapersRequest{})
	if err != nil {
		return fmt.Errorf("[ListHbvResearchPapers] Failed to list HBV research papers: %w\n\n", err)
	}
	log.Printf("[ListHbvResearchPapers] HBV research papers: %v\n\n", listHbvResearchPapersResponse)

	// Holistic Office
	listHolisticOfficeLinksResponse, err := djinServiceClient.ListHolisticOfficeLinks(ctx, &rpc.ListHolisticOfficeLinksRequest{})
	if err != nil {
		return fmt.Errorf("[ListHolisticOfficeLinks] Failed to list Holistic Office links: %w\n\n", err)
	}
	log.Printf("[ListHolisticOfficeLinks] Holistic Office links: %v\n\n", listHolisticOfficeLinksResponse)

	listHolisticOfficeModulesResponse, err := djinServiceClient.ListHolisticOfficeModules(ctx, &rpc.ListHolisticOfficeModulesRequest{})
	if err != nil {
		return fmt.Errorf("[ListHolisticOfficeModules] Failed to list Holistic Office modules: %w\n\n", err)
	}
	log.Printf("[ListHolisticOfficeModules] Holistic Office modules: %v\n\n", listHolisticOfficeModulesResponse)

	// Martial Arts Styles
	for _, martialArtsStyle := range martialArtsStyles {
		getMartialArtsStyleByIdResponse, err := djinServiceClient.GetMartialArtsStyleById(ctx, &rpc.GetMartialArtsStyleByIdRequest{
			Id: martialArtsStyle,
		})
		if err != nil {
			return fmt.Errorf("[GetMartialArtsStyleById] Failed to list martial arts style for id '%s': %w\n\n", martialArtsStyle, err)
		}
		log.Printf("[GetMartialArtsStyleById] Martial Arts Style for id '%s': %v\n\n", martialArtsStyle, getMartialArtsStyleByIdResponse)
	}

	// Music
	listInstrumentsResponse, err := djinServiceClient.ListInstruments(ctx, &rpc.ListInstrumentsRequest{})
	if err != nil {
		return fmt.Errorf("[ListInstruments] Failed to list music instruments: %w\n\n", err)
	}
	log.Printf("[ListInstruments] Music Instruments : %v\n\n", listInstrumentsResponse)

	listMusicScoresResponse, err := djinServiceClient.ListMusicScores(ctx, &rpc.ListMusicScoresRequest{})
	if err != nil {
		return fmt.Errorf("[ListMusicScores] Failed to list music scores: %w\n\n", err)
	}
	log.Printf("[ListMusicScores] Music Scores : %v\n\n", listMusicScoresResponse)

	// Organizations
	getOrganizationByIdResponse, err := djinServiceClient.GetOrganizationById(ctx, &rpc.GetOrganizationByIdRequest{
		Id: "amazon",
	})
	if err != nil {
		return fmt.Errorf("[GetOrganizationById] Failed to get organization 'amazon': %w", err)
	}
	log.Printf("[GetOrganizationById] Organization 'amazon': %v\n\n", getOrganizationByIdResponse)

	// Projects
	listProjectsResponse, err := djinServiceClient.ListProjects(ctx, &rpc.ListProjectsRequest{})
	if err != nil {
		return fmt.Errorf("[ListProjects] Failed to list projects: %w\n\n", err)
	}
	log.Printf("[ListProjects] Projects : %v\n\n", listProjectsResponse)

	return nil
}
