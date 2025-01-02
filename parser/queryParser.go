package parser

import (
	"net/url"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/proto"

	service "github.com/ebukacodes21/soleluxury-server/pb"
)

type CustomQueryParameterParser struct{}

func (p *CustomQueryParameterParser) Parse(target proto.Message, values url.Values, filter *utilities.DoubleArray) error {
	switch req := target.(type) {
	case *service.GetBillboardsRequest:
		return populateGetBillboardsParams(values, req)
	case *service.GetCategoryRequest:
		return populateGetCategoryParams(values, req)
	}

	return (*runtime.DefaultQueryParser)(nil).Parse(target, values, filter)
}

func populateGetBillboardsParams(values url.Values, r *service.GetBillboardsRequest) error {
	if storeID := values.Get("store_id"); storeID != "" {
		if parsedStoreID, err := strconv.Atoi(storeID); err == nil {
			r.StoreId = int64(parsedStoreID)
		}
	}
	return nil
}

func populateGetCategoryParams(values url.Values, r *service.GetCategoryRequest) error {
	if Id := values.Get("id"); Id != "" {
		if parsedStoreID, err := strconv.Atoi(Id); err == nil {
			r.Id = int64(parsedStoreID)
		}
	}
	return nil
}
