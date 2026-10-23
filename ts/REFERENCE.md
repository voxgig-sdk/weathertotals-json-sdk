# WeathertotalsJson TypeScript SDK Reference

Complete API reference for the WeathertotalsJson TypeScript SDK.


## WeathertotalsJsonSDK

### Constructor

```ts
new WeathertotalsJsonSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WeathertotalsJsonSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = WeathertotalsJsonSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `WeathertotalsJsonSDK` instance in test mode.


### Instance Methods

#### `AuroraNow(data?: object)`

Create a new `AuroraNow` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AuroraNowEntity` instance.

#### `ElNino(data?: object)`

Create a new `ElNino` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ElNinoEntity` instance.

#### `HailWeek(data?: object)`

Create a new `HailWeek` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `HailWeekEntity` instance.

#### `Rain(data?: object)`

Create a new `Rain` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RainEntity` instance.

#### `Snow(data?: object)`

Create a new `Snow` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SnowEntity` instance.

#### `SnowReport(data?: object)`

Create a new `SnowReport` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SnowReportEntity` instance.

#### `SnowSeason(data?: object)`

Create a new `SnowSeason` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SnowSeasonEntity` instance.

#### `Snowpack(data?: object)`

Create a new `Snowpack` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SnowpackEntity` instance.

#### `Storm(data?: object)`

Create a new `Storm` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `StormEntity` instance.

#### `WhiteChristma(data?: object)`

Create a new `WhiteChristma` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WhiteChristmaEntity` instance.

#### `Yesterday(data?: object)`

Create a new `Yesterday` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `YesterdayEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |
| `fetchargs.ctrl.signal` | `AbortSignal` | Aborts the request in flight: `ok` is then `false` and `err.code` is `request_aborted`. |

**Returns:** `Promise<{ ok, status, headers, data }>`. On a failure
`ok` is `false` and `err` holds the error.

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `WeathertotalsJsonSDK.test()`.

**Returns:** `WeathertotalsJsonSDK` instance in test mode.

#### Cancelling a call

Every entity operation takes an optional `ctrl` object after its match or
data, and an `AbortSignal` in `ctrl.signal` cancels the request in flight.
The operation then rejects with an error whose `code` is
`request_aborted` and whose `cause` is the signal's reason. A request
whose signal has already aborted is not sent. `stream()` takes the signal
as `callopts.signal`, and ends when it aborts.


---

## AuroraNowEntity

```ts
const aurora_now = client.AuroraNow()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.AuroraNow().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AuroraNowEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ElNinoEntity

```ts
const el_nino = client.ElNino()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.ElNino().load({ id: 'el_nino_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ElNinoEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## HailWeekEntity

```ts
const hail_week = client.HailWeek()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.HailWeek().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `HailWeekEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RainEntity

```ts
const rain = client.Rain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Rain().load({ city: 'city', state: 'state' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RainEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SnowEntity

```ts
const snow = client.Snow()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Snow().load({ city: 'city', state: 'state' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SnowEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SnowReportEntity

```ts
const snow_report = client.SnowReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.SnowReport().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SnowReportEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SnowSeasonEntity

```ts
const snow_season = client.SnowSeason()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.SnowSeason().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SnowSeasonEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SnowpackEntity

```ts
const snowpack = client.Snowpack()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Snowpack().load({ id: 'snowpack_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SnowpackEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## StormEntity

```ts
const storm = client.Storm()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `id` | `string` | No |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Storm().load({ kind: 'kind', slug: 'slug' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `StormEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WhiteChristmaEntity

```ts
const white_christma = client.WhiteChristma()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.WhiteChristma().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WhiteChristmaEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## YesterdayEntity

```ts
const yesterday = client.Yesterday()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `string` | Yes |  |
| `asOf` | `string` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `Record<string, any>` | Yes |  |
| `data` | `Record<string, any>` | Yes |  |
| `docs` | `string` | Yes |  |
| `page` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria. Resolves to the entity, whose record `data()` reads.

```ts
const result = await client.Yesterday().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `YesterdayEntity` instance with the same client and
options.

#### `client()`

Return the parent `WeathertotalsJsonSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Rate limiting |
| `retry` | 0.0.1 | Retry |
| `test` | 0.0.1 | Test transport |
| `timeout` | 0.0.1 | Timeout |


Features are activated via the `feature` option:

```ts
const client = new WeathertotalsJsonSDK({
  feature: {
    ratelimit: { active: true },
    retry: { active: true },
    test: { active: true },
    timeout: { active: true },
  }
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

