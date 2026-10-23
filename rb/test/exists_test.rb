# WeathertotalsJson SDK exists test

require "minitest/autorun"
require_relative "../WeathertotalsJson_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = WeathertotalsJsonSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
