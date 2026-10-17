# go-intercom

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-intercom.svg)](https://pkg.go.dev/github.com/psyb0t/go-intercom)
[![CI](https://github.com/psyb0t/go-intercom/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-intercom/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-intercom/badges/version.svg)](https://github.com/psyb0t/go-intercom/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-intercom/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-intercom/badges/importers.svg)](https://github.com/psyb0t/go-intercom/blob/badges/importers.md)

Unofficial Go client for the [Intercom API](https://developers.intercom.com/docs/references/rest-api/api.intercom.io), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Intercom publishes. I am not affiliated with Intercom.

The package is `intercom`, in the repository root. Everything in `client_*.gen.go` is generated. Don't edit those files by hand, regenerate them.

## Install

```sh
go get github.com/psyb0t/go-intercom
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-intercom"
)

func main() {
	auth := intercom.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("INTERCOM_ACCESS_TOKEN"),
			)
			req.Header.Set("Intercom-Version", "2.17")

			return nil
		},
	)

	client, err := intercom.NewClientWithResponses(
		"https://api.intercom.io",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.IdentifyAdminWithResponse(
		context.Background(),
		&intercom.IdentifyAdminParams{},
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

Method names follow the spec's operation IDs: `IdentifyAdminWithResponse` calls `GET /me`, `ListContactsWithResponse` lists contacts, and `ListConversationsWithResponse` lists conversations. Methods return transport and decoding errors; check `StatusCode()`, `Body` and the typed status field for API errors.

## Auth

Intercom uses a bearer access token. Read it from `INTERCOM_ACCESS_TOKEN` and set `Authorization: Bearer <token>` with `WithRequestEditorFn`, as above. The client does not attach authentication automatically.

The optional `Intercom-Version` header otherwise defaults to the version configured for the Intercom app. The example pins it to `2.17`, matching the generated models. Use the base URL for your workspace's data region:

| Region | Base URL |
|---|---|
| United States | `https://api.intercom.io` |
| Europe | `https://api.eu.intercom.io` |
| Australia | `https://api.au.intercom.io` |

## Regenerating

The spec is not committed. `make generate` downloads the newest numbered stable version into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` lists the provider's [version directories](https://github.com/intercom/Intercom-OpenAPI/tree/main/descriptions), skips `0` and preview names, sorts stable version components numerically, and downloads `api.intercom.io.yaml` to `api/openapi.yaml`.
- `scripts/patch-spec.sh` runs the YAML tool under `internal/cmd/patch-spec/`, writing `api/openapi.fixed.yaml`. It adds missing required string `job_identifier` path parameters to reporting export status and download. It also gives the inline `type` enums on `conversation_attribute_list` and `ticket` distinct Go type names, avoiding collisions with component names while preserving JSON fields and enum values. Single-branch contact request unions become direct schema references so typed create and update requests serialize their fields instead of sending `{}`. Inline unions for Fin escalation, conversation parts, tags and ticket replies move unchanged into named components so request body aliases retain their JSON methods.
- oapi-codegen runs with `config.yaml`, generating every operation. The config maps `number` to `float64` and prefixes enum values to avoid name collisions.
- `scripts/split-gen.sh` splits the generated output at declaration boundaries and runs goimports. Smaller files keep linting practical.

The initial client uses [Intercom API 2.17](https://raw.githubusercontent.com/intercom/Intercom-OpenAPI/main/descriptions/2.17/api.intercom.io.yaml), an OpenAPI 3.0.1 document with 258 operations across 182 paths.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-intercom) teaches a coding agent how to use go-intercom when it writes Go. It works in any agent that reads `.agents/skills/`. The marketplace commands below require publication of the repository and its catalog entries.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-intercom@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-intercom@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The release workflow publishes the skill to ClawHub. After publication:

```bash
openclaw skills install @psyb0t/go-intercom
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). Intercom labels its [published specification](https://raw.githubusercontent.com/intercom/Intercom-OpenAPI/main/descriptions/2.17/api.intercom.io.yaml) MIT too. The specification itself is not redistributed here.
