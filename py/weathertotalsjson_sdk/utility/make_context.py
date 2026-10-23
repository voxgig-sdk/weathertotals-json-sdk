# WeathertotalsJson SDK utility: make_context

from weathertotalsjson_sdk.core.context import WeathertotalsJsonContext


def make_context_util(ctxmap, basectx):
    return WeathertotalsJsonContext(ctxmap, basectx)
