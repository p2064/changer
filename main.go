package main

import (
	"log"
	"net"

	"github.com/p2064/changer/handlers"
	"github.com/p2064/changer/proto"
	"github.com/p2064/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	log.Print("Start changer")
	if config.Status != config.GOOD {
		log.Fatal("failed to get config")
	}
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterChangerServiceServer(grpcServer, &handlers.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
