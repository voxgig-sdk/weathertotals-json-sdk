# Typed models for the WeathertotalsJson SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AuroraNow(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class AuroraNowLoadMatch(TypedDict, total=False):
    format: str


class ElNinoRequired(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class ElNino(ElNinoRequired, total=False):
    id: str


class ElNinoLoadMatchRequired(TypedDict):
    id: str


class ElNinoLoadMatch(ElNinoLoadMatchRequired, total=False):
    format: str


class HailWeek(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class HailWeekLoadMatch(TypedDict, total=False):
    format: str


class RainRequired(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class Rain(RainRequired, total=False):
    id: str


class RainLoadMatchRequired(TypedDict):
    city: str
    state: str


class RainLoadMatch(RainLoadMatchRequired, total=False):
    format: str
    table: str


class SnowRequired(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class Snow(SnowRequired, total=False):
    id: str


class SnowLoadMatchRequired(TypedDict):
    city: str
    state: str


class SnowLoadMatch(SnowLoadMatchRequired, total=False):
    format: str
    table: str


class SnowReport(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class SnowReportLoadMatch(TypedDict, total=False):
    format: str


class SnowSeason(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class SnowSeasonLoadMatch(TypedDict, total=False):
    format: str


class SnowpackRequired(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class Snowpack(SnowpackRequired, total=False):
    id: str


class SnowpackLoadMatchRequired(TypedDict):
    id: str


class SnowpackLoadMatch(SnowpackLoadMatchRequired, total=False):
    format: str


class StormRequired(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class Storm(StormRequired, total=False):
    id: str


class StormLoadMatchRequired(TypedDict):
    kind: str
    slug: str


class StormLoadMatch(StormLoadMatchRequired, total=False):
    format: str


class WhiteChristma(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class WhiteChristmaLoadMatch(TypedDict, total=False):
    format: str


class Yesterday(TypedDict):
    api: str
    asOf: str
    attribution: dict
    data: dict
    docs: str
    page: str


class YesterdayLoadMatch(TypedDict, total=False):
    format: str
