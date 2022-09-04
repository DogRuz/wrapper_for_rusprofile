package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"wrapper_for_rusprofile/internal/api"
	"wrapper_for_rusprofile/internal/controller"
)

func main() {
	s := grpc.NewServer()
	srv := &controller.InnHandler{}
	api.RegisterGetInformationServer(s, srv)
	l, err := net.Listen("tcp", ":8070")
	log.Println("Serving gRPC on 0.0.0.0:8070")
	go func() {
		log.Fatalln(s.Serve(l))
	}()
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8070",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = api.RegisterGetInformationHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/inn/", gwmux)
	mux.HandleFunc("/swagger.json", controller.SwaggerHandler{}.Read)

	log.Println("Serving gRPC-Gateway on http://localhost:8090")
	log.Fatalln(http.ListenAndServe(":8090", mux))
}
