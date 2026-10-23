# WeathertotalsJson SDK utility: make_context

from projectname_sdk.core.context import WeathertotalsJsonContext


def make_context_util(ctxmap, basectx):
    return WeathertotalsJsonContext(ctxmap, basectx)
