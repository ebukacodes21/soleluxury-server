package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	violations := validateCreateOrderRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	order, products, err := s.repository.CreateOrder(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to create order: %s", err)
	}

	totalValue := float64(0)
	for _, product := range products {
		totalValue += product.Price
	}

	paymentParams := map[string]interface{}{
		"email":  req.GetEmail(),
		"amount": totalValue * 100,
	}
	paymentBody, _ := json.Marshal(paymentParams)
	paymentURL := "https://api.paystack.co/transaction/initialize"

	httpReq, err := http.NewRequest("POST", paymentURL, bytes.NewBuffer(paymentBody))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create payment request: %s", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+s.config.PAYSTACK_SECRET_KEY)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send payment request: %s", err)
	}
	defer resp.Body.Close()

	var paymentResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse payment response: %s", err)
	}

	reference := ""
	if data, ok := paymentResponse["data"].(map[string]interface{}); ok {
		if r, ok := data["reference"].(string); ok {
			reference = r
		}
	}
	res := &pb.CreateOrderResponse{
		Reference: reference,
		OrderId:   order.ID.Hex(),
	}
	return res, nil
}

func (s *Server) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.GetOrdersResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get orders")
	}

	orders, err := s.repository.GetOrders(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get orders: %s", err)
	}

	resp := &pb.GetOrdersResponse{
		Orders: convertOrders(orders),
	}

	return resp, nil
}

func (s *Server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	violations := validateUpdateOrderRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	message, err := s.repository.UpdateOrder(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to update order: %s", err)
	}

	resp := &pb.UpdateOrderResponse{
		Message: message,
	}

	return resp, nil
}

func validateCreateOrderRequest(req *pb.CreateOrderRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	for _, item := range req.GetItems() {
		if err := validate.ValidateId(item); err != nil {
			violations = append(violations, fieldViolation("items", err))
		}
	}

	if err := validate.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := validate.ValidateAddress(req.GetAddress()); err != nil {
		violations = append(violations, fieldViolation("address", err))
	}

	if err := validate.ValidatePhone(req.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}

	return violations
}

func validateUpdateOrderRequest(req *pb.UpdateOrderRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetOrderId()); err != nil {
		violations = append(violations, fieldViolation("order_id", err))
	}

	if err := validate.ValidateOrderMessage(req.GetMessage()); err != nil {
		violations = append(violations, fieldViolation("message", err))
	}

	if err := validate.ValidateOrderStatus(req.GetStatus()); err != nil {
		violations = append(violations, fieldViolation("status", err))
	}
	return violations
}
