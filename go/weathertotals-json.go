package voxgigweathertotalsjsonsdk

import (
	"github.com/voxgig-sdk/weathertotals-json-sdk/go/core"
	"github.com/voxgig-sdk/weathertotals-json-sdk/go/entity"
	"github.com/voxgig-sdk/weathertotals-json-sdk/go/feature"
	_ "github.com/voxgig-sdk/weathertotals-json-sdk/go/utility"
)

// Type aliases preserve external API.
type WeathertotalsJsonSDK = core.WeathertotalsJsonSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type WeathertotalsJsonEntity = core.WeathertotalsJsonEntity
type StreamItem = core.StreamItem
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type WeathertotalsJsonError = core.WeathertotalsJsonError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewRatelimitFeatureFunc = func() core.Feature {
		return feature.NewRatelimitFeature()
	}
	core.NewRetryFeatureFunc = func() core.Feature {
		return feature.NewRetryFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewTimeoutFeatureFunc = func() core.Feature {
		return feature.NewTimeoutFeature()
	}
	core.NewAuroraNowEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewAuroraNowEntity(client, entopts)
	}
	core.NewElNinoEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewElNinoEntity(client, entopts)
	}
	core.NewHailWeekEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewHailWeekEntity(client, entopts)
	}
	core.NewRainEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewRainEntity(client, entopts)
	}
	core.NewSnowEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewSnowEntity(client, entopts)
	}
	core.NewSnowReportEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewSnowReportEntity(client, entopts)
	}
	core.NewSnowSeasonEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewSnowSeasonEntity(client, entopts)
	}
	core.NewSnowpackEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewSnowpackEntity(client, entopts)
	}
	core.NewStormEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewStormEntity(client, entopts)
	}
	core.NewWhiteChristmaEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewWhiteChristmaEntity(client, entopts)
	}
	core.NewYesterdayEntityFunc = func(client *core.WeathertotalsJsonSDK, entopts map[string]any) core.WeathertotalsJsonEntity {
		return entity.NewYesterdayEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewWeathertotalsJsonSDK = core.NewWeathertotalsJsonSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewWeathertotalsJsonSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *WeathertotalsJsonSDK  { return NewWeathertotalsJsonSDK(nil) }
func Test() *WeathertotalsJsonSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
