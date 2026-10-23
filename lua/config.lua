-- WeathertotalsJson SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "WeathertotalsJson",
      slug = "weathertotals-json",
      version = "0.0.1",
      target = "lua",
    },
    feature = {
      ["ratelimit"] = {
        ["options"] = {
          ["active"] = false,
          ["burst"] = 5,
          ["rate"] = 5,
        },
        ["optspec"] = {
          ["now"] = "`$FUNCTION`",
          ["sleep"] = "`$FUNCTION`",
        },
        ["strict"] = false,
        ["transport"] = "wrap",
      },
      ["retry"] = {
        ["options"] = {
          ["active"] = false,
          ["factor"] = 2,
          ["maxDelay"] = 2000,
          ["minDelay"] = 50,
          ["retries"] = 2,
          ["statuses"] = {
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          },
        },
        ["optspec"] = {
          ["jitter"] = "`$BOOLEAN`",
          ["sleep"] = "`$FUNCTION`",
        },
        ["strict"] = false,
        ["transport"] = "wrap",
      },
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
        ["optspec"] = {
          ["entity"] = "`$MAP`",
          ["net"] = "`$MAP`",
        },
        ["strict"] = false,
        ["transport"] = "base",
      },
      ["timeout"] = {
        ["options"] = {
          ["active"] = false,
          ["ms"] = 30000,
        },
        ["optspec"] = {
          ["clearTimer"] = "`$FUNCTION`",
          ["now"] = "`$FUNCTION`",
          ["setTimer"] = "`$FUNCTION`",
        },
        ["strict"] = false,
        ["transport"] = "wrap",
      },
    },
    options = {
      base = "https://weathertotals.com",
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["aurora_now"] = {},
        ["el_nino"] = {},
        ["hail_week"] = {},
        ["rain"] = {},
        ["snow"] = {},
        ["snow_report"] = {},
        ["snow_season"] = {},
        ["snowpack"] = {},
        ["storm"] = {},
        ["white_christma"] = {},
        ["yesterday"] = {},
      },
    },
    entity = {
      ["aurora_now"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["name"] = "aurora_now",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/aurora/now",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "aurora",
                  },
                  {
                    ["lit"] = "now",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "aurora",
                  "now",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {},
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["el_nino"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
        },
        ["name"] = "el_nino",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/el-nino/{state}",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "el-nino",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "el-nino",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["state"] = "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "id",
                      ["orig"] = "state",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["hail_week"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["name"] = "hail_week",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/hail/week",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "hail",
                  },
                  {
                    ["lit"] = "week",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "hail",
                  "week",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {},
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["rain"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
          ["parts"] = {
            "state",
            "city",
          },
          ["sep"] = "/",
        },
        ["name"] = "rain",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/rain/{state}/{city}",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "rain",
                  },
                  {
                    ["var"] = "state",
                  },
                  {
                    ["var"] = "city",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "rain",
                  "{state}",
                  "{city}",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "city",
                      ["orig"] = "city",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                    {
                      ["name"] = "state",
                      ["orig"] = "state",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                    {
                      ["name"] = "table",
                      ["orig"] = "table",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "days",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "city",
                    "state",
                  },
                },
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["snow"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
          ["parts"] = {
            "state",
            "city",
          },
          ["sep"] = "/",
        },
        ["name"] = "snow",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/snow/{state}/{city}",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "snow",
                  },
                  {
                    ["var"] = "state",
                  },
                  {
                    ["var"] = "city",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "snow",
                  "{state}",
                  "{city}",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "city",
                      ["orig"] = "city",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                    {
                      ["name"] = "state",
                      ["orig"] = "state",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                    {
                      ["name"] = "table",
                      ["orig"] = "table",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "days",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "city",
                    "state",
                  },
                },
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["snow_report"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["name"] = "snow_report",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/snow/reports",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "snow",
                  },
                  {
                    ["lit"] = "reports",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "snow",
                  "reports",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {},
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["snow_season"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["name"] = "snow_season",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/snow/season",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "snow",
                  },
                  {
                    ["lit"] = "season",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "snow",
                  "season",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {},
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["snowpack"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
        },
        ["name"] = "snowpack",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/snowpack/{state}",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "snowpack",
                  },
                  {
                    ["var"] = "id",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "snowpack",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["state"] = "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "id",
                      ["orig"] = "state",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["storm"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "id",
            ["title"] = "Id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["id"] = {
          ["field"] = "id",
          ["name"] = "id",
          ["parts"] = {
            "kind",
            "slug",
          },
          ["sep"] = "/",
        },
        ["name"] = "storm",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/storms/{kind}/{slug}",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "storms",
                  },
                  {
                    ["var"] = "kind",
                  },
                  {
                    ["var"] = "slug",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "storms",
                  "{kind}",
                  "{slug}",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["params"] = {
                    {
                      ["name"] = "kind",
                      ["orig"] = "kind",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                    {
                      ["name"] = "slug",
                      ["orig"] = "slug",
                      ["type"] = "`$STRING`",
                      ["kind"] = "param",
                      ["reqd"] = true,
                    },
                  },
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "kind",
                    "slug",
                  },
                },
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["white_christma"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["name"] = "white_christma",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/white-christmas",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "white-christmas",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "white-christmas",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {},
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["yesterday"] = {
        ["fields"] = {
          {
            ["name"] = "api",
            ["title"] = "Api",
            ["type"] = "`$STRING`",
            ["req"] = true,
          },
          {
            ["name"] = "asOf",
            ["title"] = "As Of",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["short"] = "When the response was assembled, or the observation time where one applies.",
            ["format"] = "date-time",
          },
          {
            ["name"] = "attribution",
            ["title"] = "Attribution",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "data",
            ["title"] = "Data",
            ["type"] = "`$OBJECT`",
            ["req"] = true,
          },
          {
            ["name"] = "docs",
            ["title"] = "Docs",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
          {
            ["name"] = "page",
            ["title"] = "Page",
            ["type"] = "`$STRING`",
            ["req"] = true,
            ["format"] = "uri",
          },
        },
        ["name"] = "yesterday",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/yesterday",
                ["segments"] = {
                  {
                    ["lit"] = "api",
                  },
                  {
                    ["lit"] = "v1",
                  },
                  {
                    ["lit"] = "yesterday",
                  },
                },
                ["parts"] = {
                  "api",
                  "v1",
                  "yesterday",
                },
                ["rename"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
                ["args"] = {
                  ["query"] = {
                    {
                      ["name"] = "format",
                      ["orig"] = "format",
                      ["type"] = "`$STRING`",
                      ["kind"] = "query",
                      ["example"] = "json",
                    },
                  },
                },
                ["select"] = {},
                ["response"] = {
                  ["alternatives"] = {
                    {
                      ["kind"] = "raw",
                      ["media"] = "text/csv",
                    },
                  },
                  ["kind"] = "json",
                  ["media"] = "application/json",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
