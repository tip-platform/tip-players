# command/

Local, cross-platform helper scripts.

## golangci-lint

Runs `golangci-lint` using the repository config `.golangci.yml`.

### Linux/macOS

```bash
./command/lint-golangci.sh
```

### Windows (PowerShell)

```powershell
pwsh ./command/lint-golangci.ps1
# or (Windows PowerShell 5.1 if pwsh is not available)
powershell -ExecutionPolicy Bypass -File .\command\lint-golangci.ps1
```

### Pin / override version

```bash
GOLANGCI_LINT_VERSION=v1.64.6 ./command/lint-golangci.sh
```

```powershell
$env:GOLANGCI_LINT_VERSION = 'v1.64.6'
pwsh ./command/lint-golangci.ps1
```

---

## semantic-release

Analyzes commits and computes the next release version using [semantic-release](https://semantic-release.gitbook.io).
Runs in **dry-run mode by default** — no tags or releases are created unless explicitly requested.

Requires a GitHub token with `Contents`, `Pull requests`, and `Issues` write permissions on the repository.

### Token setup

Create `config/.env` (already in `.gitignore`) and add:

```env
GITHUB_TOKEN=your_token_here
```

Alternatively, export it in your shell before running:

```bash
export GITHUB_TOKEN=your_token_here
```

The script reads the environment variable first; `config/.env` is only loaded as fallback.

### Linux/macOS

```bash
# Dry-run (default) — shows next version and changelog, nothing is published
./command/semantic-release.sh

# Publish — creates tag, GitHub release, and CHANGELOG.md commit
./command/semantic-release.sh --no-dry-run
```

### Windows (PowerShell)

```powershell
# Dry-run (default)
pwsh ./command/semantic-release.ps1

# Publish
pwsh ./command/semantic-release.ps1 -NoDryRun

# or (Windows PowerShell 5.1)
powershell -ExecutionPolicy Bypass -File .\command\semantic-release.ps1
powershell -ExecutionPolicy Bypass -File .\command\semantic-release.ps1 -NoDryRun
```

### Override target branch

```bash
# Linux/macOS — pass extra flags directly
./command/semantic-release.sh --branches main
```

```powershell
# Windows — not yet supported via flag; set manually in the script if needed
```

> **Note:** In CI, `GITHUB_TOKEN` is injected automatically by GitHub Actions.
> The `config/.env` fallback is for local development only and must never be committed.
