package main

import (
	"context"
	entityv1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/entity/v1"
	rpcv1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/rpc/v1"
	servicev1 "go.buf.build/grpc/go/andrewtsun25/djin/proto/dev/djin/service/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type djinServiceServer struct {
	servicev1.UnimplementedDjinServiceServer
}

func (s *djinServiceServer) GetOrganizationById(ctx context.Context, req *rpcv1.GetOrganizationByIdRequest) (*rpcv1.GetOrganizationByIdResponse, error) {
	return &rpcv1.GetOrganizationByIdResponse{
		Organization: &entityv1.Organization{
			Id:      "organizationId",
			Name:    "Organization Name",
			LogoUrl: wrapperspb.String("https://upload.wikimedia.org/wikipedia/en/a/a9/Example.jpg"),
		},
	}, nil
}
