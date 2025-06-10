package main

import (
	"log"
	"net"
	"product_service/proto_impl/proto_product_impl"
	"product_service/service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// repo := NewInMemoryRepo() // bisa ganti ke database repo
	proto_product_impl.RegisterProductServiceServer(grpcServer, service.NewProductService())

	log.Println("ProductService running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
