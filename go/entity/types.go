// Typed models for the WeathertotalsJson SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields{} and per-op
// params (op.<name>.points[].g.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/weathertotals-json-sdk/go/core"
)

// AuroraNow is the typed data model for the aurora_now entity.
type AuroraNow struct {
}

// AuroraNowLoadMatch is the typed request payload for AuroraNow.LoadTyped.
type AuroraNowLoadMatch struct {
	Format *string `json:"format,omitempty"`
}

// ElNino is the typed data model for the el_nino entity.
type ElNino struct {
}

// ElNinoLoadMatch is the typed request payload for ElNino.LoadTyped.
type ElNinoLoadMatch struct {
	Id string `json:"id"`
	Format *string `json:"format,omitempty"`
}

// HailWeek is the typed data model for the hail_week entity.
type HailWeek struct {
}

// HailWeekLoadMatch is the typed request payload for HailWeek.LoadTyped.
type HailWeekLoadMatch struct {
	Format *string `json:"format,omitempty"`
}

// Rain is the typed data model for the rain entity.
type Rain struct {
}

// RainLoadMatch is the typed request payload for Rain.LoadTyped.
type RainLoadMatch struct {
	City string `json:"city"`
	State string `json:"state"`
	Format *string `json:"format,omitempty"`
	Table *string `json:"table,omitempty"`
}

// Snow is the typed data model for the snow entity.
type Snow struct {
}

// SnowLoadMatch is the typed request payload for Snow.LoadTyped.
type SnowLoadMatch struct {
	City string `json:"city"`
	State string `json:"state"`
	Format *string `json:"format,omitempty"`
	Table *string `json:"table,omitempty"`
}

// SnowReport is the typed data model for the snow_report entity.
type SnowReport struct {
}

// SnowReportLoadMatch is the typed request payload for SnowReport.LoadTyped.
type SnowReportLoadMatch struct {
	Format *string `json:"format,omitempty"`
}

// SnowSeason is the typed data model for the snow_season entity.
type SnowSeason struct {
}

// SnowSeasonLoadMatch is the typed request payload for SnowSeason.LoadTyped.
type SnowSeasonLoadMatch struct {
	Format *string `json:"format,omitempty"`
}

// Snowpack is the typed data model for the snowpack entity.
type Snowpack struct {
}

// SnowpackLoadMatch is the typed request payload for Snowpack.LoadTyped.
type SnowpackLoadMatch struct {
	Id string `json:"id"`
	Format *string `json:"format,omitempty"`
}

// Storm is the typed data model for the storm entity.
type Storm struct {
}

// StormLoadMatch is the typed request payload for Storm.LoadTyped.
type StormLoadMatch struct {
	Kind string `json:"kind"`
	Slug string `json:"slug"`
	Format *string `json:"format,omitempty"`
}

// WhiteChristma is the typed data model for the white_christma entity.
type WhiteChristma struct {
}

// WhiteChristmaLoadMatch is the typed request payload for WhiteChristma.LoadTyped.
type WhiteChristmaLoadMatch struct {
	Format *string `json:"format,omitempty"`
}

// Yesterday is the typed data model for the yesterday entity.
type Yesterday struct {
}

// YesterdayLoadMatch is the typed request payload for Yesterday.LoadTyped.
type YesterdayLoadMatch struct {
	Format *string `json:"format,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
