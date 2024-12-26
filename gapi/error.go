package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func invalidArgs(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequests := &errdetails.BadRequest{FieldViolations: violations}
	invalidStatus := status.New(codes.InvalidArgument, "invalid request parameters")

	status, err := invalidStatus.WithDetails(badRequests)
	if err != nil {
		return invalidStatus.Err()
	}

	return status.Err()
}
