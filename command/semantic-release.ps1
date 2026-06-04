#Requires -Version 5.1
<#
.SYNOPSIS
  Cross-platform (Windows) semantic-release runner.

.DESCRIPTION
  - Reads GITHUB_TOKEN from environment or config/.env.
  - Runs in dry-run mode by default; pass -NoDryRun to publish.

.PARAMETER NoDryRun
  Disables dry-run mode and publishes the release.

.EXAMPLE
  pwsh ./command/semantic-release.ps1
  pwsh ./command/semantic-release.ps1 -NoDryRun
#>
param(
  [switch]$NoDryRun
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$RootDir = Resolve-Path (Join-Path $PSScriptRoot '..')
$EnvFile = Join-Path $RootDir 'config\.env'

# ---------------------------------------------------------------------------
# Load .env if GITHUB_TOKEN is not already set
# ---------------------------------------------------------------------------
if (-not $env:GITHUB_TOKEN) {
  if (Test-Path $EnvFile) {
    $line = Get-Content $EnvFile |
      Where-Object { $_ -match '^GITHUB_TOKEN=' } |
      Select-Object -First 1

    if ($line) {
      $env:GITHUB_TOKEN = $line.Substring('GITHUB_TOKEN='.Length)
    }
  }
}

if (-not $env:GITHUB_TOKEN) {
  Write-Error @"
error: GITHUB_TOKEN is not set.
  Set it in your environment or add GITHUB_TOKEN=<token> to config/.env
"@
  exit 1
}

# ---------------------------------------------------------------------------
# Run
# ---------------------------------------------------------------------------
Push-Location $RootDir
try {
  $dryRunFlag = if ($NoDryRun) { @() } else { @('--dry-run') }
  $mode = if ($NoDryRun) { '' } else { ' (dry-run)' }

  Write-Host "Running semantic-release$mode..." -ForegroundColor Cyan

  npx `
    -p semantic-release `
    -p '@semantic-release/changelog' `
    -p '@semantic-release/git' `
    -p '@semantic-release/github' `
    semantic-release `
    --branches develop `
    @dryRunFlag
}
finally {
  Pop-Location
}
