# WeathertotalsJson Golang SDK Reference

Complete API reference for the WeathertotalsJson Golang SDK.


## WeathertotalsJsonSDK

### Constructor

```go
func NewWeathertotalsJsonSDK(options map[string]any) *WeathertotalsJsonSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *WeathertotalsJsonSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *WeathertotalsJsonSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `AuroraNow(data map[string]any) WeathertotalsJsonEntity`

Create a new `AuroraNow` entity instance. Pass `nil` for no initial data.

#### `ElNino(data map[string]any) WeathertotalsJsonEntity`

Create a new `ElNino` entity instance. Pass `nil` for no initial data.

#### `HailWeek(data map[string]any) WeathertotalsJsonEntity`

Create a new `HailWeek` entity instance. Pass `nil` for no initial data.

#### `Rain(data map[string]any) WeathertotalsJsonEntity`

Create a new `Rain` entity instance. Pass `nil` for no initial data.

#### `Snow(data map[string]any) WeathertotalsJsonEntity`

Create a new `Snow` entity instance. Pass `nil` for no initial data.

#### `SnowReport(data map[string]any) WeathertotalsJsonEntity`

Create a new `SnowReport` entity instance. Pass `nil` for no initial data.

#### `SnowSeason(data map[string]any) WeathertotalsJsonEntity`

Create a new `SnowSeason` entity instance. Pass `nil` for no initial data.

#### `Snowpack(data map[string]any) WeathertotalsJsonEntity`

Create a new `Snowpack` entity instance. Pass `nil` for no initial data.

#### `Storm(data map[string]any) WeathertotalsJsonEntity`

Create a new `Storm` entity instance. Pass `nil` for no initial data.

#### `WhiteChristma(data map[string]any) WeathertotalsJsonEntity`

Create a new `WhiteChristma` entity instance. Pass `nil` for no initial data.

#### `Yesterday(data map[string]any) WeathertotalsJsonEntity`

Create a new `Yesterday` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AuroraNowEntity

```go
auroraNow := client.AuroraNow(nil)
fmt.Println(auroraNow.GetName()) // "aurora_now"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.AuroraNow(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AuroraNowEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## ElNinoEntity

```go
elNino := client.ElNino(nil)
fmt.Println(elNino.GetName()) // "el_nino"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.ElNino(nil).Load(map[string]any{"id": "el_nino_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ElNinoEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## HailWeekEntity

```go
hailWeek := client.HailWeek(nil)
fmt.Println(hailWeek.GetName()) // "hail_week"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.HailWeek(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `HailWeekEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## RainEntity

```go
rain := client.Rain(nil)
fmt.Println(rain.GetName()) // "rain"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Rain(nil).Load(map[string]any{"city": "city", "state": "state"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RainEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SnowEntity

```go
snow := client.Snow(nil)
fmt.Println(snow.GetName()) // "snow"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Snow(nil).Load(map[string]any{"city": "city", "state": "state"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SnowEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SnowReportEntity

```go
snowReport := client.SnowReport(nil)
fmt.Println(snowReport.GetName()) // "snow_report"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.SnowReport(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SnowReportEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SnowSeasonEntity

```go
snowSeason := client.SnowSeason(nil)
fmt.Println(snowSeason.GetName()) // "snow_season"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.SnowSeason(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SnowSeasonEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## SnowpackEntity

```go
snowpack := client.Snowpack(nil)
fmt.Println(snowpack.GetName()) // "snowpack"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Snowpack(nil).Load(map[string]any{"id": "snowpack_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SnowpackEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## StormEntity

```go
storm := client.Storm(nil)
fmt.Println(storm.GetName()) // "storm"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Storm(nil).Load(map[string]any{"kind": "kind", "slug": "slug"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `StormEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## WhiteChristmaEntity

```go
whiteChristma := client.WhiteChristma(nil)
fmt.Println(whiteChristma.GetName()) // "white_christma"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.WhiteChristma(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WhiteChristmaEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

Return the entity name.


---

## YesterdayEntity

```go
yesterday := client.Yesterday(nil)
fmt.Println(yesterday.GetName()) // "yesterday"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `map[string]any` | Yes |  |
| `data` | `map[string]any` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria. Returns the entity, whose record `Data()` reads; `err` is non-nil on failure.

```go
result, err := client.Yesterday(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result.(sdk.Entity).Data())
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `YesterdayEntity` instance with the same client and
options.

#### `Stream(action string, args map[string]any, callopts map[string]any) <-chan StreamItem`

Run an operation through the pipeline and send its result items on the
returned channel, which closes when the stream ends. A `StreamItem` holds
one item in `Item`, or in `Err` the error that ended the stream: the
error the operation itself would return, sent as the last value. Under
`throw: false` in `callopts["ctrl"]`, no error is sent.

#### `GetName() string`

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

```go
client := sdk.NewWeathertotalsJsonSDK(map[string]any{
    "feature": map[string]any{
        "ratelimit": map[string]any{"active": true},
        "retry": map[string]any{"active": true},
        "test": map[string]any{"active": true},
        "timeout": map[string]any{"active": true},
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

