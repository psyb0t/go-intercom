# Go code

- **GO1** `client_*.gen.go` files are generated. Never edit them by hand. Change `config.yaml` or the spec patch in `scripts/`, then run `make generate`.
- **GO2** Spec fixes run before generation through `scripts/patch-spec.sh` and the YAML tool under `internal/cmd/patch-spec/`. Each patch explains the upstream problem it fixes.
- **GO3** The downloaded spec in `api/` is gitignored and never committed.
- **GO4** Number fields stay `float64` through the `type-mapping` in `config.yaml`. Money in `float32` loses precision.
- **GO5** Hand-written Go (`gen.go` and the tests) wraps errors with `ctxerrors.Wrap` against `ctxerrors/commerr` sentinels. No `fmt.Errorf`, no `errors.New` outside a sentinel declaration.
- **GO6** The exported API is whatever oapi-codegen emits. A config change that renames or removes an exported name is a breaking release (see `releases.md`).
- **GO7** `make lint` passes with zero issues. Hand-written lines are at most 80 columns, the `.golangci.yml` limit.
