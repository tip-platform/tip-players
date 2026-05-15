// Package main starts the ranktrack gRPC server.
package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/tip-platform/tip-players/internal/adapter/driver/di"
	"github.com/tip-platform/tip-players/internal/adapter/driver/rpc"
	pb "github.com/tip-platform/tip-players/proto"
)

func main() {
	if e := godotenv.Load(); e != nil {
		log.Printf(".env not loaded: %v", e)
	}

	container, e := di.NewContainer()

	if e != nil {
		log.Fatalf("Failed to create container: %v", e)
	}

	handler := rpc.NewPlayerHandler(container.PlayerService)

	server := grpc.NewServer()
	pb.RegisterPlayerServiceServer(server, handler)

	//nolint:gosec // Required for container networking
	lis, e := net.Listen("tcp", ":50051")

	if e != nil {
		log.Fatalf("Failed to listen: %v", e)
	}

	log.Println("gRPC server listening on :50051")

	if e := server.Serve(lis); e != nil {
		log.Fatalf("Failed to serve: %v", e)
	}
}
