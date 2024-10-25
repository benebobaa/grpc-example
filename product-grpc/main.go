package main

import (
	"log"
	"net/http"

	pb "simple-grpc-2/proto"

	"google.golang.org/grpc"
)

func main() {

	grpcConn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connect to client localhost:50051 -> ", err.Error())
	}
	defer grpcConn.Close()

	userClient := pb.NewAuthUserClient(grpcConn)
	handler := NewHandler(userClient)

	http.HandleFunc("/api/checkout-product", handler.CheckoutProduct)

	server := http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	log.Println("started server port 8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("err start server: ", err.Error())
	}
}
