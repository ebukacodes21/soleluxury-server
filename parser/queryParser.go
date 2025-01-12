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
	case *service.GetCategoriesRequest:
		return populateGetCategoriesParams(values, req)
	case *service.GetSizeRequest:
		return populateGetSizeParams(values, req)
	case *service.GetSizesRequest:
		return populateGetSizesParams(values, req)
	case *service.GetColorRequest:
		return populateGetColorParams(values, req)
	case *service.GetColorsRequest:
		return populateGetColorsParams(values, req)
	case *service.GetProductsRequest:
		return populateGetProductsParams(values, req)
	case *service.GetProductRequest:
		return populateGetProductParams(values, req)
	case *service.GetOrdersRequest:
		return populateGetOrdersParams(values, req)
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
		if parsedID, err := strconv.Atoi(Id); err == nil {
			r.Id = int64(parsedID)
		}
	}
	return nil
}

func populateGetCategoriesParams(values url.Values, r *service.GetCategoriesRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		if parsedStoreId, err := strconv.Atoi(storeId); err == nil {
			r.StoreId = int64(parsedStoreId)
		}
	}
	return nil
}

func populateGetSizeParams(values url.Values, r *service.GetSizeRequest) error {
	if Id := values.Get("id"); Id != "" {
		if parsedID, err := strconv.Atoi(Id); err == nil {
			r.Id = int64(parsedID)
		}
	}
	return nil
}

func populateGetSizesParams(values url.Values, r *service.GetSizesRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		if parsedStoreId, err := strconv.Atoi(storeId); err == nil {
			r.StoreId = int64(parsedStoreId)
		}
	}
	return nil
}

func populateGetColorParams(values url.Values, r *service.GetColorRequest) error {
	if Id := values.Get("id"); Id != "" {
		if parsedID, err := strconv.Atoi(Id); err == nil {
			r.Id = int64(parsedID)
		}
	}
	return nil
}

func populateGetColorsParams(values url.Values, r *service.GetColorsRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		if parsedStoreId, err := strconv.Atoi(storeId); err == nil {
			r.StoreId = int64(parsedStoreId)
		}
	}
	return nil
}

func populateGetProductsParams(values url.Values, r *service.GetProductsRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		if parsedStoreId, err := strconv.Atoi(storeId); err == nil {
			r.StoreId = int64(parsedStoreId)
		}
	}
	return nil
}

func populateGetProductParams(values url.Values, r *service.GetProductRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		if parsedStoreId, err := strconv.Atoi(storeId); err == nil {
			r.StoreId = int64(parsedStoreId)
		}
	}

	if pId := values.Get("product_id"); pId != "" {
		if parsedPId, err := strconv.Atoi(pId); err == nil {
			r.ProductId = int64(parsedPId)
		}
	}
	return nil
}

func populateGetOrdersParams(values url.Values, r *service.GetOrdersRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		if parsedStoreId, err := strconv.Atoi(storeId); err == nil {
			r.StoreId = int64(parsedStoreId)
		}
	}
	return nil
}
