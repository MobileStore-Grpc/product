package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/MobileStore-Grpc/product/pb"
	"github.com/MobileStore-Grpc/product/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func runRESTServer(mobileService pb.MobileServiceServer, listener net.Listener, grpcEndpoint string) error {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// in-process handler
	err := pb.RegisterMobileServiceHandlerServer(ctx, mux, mobileService)
	if err != nil {
		return err
	}
	log.Printf("Start REST server at %s", listener.Addr().String())
	return http.Serve(listener, mux)
}

func runGRPCServer(mobileService pb.MobileServiceServer, listener net.Listener) error {
	grpcServer := grpc.NewServer()
	//Like we register router with server, server:= http.Server{Addr: :8080, Handler: router}, simimarly we have to register LaptopServer object with grpcServer using RegisterLaptopServiceServer() function
	//Register our service implementation with the gRPC server
	pb.RegisterMobileServiceServer(grpcServer, mobileService)

	// Like we run server.ListenAnsServer(), similarly we do  grpcServer.serve()
	log.Printf("Start GRPC server at %s", listener.Addr().String())

	return grpcServer.Serve(listener)
}

func main() {
	port := flag.Int("port", 0, "the server port")
	serverType := flag.String("type", "grpc", "type of server (grpc/rest)")
	endPoint := flag.String("endpoint", "", "gRPC endpoint")
	flag.Parse()

	mobileStore := service.NewInMemoryMobileStore()

	mobileService := service.NewMobileService(mobileStore)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if *serverType == "grpc" {
		err = runGRPCServer(mobileService, listener)
	} else {
		err = runRESTServer(mobileService, listener, *endPoint)
	}
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
