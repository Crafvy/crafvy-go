# Crafvy — Go SDK

The official Go client of the [Crafvy](https://crafvy.com) API. Every
resource and method is generated from the API's description, the same one the
TypeScript, Python, PHP, Java and C# SDKs and the [reference](https://docs.crafvy.com/sdk)
come from — so a method here is exactly the one documented there.

Standard library only; Go 1.21 or newer.

```sh
go get github.com/Crafvy/crafvy-go@v0.2.0
```

## Use

```go
import (
	"context"
	"log"
	"os"

	"github.com/Crafvy/crafvy-go"
)

client := crafvy.NewClient(os.Getenv("CRAFVY_TOKEN")) // "" also reads CRAFVY_TOKEN
// client := crafvy.NewClient(token, crafvy.WithBaseURL("https://your-instance.example")) // self-hosted

ctx := context.Background()
projects, err := client.Projects.List(ctx, nil)
if err != nil {
	log.Fatal(err)
}
for _, p := range projects {
	log.Println(p.Name)
}

spaces, err := client.CRM.Spaces.List(ctx, &crafvy.CrmSpacesListQuery{WorkspaceID: "…"})
```

- Resources are fields (`client.CRM.Records`), methods are exported
  (`Create`), names use Go's initialisms (`SpaceID`, `CRM`, `URL`).
- Every method takes `ctx` first, then the path parameters, then the body
  (a pointer to its struct), then the query (a pointer to its struct, `nil`
  when you have none), then call options — and returns `(T, error)`.
- A field the API may leave out is a pointer: set it with `crafvy.String`,
  `crafvy.Int`, `crafvy.Float64`, `crafvy.Bool` (or `crafvy.Ptr`). An absent
  field — a nil pointer, slice or map — is not sent; an explicit `false`, `0`
  or empty list is. A required field that may be null is a pointer without
  `omitempty`: nil sends `null`.
- Enums are string types with constants (`crafvy.AssetVisibilityPrivate`).
- A body that is one of several shapes, or an answer that is, is `any`: pass
  the struct you mean; the answer comes as generic JSON (`map[string]any`).
- `client.Raw(ctx, "GET", "/api/...", query, body)` calls any endpoint and
  returns the answer as `json.RawMessage`.

## Tokens

A token acts as its owner: their roles, workspaces, feature gates and plan
limits apply exactly as in the dashboard. Create one under **Settings →
Integrations → API Access Tokens**. A token made with permissions only reaches
those families of routes; each method's doc comment names the permission it
needs.

## Errors

```go
_, err := client.Projects.Get(ctx, "missing")
if errors.Is(err, crafvy.ErrNotFound) {
	// …
}
var apiErr *crafvy.Error
if errors.As(err, &apiErr) {
	log.Println(apiErr.Status, apiErr.Code, apiErr.Message)
}
```

`Code` is the API's stable reason (`tokenPermissionMissing`, …) — branch on
it, not on the English text. Sentinels: `ErrValidation` (400),
`ErrAuthentication` (401), `ErrAuthorization` (403), `ErrNotFound` (404),
`ErrConflict` (409), `ErrRateLimit` (429, `RetryAfter` set when the API said).
429 answers are retried with back-off (respecting `Retry-After`); a 5xx is
retried only for GET, HEAD, PUT and DELETE, never for a POST that may have
been carried out. A timeout is an `*Error` with status 408.

## Options

```go
client := crafvy.NewClient(token,
	crafvy.WithRetries(5),
	crafvy.WithTimeout(30*time.Second),
	crafvy.WithHTTPClient(&http.Client{Transport: myTransport}),
)
client.Projects.List(ctx, nil, crafvy.CallTimeout(5*time.Second), crafvy.CallHeader("X-Trace", "1"))
```

Tests can pass an `http.Client` whose `Transport` records requests and
answers them, instead of real HTTP.

## Uploads

```go
file, err := crafvy.FileFromPath("logo.png", "image/png")
asset, err := client.Assets.Upload(ctx, &crafvy.AssetsUploadBody{File: file, OrganizationID: org})
```

## Releases

Go modules publish by tag: `go get github.com/Crafvy/crafvy-go@vX.Y.Z` works
as soon as the tag exists. Changelog: [CHANGELOG.md](CHANGELOG.md). Generated
— changes go into the API's description in the Crafvy repository, not into
this package.
