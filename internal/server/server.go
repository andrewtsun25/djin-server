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

func (s *djinServiceServerImpl) GetOrganizationById(ctx context.Context, req *rpc.GetOrganizationByIdRequest) (*rpc.GetOrganizationByIdResponse, error) {
	organization, err := s.firestoreDB.GetOrganizationById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &rpc.GetOrganizationByIdResponse{
		Organization: organization,
	}, nil
}

func (s *djinServiceServerImpl) ListEducationsByType(ctx context.Context, req *rpc.ListEducationsByTypeRequest) (*rpc.ListEducationsByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEducationsByType not implemented")
}
func (s *djinServiceServerImpl) ListEmployments(ctx context.Context, req *rpc.ListEmploymentsRequest) (*rpc.ListEmploymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEmployments not implemented")
}
func (s *djinServiceServerImpl) ListHbvResearchPapers(ctx context.Context, req *rpc.ListHbvResearchPapersRequest) (*rpc.ListHbvResearchPapersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHbvResearchPapers not implemented")
}
func (s *djinServiceServerImpl) ListHolisticOfficeLinks(ctx context.Context, req *rpc.ListHolisticOfficeLinksRequest) (*rpc.ListHolisticOfficeLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHolisticOfficeLinks not implemented")
}
func (s *djinServiceServerImpl) ListHolisticOfficeModules(ctx context.Context, req *rpc.ListHolisticOfficeModulesRequest) (*rpc.ListHolisticOfficeModulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHolisticOfficeModules not implemented")
}
func (s *djinServiceServerImpl) GetMartialArtsStyleById(ctx context.Context, req *rpc.GetMartialArtsStyleByIdRequest) (*rpc.GetMartialArtsStyleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMartialArtsStyleById not implemented")
}
func (s *djinServiceServerImpl) ListInstruments(ctx context.Context, req *rpc.ListInstrumentsRequest) (*rpc.ListInstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstruments not implemented")
}
func (s *djinServiceServerImpl) ListMusicScores(ctx context.Context, req *rpc.ListMusicScoresRequest) (*rpc.ListMusicScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMusicScores not implemented")
}
func (s *djinServiceServerImpl) ListProjects(ctx context.Context, req *rpc.ListProjectsRequest) (*rpc.ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
