package parser

import (
	"net/url"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/proto"

	service "github.com/ebukacodes21/soleluxury-server/pb"
)

type CustomQueryParameterParser struct{}

func (p *CustomQueryParameterParser) Parse(target proto.Message, values url.Values, filter *utilities.DoubleArray) error {
	switch req := target.(type) {
	case *service.DeleteStoreRequest:
		return populateDeleteStoreParams(values, req)
	case *service.GetBillboardsRequest:
		return populateGetBillboardsParams(values, req)
	case *service.DeleteBillboardRequest:
		return populateDeleteBillboardParams(values, req)
	case *service.GetCategoryRequest:
		return populateGetCategoryParams(values, req)
	case *service.GetCategoriesRequest:
		return populateGetCategoriesParams(values, req)
	case *service.DeleteCategoryRequest:
		return populateDeleteCategoryParams(values, req)
	case *service.GetSizeRequest:
		return populateGetSizeParams(values, req)
	case *service.GetSizesRequest:
		return populateGetSizesParams(values, req)
	case *service.DeleteSizeRequest:
		return populateDeleteSizeParams(values, req)
	case *service.GetColorRequest:
		return populateGetColorParams(values, req)
	case *service.GetColorsRequest:
		return populateGetColorsParams(values, req)
	case *service.DeleteColorRequest:
		return populateDeleteColorParams(values, req)
	case *service.DeleteProductRequest:
		return populateDeleteProductParams(values, req)
	case *service.GetProductsRequest:
		return populateGetProductsParams(values, req)
	case *service.GetCategoryProductsRequest:
		return populateGetCategoryProductsParams(values, req)
	case *service.GetProductRequest:
		return populateGetProductParams(values, req)
	}

	return (*runtime.DefaultQueryParser)(nil).Parse(target, values, filter)
}

func populateDeleteStoreParams(values url.Values, r *service.DeleteStoreRequest) error {
	if storeID := values.Get("id"); storeID != "" {
		r.Id = storeID
	}
	return nil
}

func populateDeleteBillboardParams(values url.Values, r *service.DeleteBillboardRequest) error {
	if id := values.Get("id"); id != "" {
		r.Id = id
	}
	return nil
}

func populateDeleteCategoryParams(values url.Values, r *service.DeleteCategoryRequest) error {
	if id := values.Get("id"); id != "" {
		r.Id = id
	}
	return nil
}

func populateDeleteSizeParams(values url.Values, r *service.DeleteSizeRequest) error {
	if id := values.Get("id"); id != "" {
		r.Id = id
	}
	return nil
}

func populateDeleteColorParams(values url.Values, r *service.DeleteColorRequest) error {
	if id := values.Get("id"); id != "" {
		r.Id = id
	}
	return nil
}

func populateDeleteProductParams(values url.Values, r *service.DeleteProductRequest) error {
	if id := values.Get("product_id"); id != "" {
		r.ProductId = id
	}
	return nil
}

func populateGetBillboardsParams(values url.Values, r *service.GetBillboardsRequest) error {
	if storeID := values.Get("store_id"); storeID != "" {
		r.StoreId = storeID
	}
	return nil
}

func populateGetCategoryParams(values url.Values, r *service.GetCategoryRequest) error {
	if Id := values.Get("id"); Id != "" {
		r.Id = Id
	}
	return nil
}

func populateGetCategoriesParams(values url.Values, r *service.GetCategoriesRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		r.StoreId = storeId
	}
	return nil
}

func populateGetSizeParams(values url.Values, r *service.GetSizeRequest) error {
	if Id := values.Get("id"); Id != "" {
		r.Id = Id
	}
	return nil
}

func populateGetSizesParams(values url.Values, r *service.GetSizesRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		r.StoreId = storeId
	}
	return nil
}

func populateGetColorParams(values url.Values, r *service.GetColorRequest) error {
	if Id := values.Get("id"); Id != "" {
		r.Id = Id
	}
	return nil
}

func populateGetColorsParams(values url.Values, r *service.GetColorsRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		r.StoreId = storeId
	}
	return nil
}

func populateGetProductsParams(values url.Values, r *service.GetProductsRequest) error {
	if storeId := values.Get("store_id"); storeId != "" {
		r.StoreId = storeId
	}
	return nil
}

func populateGetCategoryProductsParams(values url.Values, r *service.GetCategoryProductsRequest) error {
	if categoryId := values.Get("category_id"); categoryId != "" {
		r.CategoryId = categoryId
	}
	return nil
}

func populateGetProductParams(values url.Values, r *service.GetProductRequest) error {
	if pId := values.Get("product_id"); pId != "" {
		r.ProductId = pId
	}
	return nil
}
