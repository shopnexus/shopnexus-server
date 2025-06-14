package product

import (
	"context"
	"net/http"
	"shopnexus-go-service/internal/model"
	"shopnexus-go-service/internal/service/product"

	common_grpc "shopnexus-go-service/internal/transport/connect/handler/common"

	"connectrpc.com/connect"
	productv1 "github.com/shopnexus/shopnexus-protobuf-gen-go/pb/product/v1"
	"github.com/shopnexus/shopnexus-protobuf-gen-go/pb/product/v1/productv1connect"
)

var _ productv1connect.ProductServiceHandler = (*ImplementedProductServiceHandler)(nil)

type ImplementedProductServiceHandler struct {
	productv1connect.UnimplementedProductServiceHandler
	service product.Service
}

func NewProductServiceHandler(service product.Service, opts ...connect.HandlerOption) (string, http.Handler) {
	return productv1connect.NewProductServiceHandler(&ImplementedProductServiceHandler{service: service}, opts...)
}

func (s *ImplementedProductServiceHandler) GetProduct(ctx context.Context, req *connect.Request[productv1.GetProductRequest]) (*connect.Response[productv1.GetProductResponse], error) {
	data, err := s.service.GetProduct(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productv1.GetProductResponse{
		Data: modelToProductEntity(data),
	}), nil
}

func (s *ImplementedProductServiceHandler) ListProducts(ctx context.Context, req *connect.Request[productv1.ListProductsRequest]) (*connect.Response[productv1.ListProductsResponse], error) {
	data, err := s.service.ListProducts(ctx, product.ListProductsParams{
		PaginationParams: model.PaginationParams{
			Page:  req.Msg.GetPagination().GetPage(),
			Limit: req.Msg.GetPagination().GetLimit(),
		},
		ProductModelID:  req.Msg.ProductModelId,
		QuantityFrom:    req.Msg.QuantityFrom,
		QuantityTo:      req.Msg.QuantityTo,
		SoldFrom:        req.Msg.SoldFrom,
		SoldTo:          req.Msg.SoldTo,
		AddPriceFrom:    req.Msg.AddPriceFrom,
		AddPriceTo:      req.Msg.AddPriceTo,
		IsActive:        req.Msg.IsActive,
		Metadata:        req.Msg.Metadata,
		DateCreatedFrom: req.Msg.DateCreatedFrom,
		DateCreatedTo:   req.Msg.DateCreatedTo,
	})
	if err != nil {
		return nil, err
	}

	var products []*productv1.ProductEntity
	for _, d := range data.Data {
		products = append(products, modelToProductEntity(d))
	}

	return connect.NewResponse(&productv1.ListProductsResponse{
		Data:       products,
		Pagination: common_grpc.ToProtoPaginationResponse(data),
	}), nil
}

func (s *ImplementedProductServiceHandler) CreateProduct(ctx context.Context, req *connect.Request[productv1.CreateProductRequest]) (*connect.Response[productv1.CreateProductResponse], error) {
	data, err := s.service.CreateProduct(ctx, model.Product{
		ProductModelID: req.Msg.ProductModelId,
		Quantity:       req.Msg.Quantity,
		AddPrice:       req.Msg.AddPrice,
		IsActive:       req.Msg.IsActive,
		Metadata:       req.Msg.Metadata,
		Resources:      req.Msg.Resources,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productv1.CreateProductResponse{
		Data: modelToProductEntity(data),
	}), nil
}

func (s *ImplementedProductServiceHandler) UpdateProduct(ctx context.Context, req *connect.Request[productv1.UpdateProductRequest]) (*connect.Response[productv1.UpdateProductResponse], error) {
	var resources *[]string
	if req.Msg.Resources != nil {
		resources = &req.Msg.Resources
	}

	var metadata *[]byte
	if req.Msg.Metadata != nil {
		metadata = &req.Msg.Metadata
	}

	err := s.service.UpdateProduct(ctx, product.UpdateProductParams{
		ID:             req.Msg.GetId(),
		ProductModelID: req.Msg.ProductModelId,
		Quantity:       req.Msg.Quantity,
		Sold:           req.Msg.Sold,
		AddPrice:       req.Msg.AddPrice,
		IsActive:       req.Msg.IsActive,
		CanCombine:     req.Msg.CanCombine,
		Metadata:       metadata,
		Resources:      resources,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productv1.UpdateProductResponse{}), nil
}

func (s *ImplementedProductServiceHandler) DeleteProduct(ctx context.Context, req *connect.Request[productv1.DeleteProductRequest]) (*connect.Response[productv1.DeleteProductResponse], error) {
	err := s.service.DeleteProduct(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productv1.DeleteProductResponse{}), nil
}

func (s *ImplementedProductServiceHandler) GetProductByPOPID(ctx context.Context, req *connect.Request[productv1.GetProductByPOPIDRequest]) (*connect.Response[productv1.GetProductByPOPIDResponse], error) {
	data, err := s.service.GetProductByPOPID(ctx, req.Msg.ProductOnPaymentId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&productv1.GetProductByPOPIDResponse{
		Data: modelToProductEntity(data),
	}), nil
}

func modelToProductEntity(data model.Product) *productv1.ProductEntity {
	return &productv1.ProductEntity{
		Id:             data.ID,
		ProductModelId: data.ProductModelID,
		Quantity:       data.Quantity,
		Sold:           data.Sold,
		AddPrice:       data.AddPrice,
		IsActive:       data.IsActive,
		CanCombine:     data.CanCombine,
		Metadata:       data.Metadata,
		DateCreated:    data.DateCreated,
		DateUpdated:    data.DateUpdated,
		Resources:      data.Resources,
	}
}
