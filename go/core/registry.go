package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewRatelimitFeatureFunc func() Feature

var NewRetryFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewTimeoutFeatureFunc func() Feature

var NewAuroraNowEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewElNinoEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewHailWeekEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewRainEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewSnowEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewSnowReportEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewSnowSeasonEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewSnowpackEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewStormEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewWhiteChristmaEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

var NewYesterdayEntityFunc func(client *WeathertotalsJsonSDK, entopts map[string]any) WeathertotalsJsonEntity

