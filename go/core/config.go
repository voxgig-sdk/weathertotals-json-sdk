package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "WeathertotalsJson",
			"slug": "weathertotals-json",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"now": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://weathertotals.com",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"aurora_now": map[string]any{},
				"el_nino": map[string]any{},
				"hail_week": map[string]any{},
				"rain": map[string]any{},
				"snow": map[string]any{},
				"snow_report": map[string]any{},
				"snow_season": map[string]any{},
				"snowpack": map[string]any{},
				"storm": map[string]any{},
				"white_christma": map[string]any{},
				"yesterday": map[string]any{},
			},
		},
		"entity": map[string]any{
			"aurora_now": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"name": "aurora_now",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/aurora/now",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "aurora",
									},
									map[string]any{
										"lit": "now",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"aurora",
									"now",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"el_nino": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "el_nino",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/el-nino/{state}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "el-nino",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"el-nino",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"state": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "state",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"hail_week": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"name": "hail_week",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/hail/week",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "hail",
									},
									map[string]any{
										"lit": "week",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"hail",
									"week",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"rain": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
					"parts": []any{
						"state",
						"city",
					},
					"sep": "/",
				},
				"name": "rain",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/rain/{state}/{city}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "rain",
									},
									map[string]any{
										"var": "state",
									},
									map[string]any{
										"var": "city",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"rain",
									"{state}",
									"{city}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
										map[string]any{
											"name": "table",
											"orig": "table",
											"type": "`$STRING`",
											"kind": "query",
											"example": "days",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"city",
										"state",
									},
								},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"snow": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
					"parts": []any{
						"state",
						"city",
					},
					"sep": "/",
				},
				"name": "snow",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/snow/{state}/{city}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "snow",
									},
									map[string]any{
										"var": "state",
									},
									map[string]any{
										"var": "city",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"snow",
									"{state}",
									"{city}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "city",
											"orig": "city",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
										map[string]any{
											"name": "table",
											"orig": "table",
											"type": "`$STRING`",
											"kind": "query",
											"example": "days",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"city",
										"state",
									},
								},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"snow_report": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"name": "snow_report",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/snow/reports",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "snow",
									},
									map[string]any{
										"lit": "reports",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"snow",
									"reports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"snow_season": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"name": "snow_season",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/snow/season",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "snow",
									},
									map[string]any{
										"lit": "season",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"snow",
									"season",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"snowpack": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "snowpack",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/snowpack/{state}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "snowpack",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"snowpack",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"state": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "state",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"storm": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
					"parts": []any{
						"kind",
						"slug",
					},
					"sep": "/",
				},
				"name": "storm",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/storms/{kind}/{slug}",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "storms",
									},
									map[string]any{
										"var": "kind",
									},
									map[string]any{
										"var": "slug",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"storms",
									"{kind}",
									"{slug}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "kind",
											"orig": "kind",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "slug",
											"orig": "slug",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"kind",
										"slug",
									},
								},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"white_christma": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"name": "white_christma",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/white-christmas",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "white-christmas",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"white-christmas",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"yesterday": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "api",
						"title": "Api",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "asOf",
						"title": "As Of",
						"type": "`$STRING`",
						"req": true,
						"short": "When the response was assembled, or the observation time where one applies.",
						"format": "date-time",
					},
					map[string]any{
						"name": "attribution",
						"title": "Attribution",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "docs",
						"title": "Docs",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"name": "yesterday",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/yesterday",
								"segments": []any{
									map[string]any{
										"lit": "api",
									},
									map[string]any{
										"lit": "v1",
									},
									map[string]any{
										"lit": "yesterday",
									},
								},
								"parts": []any{
									"api",
									"v1",
									"yesterday",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "json",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"alternatives": []any{
										map[string]any{
											"kind": "raw",
											"media": "text/csv",
										},
									},
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
