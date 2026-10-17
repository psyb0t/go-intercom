// Package intercom is a Go client for the Intercom REST API,
// generated with oapi-codegen from the OpenAPI document Intercom publishes.
package intercom

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/patch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.fixed.yaml
//go:generate bash scripts/split-gen.sh
