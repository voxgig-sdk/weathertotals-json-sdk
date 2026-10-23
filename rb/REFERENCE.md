# WeathertotalsJson Ruby SDK Reference

Complete API reference for the WeathertotalsJson Ruby SDK.


## WeathertotalsJsonSDK

### Constructor

```ruby
require_relative 'WeathertotalsJson_sdk'

client = WeathertotalsJsonSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WeathertotalsJsonSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = WeathertotalsJsonSDK.test
```


### Instance Methods

#### `AuroraNow(data = nil)`

Create a new `AuroraNow` entity instance. Pass `nil` for no initial data.

#### `ElNino(data = nil)`

Create a new `ElNino` entity instance. Pass `nil` for no initial data.

#### `HailWeek(data = nil)`

Create a new `HailWeek` entity instance. Pass `nil` for no initial data.

#### `Rain(data = nil)`

Create a new `Rain` entity instance. Pass `nil` for no initial data.

#### `Snow(data = nil)`

Create a new `Snow` entity instance. Pass `nil` for no initial data.

#### `SnowReport(data = nil)`

Create a new `SnowReport` entity instance. Pass `nil` for no initial data.

#### `SnowSeason(data = nil)`

Create a new `SnowSeason` entity instance. Pass `nil` for no initial data.

#### `Snowpack(data = nil)`

Create a new `Snowpack` entity instance. Pass `nil` for no initial data.

#### `Storm(data = nil)`

Create a new `Storm` entity instance. Pass `nil` for no initial data.

#### `WhiteChristma(data = nil)`

Create a new `WhiteChristma` entity instance. Pass `nil` for no initial data.

#### `Yesterday(data = nil)`

Create a new `Yesterday` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AuroraNowEntity

```ruby
aurora_now = client.AuroraNow
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.AuroraNow.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AuroraNowEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ElNinoEntity

```ruby
el_nino = client.ElNino
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `id` | `String` | No |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.ElNino.load({ "id" => "el_nino_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ElNinoEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## HailWeekEntity

```ruby
hail_week = client.HailWeek
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.HailWeek.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `HailWeekEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RainEntity

```ruby
rain = client.Rain
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `id` | `String` | No |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.Rain.load({ "city" => "city", "state" => "state" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RainEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SnowEntity

```ruby
snow = client.Snow
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `id` | `String` | No |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.Snow.load({ "city" => "city", "state" => "state" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SnowEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SnowReportEntity

```ruby
snow_report = client.SnowReport
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.SnowReport.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SnowReportEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SnowSeasonEntity

```ruby
snow_season = client.SnowSeason
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.SnowSeason.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SnowSeasonEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SnowpackEntity

```ruby
snowpack = client.Snowpack
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `id` | `String` | No |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.Snowpack.load({ "id" => "snowpack_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SnowpackEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## StormEntity

```ruby
storm = client.Storm
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `id` | `String` | No |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.Storm.load({ "kind" => "kind", "slug" => "slug" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `StormEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## WhiteChristmaEntity

```ruby
white_christma = client.WhiteChristma
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.WhiteChristma.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `WhiteChristmaEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## YesterdayEntity

```ruby
yesterday = client.Yesterday
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `String` | Yes |  |
| `asOf` | `String` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Hash` | Yes |  |
| `data` | `Hash` | Yes |  |
| `docs` | `String` | Yes |  |
| `page` | `String` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get` reads, and raises on error.

```ruby
result = client.Yesterday.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `YesterdayEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```ruby
client = WeathertotalsJsonSDK.new({
  "feature" => {
    "ratelimit" => { "active" => true },
    "retry" => { "active" => true },
    "test" => { "active" => true },
    "timeout" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Rate limiting.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Retry.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Timeout.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `now` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

