// Typed models for the WeathertotalsJson SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface AuroraNow {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  page: string
}

export interface AuroraNowLoadMatch {
  format?: string
}

export interface ElNino {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  id?: string
  page: string
}

export interface ElNinoLoadMatch {
  id: string
  format?: string
}

export interface HailWeek {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  page: string
}

export interface HailWeekLoadMatch {
  format?: string
}

export interface Rain {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  id?: string
  page: string
}

export interface RainLoadMatch {
  city: string
  state: string
  format?: string
  table?: string
}

export interface Snow {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  id?: string
  page: string
}

export interface SnowLoadMatch {
  city: string
  state: string
  format?: string
  table?: string
}

export interface SnowReport {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  page: string
}

export interface SnowReportLoadMatch {
  format?: string
}

export interface SnowSeason {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  page: string
}

export interface SnowSeasonLoadMatch {
  format?: string
}

export interface Snowpack {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  id?: string
  page: string
}

export interface SnowpackLoadMatch {
  id: string
  format?: string
}

export interface Storm {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  id?: string
  page: string
}

export interface StormLoadMatch {
  kind: string
  slug: string
  format?: string
}

export interface WhiteChristma {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  page: string
}

export interface WhiteChristmaLoadMatch {
  format?: string
}

export interface Yesterday {
  api: string
  asOf: string
  attribution: Record<string, any>
  data: Record<string, any>
  docs: string
  page: string
}

export interface YesterdayLoadMatch {
  format?: string
}

