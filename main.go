package main

import (
	"log"
	"net"

	"github.com/0x41gawor/grpc-demo/invoicer"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot start a listner: %s", err)
	}
	grpcServer := grpc.NewServer()
	service := &MyInvoicerServer{}

	invoicer.RegisterInvoicerServer(grpcServer, service)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}
