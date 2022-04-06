package server

import (
	"context"
	"djin-server/internal/db"
	rpc "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/rpc/v1"
	service "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type djinServiceServerImpl struct {
	service.UnimplementedDjinServiceServer
	firestoreDB *db.FirestoreDB
}

func NewDjinServiceServer(ctx context.Context, serviceAccount option.ClientOption) *djinServiceServerImpl {
	fdb := db.NewFirestoreDB(ctx, serviceAccount)
	if fdb == nil {
		return nil
	}
	return &djinServiceServerImpl{
		firestoreDB: fdb,
	}
}

func (s *djinServiceServerImpl) Close() error {
	return s.firestoreDB.Close()
}

// Educations

func (s *djinServiceServerImpl) ListEducationsByType(ctx context.Context, req *rpc.ListEducationsByTypeRequest) (*rpc.ListEducationsByTypeResponse, error) {
	educations, err := s.firestoreDB.ListEducationsByType(ctx, req.GetEducationType())
	if err != nil {
		return nil, err
	}
	return &rpc.ListEducationsByTypeResponse{
		Educations: educations,
	}, nil
}

// Employments

func (s *djinServiceServerImpl) ListEmployments(ctx context.Context, _ *rpc.ListEmploymentsRequest) (*rpc.ListEmploymentsResponse, error) {
	employments, err := s.firestoreDB.ListEmployments(ctx)
	if err != nil {
		return nil, err
	}
	return &rpc.ListEmploymentsResponse{
		Employments: employments,
	}, nil
}

// HBV Research Papers

func (s *djinServiceServerImpl) ListHbvResearchPapers(ctx context.Context, _ *rpc.ListHbvResearchPapersRequest) (*rpc.ListHbvResearchPapersResponse, error) {
	hbvResearchPapers, err := s.firestoreDB.ListHbvResearchPapers(ctx)
	if err != nil {
		return nil, err
	}
	return &rpc.ListHbvResearchPapersResponse{
		HbvResearchPapers: hbvResearchPapers,
	}, nil
}

// Holistic Office

func (s *djinServiceServerImpl) ListHolisticOfficeLinks(ctx context.Context, _ *rpc.ListHolisticOfficeLinksRequest) (*rpc.ListHolisticOfficeLinksResponse, error) {
	holisticOfficeLinks, err := s.firestoreDB.ListHolisticOfficeLinks(ctx)
	if err != nil {
		return nil, err
	}
	return &rpc.ListHolisticOfficeLinksResponse{
		HolisticOfficeLinks: holisticOfficeLinks,
	}, nil
}

func (s *djinServiceServerImpl) ListHolisticOfficeModules(ctx context.Context, _ *rpc.ListHolisticOfficeModulesRequest) (*rpc.ListHolisticOfficeModulesResponse, error) {
	holisticOfficeModules, err := s.firestoreDB.ListHolisticOfficeModules(ctx)
	if err != nil {
		return nil, err
	}
	return &rpc.ListHolisticOfficeModulesResponse{
		HolisticOfficeModules: holisticOfficeModules,
	}, nil
}

// Martial Arts

func (s *djinServiceServerImpl) GetMartialArtsStyleById(ctx context.Context, req *rpc.GetMartialArtsStyleByIdRequest) (*rpc.GetMartialArtsStyleByIdResponse, error) {
	martialArtsStyle, err := s.firestoreDB.GetMartialArtsStyleById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &rpc.GetMartialArtsStyleByIdResponse{
		MartialArtsStyle: martialArtsStyle,
	}, nil
}

// Music

func (s *djinServiceServerImpl) ListInstruments(ctx context.Context, req *rpc.ListInstrumentsRequest) (*rpc.ListInstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstruments not implemented")
}

func (s *djinServiceServerImpl) ListMusicScores(ctx context.Context, req *rpc.ListMusicScoresRequest) (*rpc.ListMusicScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMusicScores not implemented")
}

// Organizations

func (s *djinServiceServerImpl) GetOrganizationById(ctx context.Context, req *rpc.GetOrganizationByIdRequest) (*rpc.GetOrganizationByIdResponse, error) {
	organization, err := s.firestoreDB.GetOrganizationById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &rpc.GetOrganizationByIdResponse{
		Organization: organization,
	}, nil
}

// Projects

func (s *djinServiceServerImpl) ListProjects(ctx context.Context, req *rpc.ListProjectsRequest) (*rpc.ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
