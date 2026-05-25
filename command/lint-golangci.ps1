$ErrorActionPreference = 'Stop'

<#
Cross-platform (Windows/macOS/Linux with PowerShell) golangci-lint runner.
- Installs golangci-lint if missing (via `go install`).
- Uses repo config `.golangci.yml`.

Usage:
  pwsh ./command/lint-golangci.ps1
  pwsh ./command/lint-golangci.ps1 --fix
#>

$RootDir = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path

if (-not $env:GOLANGCI_LINT_VERSION) {
  $env:GOLANGCI_LINT_VERSION = 'v1.64.6'
}

function Ensure-GolangciLint {
  $cmd = Get-Command golangci-lint -ErrorAction SilentlyContinue
  if ($cmd) { return }

  Write-Host "golangci-lint not found; installing $($env:GOLANGCI_LINT_VERSION) ..." -ForegroundColor Yellow

  Push-Location $RootDir
  try {
    go install "github.com/golangci/golangci-lint/cmd/golangci-lint@$($env:GOLANGCI_LINT_VERSION)"
  } finally {
    Pop-Location
  }

  $gopath = (go env GOPATH)
  $bin = Join-Path $gopath 'bin'
  $exe = Join-Path $bin 'golangci-lint.exe'
  $nix = Join-Path $bin 'golangci-lint'

  if (Test-Path $exe) {
    $env:PATH = "$bin;$env:PATH"
    return
  }
  if (Test-Path $nix) {
    $env:PATH = "$bin:$env:PATH"
    return
  }

  throw "golangci-lint still not found after install. Ensure GOPATH/bin is in PATH."
}

Ensure-GolangciLint

Push-Location $RootDir
try {
  golangci-lint run --config .golangci.yml --timeout 5m @args
} finally {
  Pop-Location
}
