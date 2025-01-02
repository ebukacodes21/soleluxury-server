package gapi

import (
	"context"
	"time"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"github.com/ebukacodes21/soleluxury-server/worker"
	"github.com/hibiken/asynq"
	pg "github.com/lib/pq"
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

	hash, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to hash password %s", err)
	}
	code := validate.RandomString(32)

	args := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:         req.GetUsername(),
			Email:            req.GetEmail(),
			Password:         hash,
			VerificationCode: code,
		},
		AfterCreate: func(user db.User) error {
			payload := worker.RegisterMailPayload{
				Username: user.Username,
				Email:    user.Email,
			}
			options := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.Critical),
			}
			return s.td.DistributeTaskRegisterMail(ctx, &payload, options...)
		},
	}

	result, err := s.repository.CreateUserTx(ctx, args)
	if err != nil {
		if pgErr, ok := err.(*pg.Error); ok {
			switch pgErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.Internal, "unique violation %s ", pgErr)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user -- %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(result.User),
	}

	return resp, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginUserRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	user, err := s.repository.GetUser(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no user found %s ", err)
	}

	err = utils.ComparePassword(req.GetPassword(), user.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password %s ", err)
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
	session, err := s.repository.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.Id,
		Username:     user.Username,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    metaData.UserAgent,
		ClientIp:     metaData.ClientIp,
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session")
	}

	resp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		SessionId:             session.ID.String(),
	}

	return resp, nil
}

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
