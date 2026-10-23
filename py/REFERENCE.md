# WeathertotalsJson Python SDK Reference

Complete API reference for the WeathertotalsJson Python SDK.


## WeathertotalsJsonSDK

### Constructor

```python
from weathertotalsjson_sdk import WeathertotalsJsonSDK

client = WeathertotalsJsonSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WeathertotalsJsonSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = WeathertotalsJsonSDK.test()
```


### Instance Methods

#### `AuroraNow(data=None)`

Create a new `AuroraNowEntity` instance. Pass `None` for no initial data.

#### `ElNino(data=None)`

Create a new `ElNinoEntity` instance. Pass `None` for no initial data.

#### `HailWeek(data=None)`

Create a new `HailWeekEntity` instance. Pass `None` for no initial data.

#### `Rain(data=None)`

Create a new `RainEntity` instance. Pass `None` for no initial data.

#### `Snow(data=None)`

Create a new `SnowEntity` instance. Pass `None` for no initial data.

#### `SnowReport(data=None)`

Create a new `SnowReportEntity` instance. Pass `None` for no initial data.

#### `SnowSeason(data=None)`

Create a new `SnowSeasonEntity` instance. Pass `None` for no initial data.

#### `Snowpack(data=None)`

Create a new `SnowpackEntity` instance. Pass `None` for no initial data.

#### `Storm(data=None)`

Create a new `StormEntity` instance. Pass `None` for no initial data.

#### `WhiteChristma(data=None)`

Create a new `WhiteChristmaEntity` instance. Pass `None` for no initial data.

#### `Yesterday(data=None)`

Create a new `YesterdayEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AuroraNowEntity

```python
aurora_now = client.AuroraNow()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> AuroraNowEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.AuroraNow().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AuroraNowEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ElNinoEntity

```python
el_nino = client.ElNino()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `id` | `str` | No |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> ElNinoEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.ElNino().load({"id": "el_nino_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ElNinoEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HailWeekEntity

```python
hail_week = client.HailWeek()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> HailWeekEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.HailWeek().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HailWeekEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RainEntity

```python
rain = client.Rain()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `id` | `str` | No |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> RainEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Rain().load({"city": "city", "state": "state"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RainEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SnowEntity

```python
snow = client.Snow()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `id` | `str` | No |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> SnowEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Snow().load({"city": "city", "state": "state"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SnowReportEntity

```python
snow_report = client.SnowReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> SnowReportEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.SnowReport().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowReportEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SnowSeasonEntity

```python
snow_season = client.SnowSeason()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> SnowSeasonEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.SnowSeason().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowSeasonEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SnowpackEntity

```python
snowpack = client.Snowpack()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `id` | `str` | No |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> SnowpackEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Snowpack().load({"id": "snowpack_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SnowpackEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## StormEntity

```python
storm = client.Storm()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `id` | `str` | No |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> StormEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Storm().load({"kind": "kind", "slug": "slug"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `StormEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WhiteChristmaEntity

```python
white_christma = client.WhiteChristma()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> WhiteChristmaEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.WhiteChristma().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WhiteChristmaEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## YesterdayEntity

```python
yesterday = client.Yesterday()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api` | `str` | Yes |  |
| `asOf` | `str` | Yes | When the response was assembled, or the observation time where one applies. |
| `attribution` | `dict` | Yes |  |
| `data` | `dict` | Yes |  |
| `docs` | `str` | Yes |  |
| `page` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> YesterdayEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Yesterday().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `YesterdayEntity` instance with the same options.

#### `get_name() -> str`

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

```python
client = WeathertotalsJsonSDK({
    "feature": {
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
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

