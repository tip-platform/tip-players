#!/usr/bin/env sh
set -eu

# Cross-platform (Linux/macOS) semantic-release runner.
# - Reads GITHUB_TOKEN from environment or config/.env.
# - Runs in dry-run mode by default; pass --no-dry-run to publish.
#
# Usage:
#   ./command/semantic-release.sh
#   ./command/semantic-release.sh --no-dry-run

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)

ENV_FILE="${ROOT_DIR}/config/.env"

# ---------------------------------------------------------------------------
# Load .env if GITHUB_TOKEN is not already set
# ---------------------------------------------------------------------------
if [ -z "${GITHUB_TOKEN:-}" ]; then
  if [ -f "$ENV_FILE" ]; then
    # Export only GITHUB_TOKEN from the .env file, ignore comments and blanks
    GITHUB_TOKEN=$(grep -E '^GITHUB_TOKEN=' "$ENV_FILE" | head -1 | cut -d'=' -f2-)
    export GITHUB_TOKEN
  fi
fi

if [ -z "${GITHUB_TOKEN:-}" ]; then
  echo "error: GITHUB_TOKEN is not set." >&2
  echo "  Set it in your environment or add GITHUB_TOKEN=<token> to config/.env" >&2
  exit 1
fi

# ---------------------------------------------------------------------------
# Resolve dry-run flag (default: dry-run ON)
# ---------------------------------------------------------------------------
DRY_RUN="--dry-run"
for arg in "$@"; do
  if [ "$arg" = "--no-dry-run" ]; then
    DRY_RUN=""
  fi
done

# Remove --no-dry-run from args forwarded to semantic-release
EXTRA_ARGS=""
for arg in "$@"; do
  if [ "$arg" != "--no-dry-run" ]; then
    EXTRA_ARGS="$EXTRA_ARGS $arg"
  fi
done

# ---------------------------------------------------------------------------
# Run
# ---------------------------------------------------------------------------
cd "$ROOT_DIR"

echo "Running semantic-release${DRY_RUN:+ (dry-run)}..." >&2

# shellcheck disable=SC2086
npx \
  -p semantic-release \
  -p @semantic-release/changelog \
  -p @semantic-release/git \
  -p @semantic-release/github \
  semantic-release \
  --branches develop \
  $DRY_RUN \
  $EXTRA_ARGS
