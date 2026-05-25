#!/usr/bin/env sh
set -eu

# Cross-platform (Linux/macOS) golangci-lint runner.
# - Installs golangci-lint if missing (via `go install`).
# - Uses repo config `.golangci.yml`.
#
# Usage:
#   ./command/lint-golangci.sh
#   ./command/lint-golangci.sh --fix

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)

GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION:-v1.64.6}

have_cmd() { command -v "$1" >/dev/null 2>&1; }

ensure_golangci_lint() {
  if have_cmd golangci-lint; then
    return 0
  fi

  echo "golangci-lint not found; installing ${GOLANGCI_LINT_VERSION} ..." >&2
  (cd "$ROOT_DIR" && go install "github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}")

  GOPATH_BIN=$(go env GOPATH)/bin
  if [ -x "${GOPATH_BIN}/golangci-lint" ]; then
    export PATH="${GOPATH_BIN}:${PATH}"
  fi

  if ! have_cmd golangci-lint; then
    echo "golangci-lint still not found after install. Ensure GOPATH/bin is in PATH." >&2
    exit 1
  fi
}

ensure_golangci_lint

cd "$ROOT_DIR"

golangci-lint run --config .golangci.yml --timeout 5m "$@"
