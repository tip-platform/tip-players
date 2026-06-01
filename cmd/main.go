// Package main starts the ranktrack gRPC server.
package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/tip-platform/tip-players/api/proto"
	"github.com/tip-platform/tip-players/internal/infra/driver/config"
	"github.com/tip-platform/tip-players/internal/infra/driver/di"
	"github.com/tip-platform/tip-players/internal/infra/driver/health"
	"github.com/tip-platform/tip-players/internal/infra/driver/rpc"
)

func main() {
	cfg, e := config.Load()

	if e != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", e)
		os.Exit(1)
	}

	container, e := di.NewContainer()

	if e != nil {
		fmt.Fprintf(os.Stderr, "Failed to create container: %v\n", e)
		os.Exit(1)
	}

	if container.PlayerService == nil {
		fmt.Fprintln(os.Stderr, "Failed to create container: PlayerService is nil")
		os.Exit(1)
	}

	handler := rpc.NewPlayerHandler(container.PlayerService)

	server := grpc.NewServer()
	if cfg.GRPCReflectionEnable {
		reflection.Register(server)
	}

	// Health services live in adapters; main only wires them.
	_ = health.RegisterGRPC(server, func(ctx context.Context) error {
		if container.PlayerStore == nil {
			return nil
		}

		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

		defer cancel()

		return container.PlayerStore.Ping(ctx)
	})

	pb.RegisterPlayerServiceServer(server, handler)

	grpcAddr := cfg.GRPCAddr

	//nolint:gosec // Required for container networking
	lis, e := net.Listen("tcp", grpcAddr)

	if e != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", e)
		os.Exit(1)
	}

	// Optional HTTP health endpoints.
	healthAddr := cfg.HealthAddr
	if cfg.HealthEnable {
		_ = health.StartHTTP(healthAddr, func(ctx context.Context) error {
			if container.PlayerStore == nil {
				return nil
			}
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

			defer cancel()

			return container.PlayerStore.Ping(ctx)
		})
	}

	fmt.Printf("gRPC server listening on %s\n", grpcAddr)

	if cfg.HealthEnable {
		fmt.Printf("HTTP health listening on %s (GET /healthz, /readyz)\n", healthAddr)
	}

	if e := server.Serve(lis); e != nil {
		fmt.Fprintf(os.Stderr, "Failed to serve: %v\n", e)
		os.Exit(1)
	}
}
