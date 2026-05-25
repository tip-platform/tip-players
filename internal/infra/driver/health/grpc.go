// Package health provides gRPC health reporting.
package health

import (
	"context"
	"time"

	"google.golang.org/grpc"
	grpc_health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// RegisterGRPC registers the standard gRPC health service.
// If ready is provided, it periodically updates the serving status.
func RegisterGRPC(s *grpc.Server, ready ReadyFunc) *grpc_health.Server {
	hs := grpc_health.NewServer()
	healthpb.RegisterHealthServer(s, hs)

	// Default to SERVING for liveness: process is up.
	hs.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	if ready == nil {
		return hs
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			err := ready(ctx)
			cancel()

			if err != nil {
				hs.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
			} else {
				hs.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
			}
		}
	}()

	return hs
}
