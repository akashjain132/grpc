package main

import (
	"context"
	"fmt"
	"grpcClient/amazon"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Client")

	conn, err := grpc.Dial("0.0.0.0:60001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial client %v", err)
	}

	c := amazon.NewProductServiceClient(conn)

	doUnary(c)
}

func doUnary(c amazon.ProductServiceClient) {

	req := amazon.ProductID{
		Id: "49",
	}

	res, err := c.GetProduct(context.Background(), &req)
	if err != nil {
		log.Fatalf("Unable to fetch product detail %v", err)
	}

	log.Printf("Product details: %v", res)
}
