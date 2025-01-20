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
)

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	violations := validateCreateOrderRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	products, err := s.repository.CreateOrder(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to create order: %s", err)
	}

	totalValue := float64(0)
	for _, product := range products {
		totalValue += product.Price
	}

	paymentParams := map[string]interface{}{
		"email":  "georgeokafo1@gmail.com",
		"amount": totalValue,
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

	authURL := ""
	if data, ok := paymentResponse["data"].(map[string]interface{}); ok {
		if url, ok := data["authorization_url"].(string); ok {
			authURL = url
		}
	}
	res := &pb.CreateOrderResponse{
		Url: authURL,
	}
	return res, nil
}

func validateCreateOrderRequest(req *pb.CreateOrderRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	for _, item := range req.GetItems() {
		if err := validate.ValidateId(item); err != nil {
			violations = append(violations, fieldViolation("items", err))
		}
	}

	return violations
}
