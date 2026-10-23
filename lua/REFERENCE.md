# WeathertotalsJson Lua SDK Reference

Complete API reference for the WeathertotalsJson Lua SDK.


## WeathertotalsJsonSDK

### Constructor

```lua
local sdk = require("weathertotals-json_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `AuroraNow(data)`

Create a new `AuroraNow` entity instance. Pass `nil` for no initial data.

#### `ElNino(data)`

Create a new `ElNino` entity instance. Pass `nil` for no initial data.

#### `HailWeek(data)`

Create a new `HailWeek` entity instance. Pass `nil` for no initial data.

#### `Rain(data)`

Create a new `Rain` entity instance. Pass `nil` for no initial data.

#### `Snow(data)`

Create a new `Snow` entity instance. Pass `nil` for no initial data.

#### `SnowReport(data)`

Create a new `SnowReport` entity instance. Pass `nil` for no initial data.

#### `SnowSeason(data)`

Create a new `SnowSeason` entity instance. Pass `nil` for no initial data.

#### `Snowpack(data)`

Create a new `Snowpack` entity instance. Pass `nil` for no initial data.

#### `Storm(data)`

Create a new `Storm` entity instance. Pass `nil` for no initial data.

#### `WhiteChristma(data)`

Create a new `WhiteChristma` entity instance. Pass `nil` for no initial data.

#### `Yesterday(data)`

Create a new `Yesterday` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AuroraNowEntity

```lua
local aurora_now = client:AuroraNow(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:AuroraNow():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AuroraNowEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ElNinoEntity

```lua
local el_nino = client:ElNino(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:ElNino():load({ id = "el_nino_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ElNinoEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## HailWeekEntity

```lua
local hail_week = client:HailWeek(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:HailWeek():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HailWeekEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RainEntity

```lua
local rain = client:Rain(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:Rain():load({ city = "city", state = "state" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RainEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SnowEntity

```lua
local snow = client:Snow(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:Snow():load({ city = "city", state = "state" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SnowReportEntity

```lua
local snow_report = client:SnowReport(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:SnowReport():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowReportEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SnowSeasonEntity

```lua
local snow_season = client:SnowSeason(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:SnowSeason():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowSeasonEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SnowpackEntity

```lua
local snowpack = client:Snowpack(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:Snowpack():load({ id = "snowpack_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowpackEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## StormEntity

```lua
local storm = client:Storm(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:Storm():load({ kind = "kind", slug = "slug" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `StormEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WhiteChristmaEntity

```lua
local white_christma = client:WhiteChristma(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:WhiteChristma():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WhiteChristmaEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## YesterdayEntity

```lua
local yesterday = client:Yesterday(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `table` | Yes |  |
| `data` | `table` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, or `nil` and an error on failure.

```lua
local result, err = client:Yesterday():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `YesterdayEntity` instance with the same client and
options.

#### `get_name() -> string`

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

```lua
local client = sdk.new({
  feature = {
    ratelimit = { active = true },
    retry = { active = true },
    test = { active = true },
    timeout = { active = true },
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

