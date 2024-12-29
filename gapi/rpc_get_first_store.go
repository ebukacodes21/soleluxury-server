package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetFirstStore(ctx context.Context, _ *emptypb.Empty) (*pb.GetStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get store")
	}

	store, err := s.repository.GetFirstStore(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get store")
	}

	resp := &pb.GetStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}
