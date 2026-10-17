---
name: go-intercom
description: Unofficial Go client for the Intercom REST API, generated with oapi-codegen. Package `intercom` (`github.com/psyb0t/go-intercom`) provides typed methods for admins, contacts, companies, conversations, messages, articles, tickets, tags, teams and exports. Use when Go code needs Intercom customer support or workspace data, must perform a requested Intercom action, or already imports go-intercom.
homepage: https://github.com/psyb0t/go-intercom
user-invocable: true
metadata:
  openclaw:
    emoji: "💬"
    requires:
      bins:
        - go
      env:
        - INTERCOM_ACCESS_TOKEN
permissions:
  filesystem: "Guides edits to Go source in the current project. The library reads and writes no files."
  shell: "The Go toolchain (`go get`, `go build`, `go test`) to add and compile the dependency."
  network: "Module downloads use the configured Go proxy. Client requests send the chosen API payload and bearer token over HTTPS to api.intercom.io, api.eu.intercom.io or api.au.intercom.io for the workspace's region."
---

# go-intercom

Unofficial, typed Go client for the Intercom REST API, in package `intercom` at the module root.

For install, main operations, pagination and response fields, see [references/setup.md](references/setup.md).

## Security & safety

- Read the bearer token from `INTERCOM_ACCESS_TOKEN`. Never hardcode, commit or log it, or print a request with its `Authorization` header.
- Use the HTTPS API host matching the workspace's region. Do not send credentials to an arbitrary URL. Pin `Intercom-Version: 2.17` to match the generated models.
- Sending messages, changing customer data and deleting workspace resources act on the connected workspace. Perform them only when requested and only for the requested resources.

## When to use

- Go code needs Intercom contacts, companies, conversations, admins, articles or tickets.
- Go code must send a message or update Intercom resources at the user's request.
- The project already imports `github.com/psyb0t/go-intercom`.

## When NOT to use

- The code is not Go. This module ships only a Go package.
- You need a client maintained by Intercom. This one is unofficial and not affiliated with Intercom.
- You want to change the generated client itself. See `AGENTS.md` in the repo.

## Usage

```go
import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-intercom"
)

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
	return err
}

res, err := client.IdentifyAdminWithResponse(
	ctx,
	&intercom.IdentifyAdminParams{},
)
if err != nil {
	return err
}

if res.JSON200 != nil {
	_ = res.JSON200
}
```

Many response fields are pointers; check for nil before dereferencing. The error return covers transport, request construction and decoding failures. HTTP error statuses still return a response: inspect `StatusCode()`, `Body` and any typed error field. The generated client does not retry or paginate automatically.
