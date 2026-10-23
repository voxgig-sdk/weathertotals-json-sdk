-- WeathertotalsJson SDK exists test

local sdk = require("weathertotals-json_sdk")

describe("WeathertotalsJsonSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
