# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `intercom`, a REST client generated with oapi-codegen from Intercom's OpenAPI document (Intercom API, spec version 2.17, OpenAPI 3.0.1). It covers all 258 operations across 182 paths, including admins, contacts, companies, conversations, messages, articles, tickets and exports.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- The spec patch adds missing reporting job path parameters, resolves two inline enum Go name collisions, collapses single-branch contact body unions, and names four other inline union schemas so typed requests retain their JSON methods and serialize their fields.
- `make generate` selects the newest numbered stable spec from Intercom's public repository and regenerates the client. Initial source: https://raw.githubusercontent.com/intercom/Intercom-OpenAPI/main/descriptions/2.17/api.intercom.io.yaml. The spec is not committed.
- Added the `go-intercom` agent skill and Claude Code and Codex plugin manifests.
