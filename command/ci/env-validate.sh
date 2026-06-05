#!/usr/bin/env bash
set -euo pipefail

require() {
  local name="$1"
  if [ -z "${!name:-}" ]; then
    echo "Missing required env: ${name}" >&2
    exit 1
  fi
}

require APP_PORT
require HEALTH_PORT
require DB_PJ_PORT
require DB_USER
require DB_PJ_NAME

if [ -z "${SQL_PASSWORD:-}" ]; then
  echo "Missing required secret env: SQL_PASSWORD" >&2
  exit 1
fi
