## What does this PR do?

<!-- Describe the purpose the change concisely. One sentence is enough if it is
simple. -->

## Motivation

<!-- Why is this change needed? What problem does it solve? Link the issue if applicable: Closes #123 -->

## Type of change

- [ ] `feat` — new feature
- [ ] `fix` — bug fix
- [ ] `refactor` — change code without altering behavior
- [ ] `chore` — maintenance tasks / update dependencies
- [ ] `docs` — documentation only changes
- [ ] `test` — adding or correcting tests
- [ ] `perf` — performance improvements
- [ ] `ci` — changes pipelines / workflow
- [ ] `build` — system build / packaging / release
- [ ] `infra` — infrastructure as code
- [ ] `ops` — runtime configuration

## Who was tested?

<!-- Describe what you ran locally to verify the changes. -->

```bash
go test -race ./...
# o bien:
# go run ./cmd/... <args>
```

## Checklist

- [ ] Tests pass locally (`go test -race ./...`)
- [ ] The code compiles cleanly (`go build ./...`)
- [ ] No new warnings for `golangci-lint`
- [ ] Public API changes have godoc update
- [ ] The title of the PR follows Conventional Commits (`type(scope): description`)
- [ ] The branch follows the convention (`type/descripcion-kebab-case`)

## Notes for reviewers

<!-- extra context, design decisions, areas that deserve special attention. -->
