#!/usr/bin/env bash
set -euo pipefail

docker compose -f build/compose.yml down -v >/dev/null 2>&1 || true
rm -f .env
