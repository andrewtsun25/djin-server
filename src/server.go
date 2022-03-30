package main

import (
	"context"
	rpcv1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/rpc/v1"
	servicev1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type djinServiceServer struct {
	servicev1.UnimplementedDjinServiceServer
	firestoreDB *FirestoreDB
}

func NewDjinServiceServer(ctx context.Context, serviceAccount option.ClientOption) *djinServiceServer {
	fdb := NewFirestoreDB(ctx, serviceAccount)
	if fdb == nil {
		return nil
	}
	return &djinServiceServer{
		firestoreDB: fdb,
	}
}

func (s *djinServiceServer) Close() error {
	return s.firestoreDB.Close()
}

func (s *djinServiceServer) GetOrganizationById(ctx context.Context, req *rpcv1.GetOrganizationByIdRequest) (*rpcv1.GetOrganizationByIdResponse, error) {
	organization, err := s.firestoreDB.GetOrganizationById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &rpcv1.GetOrganizationByIdResponse{
		Organization: organization,
	}, nil
}

func (s *djinServiceServer) ListEducationsByType(ctx context.Context, req *rpcv1.ListEducationsByTypeRequest) (*rpcv1.ListEducationsByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEducationsByType not implemented")
}
func (s *djinServiceServer) ListEmployments(ctx context.Context, req *rpcv1.ListEmploymentsRequest) (*rpcv1.ListEmploymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEmployments not implemented")
}
func (s *djinServiceServer) ListHbvResearchPapers(ctx context.Context, req *rpcv1.ListHbvResearchPapersRequest) (*rpcv1.ListHbvResearchPapersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHbvResearchPapers not implemented")
}
func (s *djinServiceServer) ListHolisticOfficeLinks(ctx context.Context, req *rpcv1.ListHolisticOfficeLinksRequest) (*rpcv1.ListHolisticOfficeLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHolisticOfficeLinks not implemented")
}
func (s *djinServiceServer) ListHolisticOfficeModules(ctx context.Context, req *rpcv1.ListHolisticOfficeModulesRequest) (*rpcv1.ListHolisticOfficeModulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHolisticOfficeModules not implemented")
}
func (s *djinServiceServer) GetMartialArtsStyleById(ctx context.Context, req *rpcv1.GetMartialArtsStyleByIdRequest) (*rpcv1.GetMartialArtsStyleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMartialArtsStyleById not implemented")
}
func (s *djinServiceServer) ListInstruments(ctx context.Context, req *rpcv1.ListInstrumentsRequest) (*rpcv1.ListInstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstruments not implemented")
}
func (s *djinServiceServer) ListMusicScores(ctx context.Context, req *rpcv1.ListMusicScoresRequest) (*rpcv1.ListMusicScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMusicScores not implemented")
}
func (s *djinServiceServer) ListProjects(ctx context.Context, req *rpcv1.ListProjectsRequest) (*rpcv1.ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
