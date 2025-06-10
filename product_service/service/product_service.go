package service

import (
	"context"
	"log"
	"product_service/proto_impl/proto_product_impl"
)

type ProductServiceServer struct {
	proto_product_impl.UnimplementedProductServiceServer
}

func NewProductService() *ProductServiceServer {
	return &ProductServiceServer{}
}

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *proto_product_impl.ProductRequest) (*proto_product_impl.ProductResponse, error) {

	log.Println("Hasil req : ", req)
	// save to db using repo

	return &proto_product_impl.ProductResponse{}, nil
}
