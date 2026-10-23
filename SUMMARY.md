# WeatherTotals JSON API

The WeatherTotals JSON API serves the same figures that appear on the site&#39;s pages as structured data. It is free, requires no key or signup, and uses only GET requests.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 11 entities and 11 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### AuroraNow

Results: Cached for 300 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### ElNino

Results: Cached for 86400 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### HailWeek

Results: Cached for 900 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### Rain

Results: Cached for 1800 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.
- `id`: GHCN-Daily station identifier.

### Snow

Results: Cached for 1800 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.
- `id`: GHCN-Daily station identifier.

### SnowReport

Results: Cached for 900 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### SnowSeason

Results: Cached for 1800 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### Snowpack

Results: Cached for 3600 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### Storm

Results: Cached for 900 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### WhiteChristma

Results: Cached for 86400 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### Yesterday

Results: Cached for 3600 seconds.

SDK operations: `load`.

Key fields to recognise:

- `asOf`: When the response was assembled, or the observation time where one applies.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| AuroraNow | `load` | `GET /api/v1/aurora/now` | Not required |
| ElNino | `load` | `GET /api/v1/el-nino/{state}` | Not required |
| HailWeek | `load` | `GET /api/v1/hail/week` | Not required |
| Rain | `load` | `GET /api/v1/rain/{state}/{city}` | Not required |
| Snow | `load` | `GET /api/v1/snow/{state}/{city}` | Not required |
| SnowReport | `load` | `GET /api/v1/snow/reports` | Not required |
| SnowSeason | `load` | `GET /api/v1/snow/season` | Not required |
| Snowpack | `load` | `GET /api/v1/snowpack/{state}` | Not required |
| Storm | `load` | `GET /api/v1/storms/{kind}/{slug}` | Not required |
| WhiteChristma | `load` | `GET /api/v1/white-christmas` | Not required |
| Yesterday | `load` | `GET /api/v1/yesterday` | Not required |

## Connect to the API

- API server: `https://weathertotals.com`

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

A read request without required parameters or authentication is `GET /api/v1/aurora/now`. For example:

```sh
curl --fail-with-body --silent --show-error 'https://weathertotals.com/api/v1/aurora/now'
```

Inspect the response using the AuroraNow reference. This checks the public route; authenticated operations need their own credentials and request data.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Lua | `lua/` | Build from source |
| PHP | `php/` | Build from source |
| Python | `py/` | Build from source |
| Ruby | `rb/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### Go CLI

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### Go MCP server

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `weathertotals-json_list`: List records for an entity. No active entity supports this operation.
- `weathertotals-json_load`: Load one record for an entity. Supported entities: `aurora_now`, `el_nino`, `hail_week`, `rain`, `snow`, `snow_report`, `snow_season`, `snowpack`, `storm`, `white_christma`, `yesterday`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `ratelimit`: Client-side rate limiting via a token bucket
- `retry`: Automatic retry of transient failures with exponential backoff
- `test`: In-memory mock transport for testing without a live server
- `timeout`: Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

