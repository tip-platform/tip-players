#!/usr/bin/env bash
set -euo pipefail

# gRPC smoke test against local compose stack.
# Requires: grpcurl installed, docker compose available.

: "${APP_PORT:?APP_PORT is required}"

SUCCESS=false
for i in $(seq 1 30); do
  if grpcurl -plaintext "localhost:${APP_PORT}" grpc.health.v1.Health/Check >/dev/null 2>&1; then
    SUCCESS=true
    break
  fi
  sleep 2
done

if [ "$SUCCESS" = false ]; then
  echo "❌ Service did not become healthy in time"
  docker compose -f build/compose.yml logs || true
  exit 1
fi

grpcurl -plaintext "localhost:${APP_PORT}" grpc.health.v1.Health/Check
