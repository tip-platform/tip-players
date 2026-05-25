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
