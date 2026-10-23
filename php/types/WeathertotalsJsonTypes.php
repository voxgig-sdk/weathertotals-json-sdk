<?php
declare(strict_types=1);

// Typed models for the WeathertotalsJson SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** AuroraNow entity data model. */
class AuroraNow
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public string $page;
}

/** Request payload for AuroraNow#load. */
class AuroraNowLoadMatch
{
    public ?string $format = null;
}

/** ElNino entity data model. */
class ElNino
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public ?string $id = null;
    public string $page;
}

/** Request payload for ElNino#load. */
class ElNinoLoadMatch
{
    public string $id;
    public ?string $format = null;
}

/** HailWeek entity data model. */
class HailWeek
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public string $page;
}

/** Request payload for HailWeek#load. */
class HailWeekLoadMatch
{
    public ?string $format = null;
}

/** Rain entity data model. */
class Rain
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public ?string $id = null;
    public string $page;
}

/** Request payload for Rain#load. */
class RainLoadMatch
{
    public string $city;
    public string $state;
    public ?string $format = null;
    public ?string $table = null;
}

/** Snow entity data model. */
class Snow
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public ?string $id = null;
    public string $page;
}

/** Request payload for Snow#load. */
class SnowLoadMatch
{
    public string $city;
    public string $state;
    public ?string $format = null;
    public ?string $table = null;
}

/** SnowReport entity data model. */
class SnowReport
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public string $page;
}

/** Request payload for SnowReport#load. */
class SnowReportLoadMatch
{
    public ?string $format = null;
}

/** SnowSeason entity data model. */
class SnowSeason
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public string $page;
}

/** Request payload for SnowSeason#load. */
class SnowSeasonLoadMatch
{
    public ?string $format = null;
}

/** Snowpack entity data model. */
class Snowpack
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public ?string $id = null;
    public string $page;
}

/** Request payload for Snowpack#load. */
class SnowpackLoadMatch
{
    public string $id;
    public ?string $format = null;
}

/** Storm entity data model. */
class Storm
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public ?string $id = null;
    public string $page;
}

/** Request payload for Storm#load. */
class StormLoadMatch
{
    public string $kind;
    public string $slug;
    public ?string $format = null;
}

/** WhiteChristma entity data model. */
class WhiteChristma
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public string $page;
}

/** Request payload for WhiteChristma#load. */
class WhiteChristmaLoadMatch
{
    public ?string $format = null;
}

/** Yesterday entity data model. */
class Yesterday
{
    public string $api;
    public string $asOf;
    public array $attribution;
    public array $data;
    public string $docs;
    public string $page;
}

/** Request payload for Yesterday#load. */
class YesterdayLoadMatch
{
    public ?string $format = null;
}

