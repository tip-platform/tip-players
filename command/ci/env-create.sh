#!/usr/bin/env bash
set -euo pipefail

# Creates a temporary .env in repo root for build/compose.yml (env_file: ../.env)

umask 077

cat > .env <<EOF
DB_USER=${DB_USER}
SQL_PASSWORD=${SQL_PASSWORD}
DB_PJ_NAME=${DB_PJ_NAME}
DB_PJ_PORT=${DB_PJ_PORT}
APP_PORT=${APP_PORT}
HEALTH_PORT=${HEALTH_PORT}
EOF

chmod 600 .env
