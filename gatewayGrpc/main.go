package main

import (
	"context"
	"gatewayGrpc/proto_impl/proto_product_impl"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// client := productpb.NewProductServiceClient(conn)
	client := proto_product_impl.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call AddProduct
	res, err := client.CreateProduct(ctx, &proto_product_impl.ProductRequest{
		Product: "Laptop",
		Email:   "mike@gmail.coms",
	})
	if err != nil {
		log.Fatalf("Error calling AddProduct: %v", err)
	}
	log.Printf("AddProduct Response: %+v", res)
}
