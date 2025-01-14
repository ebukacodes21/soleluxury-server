package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/db"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	user, err := s.repository.CreateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user  %s ", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginUserRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	user, err := s.repository.FindUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user not found  %s ", err)
	}

	// if !user.IsVerified {
	// 	return nil, status.Errorf(codes.Unauthenticated, "user is yet to be verified")
	// }

	accessToken, accessPayload, err := s.token.CreateToken(user.Username, user.ID, user.Role, s.config.TokenAccess)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create access token")
	}

	refreshToken, refreshPayload, err := s.token.CreateToken(user.Username, user.ID, user.Role, s.config.RefreshAccess)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create refresh token")
	}

	metaData := s.extractMetaData(ctx)
	args := db.SessionReq{
		Username:     user.Username,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ClientIp:     metaData.ClientIp,
		UserAgent:    metaData.UserAgent,
		IsBlocked:    false,
		ExpiredAt:    &refreshPayload.ExpiredAt,
	}

	session, err := s.repository.CreateSession(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to create session. %s ", err)
	}

	resp := &pb.LoginUserResponse{
		SessionId:             session.ID.String(),
		User:                  convertUser(user),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return resp, nil
}

func (s *Server) LogoutUser(ctx context.Context, req *emptypb.Empty) (*pb.LogoutResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	err = s.repository.LogOut(ctx, payload.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to logout user %s ", err)
	}

	resp := &pb.LogoutResponse{
		Message: "logout successful",
	}

	return resp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := validate.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := validate.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}

func validateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := validate.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
