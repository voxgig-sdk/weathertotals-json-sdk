-- Typed models for the WeathertotalsJson SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
-- params (op.<name>.points[].g.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class AuroraNow
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field page string

---@class AuroraNowLoadMatch
---@field format? string

---@class ElNino
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field id? string
---@field page string

---@class ElNinoLoadMatch
---@field id string
---@field format? string

---@class HailWeek
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field page string

---@class HailWeekLoadMatch
---@field format? string

---@class Rain
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field id? string
---@field page string

---@class RainLoadMatch
---@field city string
---@field state string
---@field format? string
---@field table? string

---@class Snow
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field id? string
---@field page string

---@class SnowLoadMatch
---@field city string
---@field state string
---@field format? string
---@field table? string

---@class SnowReport
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field page string

---@class SnowReportLoadMatch
---@field format? string

---@class SnowSeason
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field page string

---@class SnowSeasonLoadMatch
---@field format? string

---@class Snowpack
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field id? string
---@field page string

---@class SnowpackLoadMatch
---@field id string
---@field format? string

---@class Storm
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field id? string
---@field page string

---@class StormLoadMatch
---@field kind string
---@field slug string
---@field format? string

---@class WhiteChristma
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field page string

---@class WhiteChristmaLoadMatch
---@field format? string

---@class Yesterday
---@field api string
---@field asOf string
---@field attribution table
---@field data table
---@field docs string
---@field page string

---@class YesterdayLoadMatch
---@field format? string

local M = {}

return M
