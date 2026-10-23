# WeathertotalsJson SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module WeathertotalsJsonFeatures
  def self.make_feature(name)
    case name
    when "base"
      WeathertotalsJsonBaseFeature.new
    when "ratelimit"
      WeathertotalsJsonRatelimitFeature.new
    when "retry"
      WeathertotalsJsonRetryFeature.new
    when "test"
      WeathertotalsJsonTestFeature.new
    when "timeout"
      WeathertotalsJsonTimeoutFeature.new
    else
      WeathertotalsJsonBaseFeature.new
    end
  end
end
