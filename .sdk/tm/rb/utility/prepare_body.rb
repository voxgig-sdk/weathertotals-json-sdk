# WeathertotalsJson SDK utility: prepare_body
require_relative 'media'
module WeathertotalsJsonUtilities
  PrepareBody = ->(ctx) {
    return nil unless ctx.op.input == "data"
    return WeathertotalsJsonUtilities.raw_body(ctx.reqdata) if WeathertotalsJsonUtilities.raw_request?(ctx.point)
    ctx.utility.transform_request.call(ctx)
  }
end
