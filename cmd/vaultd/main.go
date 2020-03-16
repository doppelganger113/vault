package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	vault "vault/internals"
	"vault/internals/pb"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listener address")
		grpcAddr = flag.String("grpc", ":8081", "gRPC listener address")
	)
	flag.Parse()
	service := vault.NewService()
	errChan := make(chan error)

	// Look for interrupt signal
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	hashEndpoint := vault.MakeHashEndpoint(service)
	validateEndpoint := vault.MakeValidateEndpoint(service)
	endpoints := vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}

	// HTTP transport
	go func() {
		log.Println("http: ", *httpAddr)
		handler := vault.NewHttpServer(endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// gRPC transport
	go func() {
		listener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("gRPC:", *grpcAddr)
		handler := vault.NewGRPCServer(endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	log.Fatalln(<-errChan)
}
