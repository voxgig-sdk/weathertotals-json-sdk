# frozen_string_literal: true

# Typed models for the WeathertotalsJson SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
# params (op.<name>.points[].g.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# AuroraNow entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] page
#   @return [String]
AuroraNow = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :page,
  keyword_init: true
)

# Request payload for AuroraNow#load.
#
# @!attribute [rw] format
#   @return [String, nil]
AuroraNowLoadMatch = Struct.new(
  :format,
  keyword_init: true
)

# ElNino entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [String]
ElNino = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :id,
  :page,
  keyword_init: true
)

# Request payload for ElNino#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] format
#   @return [String, nil]
ElNinoLoadMatch = Struct.new(
  :id,
  :format,
  keyword_init: true
)

# HailWeek entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] page
#   @return [String]
HailWeek = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :page,
  keyword_init: true
)

# Request payload for HailWeek#load.
#
# @!attribute [rw] format
#   @return [String, nil]
HailWeekLoadMatch = Struct.new(
  :format,
  keyword_init: true
)

# Rain entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [String]
Rain = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :id,
  :page,
  keyword_init: true
)

# Request payload for Rain#load.
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] state
#   @return [String]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] table
#   @return [String, nil]
RainLoadMatch = Struct.new(
  :city,
  :state,
  :format,
  :table,
  keyword_init: true
)

# Snow entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [String]
Snow = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :id,
  :page,
  keyword_init: true
)

# Request payload for Snow#load.
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] state
#   @return [String]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] table
#   @return [String, nil]
SnowLoadMatch = Struct.new(
  :city,
  :state,
  :format,
  :table,
  keyword_init: true
)

# SnowReport entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] page
#   @return [String]
SnowReport = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :page,
  keyword_init: true
)

# Request payload for SnowReport#load.
#
# @!attribute [rw] format
#   @return [String, nil]
SnowReportLoadMatch = Struct.new(
  :format,
  keyword_init: true
)

# SnowSeason entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] page
#   @return [String]
SnowSeason = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :page,
  keyword_init: true
)

# Request payload for SnowSeason#load.
#
# @!attribute [rw] format
#   @return [String, nil]
SnowSeasonLoadMatch = Struct.new(
  :format,
  keyword_init: true
)

# Snowpack entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [String]
Snowpack = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :id,
  :page,
  keyword_init: true
)

# Request payload for Snowpack#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] format
#   @return [String, nil]
SnowpackLoadMatch = Struct.new(
  :id,
  :format,
  keyword_init: true
)

# Storm entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [String]
Storm = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :id,
  :page,
  keyword_init: true
)

# Request payload for Storm#load.
#
# @!attribute [rw] kind
#   @return [String]
#
# @!attribute [rw] slug
#   @return [String]
#
# @!attribute [rw] format
#   @return [String, nil]
StormLoadMatch = Struct.new(
  :kind,
  :slug,
  :format,
  keyword_init: true
)

# WhiteChristma entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] page
#   @return [String]
WhiteChristma = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :page,
  keyword_init: true
)

# Request payload for WhiteChristma#load.
#
# @!attribute [rw] format
#   @return [String, nil]
WhiteChristmaLoadMatch = Struct.new(
  :format,
  keyword_init: true
)

# Yesterday entity data model.
#
# @!attribute [rw] api
#   @return [String]
#
# @!attribute [rw] asOf
#   @return [String]
#
# @!attribute [rw] attribution
#   @return [Hash]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] docs
#   @return [String]
#
# @!attribute [rw] page
#   @return [String]
Yesterday = Struct.new(
  :api,
  :asOf,
  :attribution,
  :data,
  :docs,
  :page,
  keyword_init: true
)

# Request payload for Yesterday#load.
#
# @!attribute [rw] format
#   @return [String, nil]
YesterdayLoadMatch = Struct.new(
  :format,
  keyword_init: true
)

