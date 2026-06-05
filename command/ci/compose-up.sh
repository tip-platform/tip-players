#!/usr/bin/env bash
set -euo pipefail

docker compose -f build/compose.yml up -d --build
