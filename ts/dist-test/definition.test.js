"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const node_test_1 = require("node:test");
const __1 = require("..");
const definition_runner_1 = require("./definition-runner");
const utility_1 = require("./utility");
// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN = [
    {
        "entity": "aurora_now",
        "accessor": "AuroraNow",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/aurora/now",
        "args": [],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "kp": 1,
                "observedAt": "2026-01-01T00:00:00Z",
                "forecastPeak": 1,
                "states": [
                    {
                        "state": "x",
                        "abbr": "x",
                        "kpNeeded": 1,
                        "verdict": "No chance",
                        "severity": "none",
                        "score": 1,
                        "page": "x"
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "el_nino",
        "accessor": "ElNino",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/el-nino/{state}",
        "args": [
            {
                "name": "id",
                "wire": "state",
                "value": "p1"
            }
        ],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "state": "x",
                "abbr": "x",
                "units": {},
                "strongWinters": [
                    "x"
                ],
                "medians": {
                    "precipitationPercentChange": 1,
                    "snowfallPercentChange": 1,
                    "temperatureDifference": 1
                },
                "cities": [
                    {
                        "city": "x",
                        "slug": "x",
                        "station": "x",
                        "precipitation": {
                            "strongMean": 1,
                            "otherMean": 1,
                            "difference": 1,
                            "percentChange": 1,
                            "strongWinters": [],
                            "strongAboveOther": 1,
                            "otherWinters": 1
                        },
                        "snowfall": {
                            "strongMean": 1,
                            "otherMean": 1,
                            "difference": 1,
                            "percentChange": 1,
                            "strongWinters": [],
                            "strongAboveOther": 1,
                            "otherWinters": 1
                        },
                        "temperature": {
                            "strongMean": 1,
                            "otherMean": 1,
                            "difference": 1,
                            "percentChange": 1,
                            "strongWinters": [],
                            "strongAboveOther": 1,
                            "otherWinters": 1
                        },
                        "whiteChristmas": {
                            "strong": 1,
                            "strongYears": 1,
                            "other": 1,
                            "otherYears": 1
                        }
                    }
                ],
                "pageUrl": "x"
            }
        },
        "idField": "id"
    },
    {
        "entity": "hail_week",
        "accessor": "HailWeek",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/hail/week",
        "args": [],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "window": "x",
                "unit": "x",
                "total": 1,
                "largest": {
                    "inches": 1,
                    "place": "x",
                    "county": "x",
                    "state": "x",
                    "source": "x",
                    "measured": true,
                    "at": "2026-01-01T00:00:00Z",
                    "lat": 1,
                    "lon": 1
                },
                "last24h": {
                    "count": 1,
                    "top": {
                        "inches": 1,
                        "place": "x",
                        "county": "x",
                        "state": "x",
                        "source": "x",
                        "measured": true,
                        "at": "2026-01-01T00:00:00Z",
                        "lat": 1,
                        "lon": 1
                    }
                },
                "states": [
                    {
                        "state": "x",
                        "abbr": "x",
                        "reports": 1,
                        "largest": 1,
                        "page": "x"
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "rain",
        "accessor": "Rain",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/rain/{state}/{city}",
        "args": [
            {
                "name": "city",
                "wire": "city",
                "value": "p1"
            },
            {
                "name": "state",
                "wire": "state",
                "value": "p2"
            }
        ],
        "select": {
            "format": "json",
            "table": "days"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format",
            "table"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            },
            {
                "name": "table",
                "wire": "table"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "city": "x",
                "state": "x",
                "abbr": "x",
                "station": {
                    "id": "x",
                    "name": "x",
                    "cli": "x"
                },
                "unit": "x",
                "official": {
                    "date": "2026-01-01",
                    "precip": 1,
                    "trace": true,
                    "precipNormal": 1,
                    "monthToDate": 1,
                    "monthNormal": 1,
                    "yearToDate": 1,
                    "yearNormal": 1,
                    "record": 1,
                    "recordYears": [
                        1
                    ],
                    "name": "x",
                    "last2Days": 1,
                    "last3Days": 1,
                    "last7Days": 1
                },
                "year": {
                    "year": 1,
                    "total": 1,
                    "normalToDate": 1,
                    "rainDays": 1,
                    "wettest": {
                        "date": "2026-01-01",
                        "inches": 1
                    },
                    "latest": {
                        "date": "2026-01-01",
                        "inches": 1
                    },
                    "through": "2026-01-01"
                },
                "records": {
                    "firstYear": 1,
                    "wettestDay": {
                        "date": "2026-01-01",
                        "inches": 1
                    },
                    "wettestYear": {
                        "year": 1,
                        "total": 1,
                        "days": 1
                    },
                    "driestYear": {
                        "year": 1,
                        "total": 1,
                        "days": 1
                    },
                    "averageYear": 1,
                    "averageRainDays": 1,
                    "monthlyNormals": [
                        1
                    ]
                }
            }
        },
        "idField": "id"
    },
    {
        "entity": "snow",
        "accessor": "Snow",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/snow/{state}/{city}",
        "args": [
            {
                "name": "city",
                "wire": "city",
                "value": "p1"
            },
            {
                "name": "state",
                "wire": "state",
                "value": "p2"
            }
        ],
        "select": {
            "format": "json",
            "table": "days"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format",
            "table"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            },
            {
                "name": "table",
                "wire": "table"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "city": "x",
                "state": "x",
                "abbr": "x",
                "station": {
                    "id": "x",
                    "name": "x"
                },
                "unit": "x",
                "season": {
                    "label": "x",
                    "total": 1,
                    "normalToDate": 1,
                    "snowDays": 1,
                    "biggestDay": {
                        "date": "2026-01-01",
                        "inches": 1
                    },
                    "latest": {
                        "date": "2026-01-01",
                        "inches": 1
                    },
                    "days": [
                        {
                            "date": "2026-01-01",
                            "inches": 1
                        }
                    ]
                },
                "records": {
                    "firstYear": 1,
                    "biggestDay": {
                        "date": "2026-01-01",
                        "inches": 1
                    },
                    "biggestDays": [
                        {
                            "date": "2026-01-01",
                            "inches": 1
                        }
                    ],
                    "snowiestSeason": {
                        "season": 1,
                        "total": 1,
                        "days": 1
                    },
                    "leastSnowySeason": {
                        "season": 1,
                        "total": 1,
                        "days": 1
                    },
                    "averageSeason": 1,
                    "completeSeasons": 1,
                    "monthlyNormals": [
                        1
                    ]
                }
            }
        },
        "idField": "id"
    },
    {
        "entity": "snow_report",
        "accessor": "SnowReport",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/snow/reports",
        "args": [],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "window": "x",
                "unit": "x",
                "count": 1,
                "highest": {
                    "inches": 1,
                    "place": "x",
                    "county": "x",
                    "state": "x",
                    "source": "x",
                    "measured": true,
                    "at": "2026-01-01T00:00:00Z",
                    "lat": 1,
                    "lon": 1
                },
                "reports": [
                    {
                        "inches": 1,
                        "place": "x",
                        "county": "x",
                        "state": "x",
                        "source": "x",
                        "measured": true,
                        "at": "2026-01-01T00:00:00Z",
                        "lat": 1,
                        "lon": 1
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "snow_season",
        "accessor": "SnowSeason",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/snow/season",
        "args": [],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "season": "x",
                "unit": "x",
                "states": [
                    {
                        "state": "x",
                        "abbr": "x",
                        "measuredAt": "x",
                        "seasonToDate": 1,
                        "normalToDate": 1,
                        "page": "x",
                        "cityPage": "x"
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "snowpack",
        "accessor": "Snowpack",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/snowpack/{state}",
        "args": [
            {
                "name": "id",
                "wire": "state",
                "value": "p1"
            }
        ],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "state": "x",
                "abbr": "x",
                "units": {},
                "indexDate": "x",
                "stations": [
                    {
                        "station": "x",
                        "triplet": "x",
                        "elevationFeet": 1,
                        "lat": 1,
                        "lon": 1,
                        "date": "x",
                        "swe": 1,
                        "sweMedian": 1,
                        "percentOfMedian": 1,
                        "depth": 1,
                        "change7": 1,
                        "precip": 1,
                        "precipMedian": 1,
                        "page": "x"
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "storm",
        "accessor": "Storm",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/storms/{kind}/{slug}",
        "args": [
            {
                "name": "kind",
                "wire": "kind",
                "value": "p1"
            },
            {
                "name": "slug",
                "wire": "slug",
                "value": "p2"
            }
        ],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "kind": "snow",
                "slug": "x",
                "startedAt": "2026-01-01T00:00:00Z",
                "endedAt": "2026-01-01T00:00:00Z",
                "states": [
                    "x"
                ],
                "unit": "x",
                "reportCount": 1,
                "top": {
                    "inches": 1,
                    "place": "x",
                    "county": "x",
                    "state": "x",
                    "at": "2026-01-01T00:00:00Z"
                },
                "reports": [
                    {
                        "inches": 1,
                        "place": "x",
                        "county": "x",
                        "state": "x",
                        "source": "x",
                        "measured": true,
                        "at": "2026-01-01T00:00:00Z",
                        "lat": 1,
                        "lon": 1,
                        "product": "x",
                        "remark": "x"
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "white_christma",
        "accessor": "WhiteChristma",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/white-christmas",
        "args": [],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "snowDepthInchesAtLeast": 1,
                "date": "x",
                "since": 1,
                "units": {},
                "cities": [
                    {
                        "city": "x",
                        "state": "x",
                        "abbr": "x",
                        "share": 1,
                        "white": 1,
                        "reported": 1,
                        "last": 1,
                        "recentShare": 1,
                        "strongElNino": {
                            "white": 1,
                            "reported": 1
                        },
                        "lat": 1,
                        "lon": 1,
                        "page": "x"
                    }
                ]
            }
        },
        "idField": "id"
    },
    {
        "entity": "yesterday",
        "accessor": "Yesterday",
        "op": "load",
        "method": "GET",
        "path": "/api/v1/yesterday",
        "args": [],
        "select": {
            "format": "json"
        },
        "headers": [],
        "cookies": [],
        "responseMedia": [
            "application/json",
            "text/csv"
        ],
        "query": [
            "format"
        ],
        "queryArgs": [
            {
                "name": "format",
                "wire": "format"
            }
        ],
        "auth": null,
        "status": 200,
        "sample": {
            "api": "x",
            "asOf": "2026-01-01T00:00:00Z",
            "page": "x",
            "docs": "x",
            "attribution": {
                "site": "x",
                "license": {
                    "name": "x",
                    "spdx": "x",
                    "url": "x"
                },
                "notice": "x",
                "sources": [
                    "x"
                ],
                "required": "x"
            },
            "data": {
                "date": "2026-01-01",
                "units": {
                    "temperature": "x",
                    "precip": "x"
                },
                "cities": [
                    {
                        "city": "x",
                        "state": "x",
                        "abbr": "x",
                        "station": "x",
                        "page": "x",
                        "report": {
                            "date": "2026-01-01",
                            "high": 1,
                            "highTime": "x",
                            "highNormal": 1,
                            "highDepart": 1,
                            "highRecord": 1,
                            "highRecordYears": [],
                            "low": 1,
                            "lowTime": "x",
                            "lowNormal": 1,
                            "lowDepart": 1,
                            "lowRecord": 1,
                            "lowRecordYears": [],
                            "precip": 1,
                            "trace": true
                        }
                    }
                ]
            }
        },
        "idField": "id"
    }
];
(0, node_test_1.describe)('definition', () => {
    for (const point of PLAN) {
        (0, node_test_1.test)(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
            const control = (0, utility_1.isControlSkipped)('entityOp', point.entity + '.' + point.op, 'definition');
            if (control.skip) {
                t.skip(control.reason || 'skipped via sdk-test-control.json');
                return;
            }
            await (0, definition_runner_1.runDefinitionPoint)(__1.SDK, point);
        });
    }
});
//# sourceMappingURL=definition.test.js.map