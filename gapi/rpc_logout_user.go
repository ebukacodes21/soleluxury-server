package gapi

import (
	"context"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) LogoutUser(ctx context.Context, _ *emptypb.Empty) (*pb.LogoutResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	args := db.LogoutParams{
		UserID:   payload.UserId,
		Username: payload.Username,
	}

	err = s.repository.Logout(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to log user out %s ", err)
	}

	message := "logout successful"
	resp := &pb.LogoutResponse{
		Message: message,
	}
	return resp, nil
}
