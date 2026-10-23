# WeathertotalsJson SDK utility: make_context
require_relative '../core/context'
module WeathertotalsJsonUtilities
  MakeContext = ->(ctxmap, basectx) {
    WeathertotalsJsonContext.new(ctxmap, basectx)
  }
end
