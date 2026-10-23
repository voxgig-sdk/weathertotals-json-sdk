# WeathertotalsJson SDK feature factory

from weathertotalsjson_sdk.feature.base_feature import WeathertotalsJsonBaseFeature
from weathertotalsjson_sdk.feature.ratelimit_feature import WeathertotalsJsonRatelimitFeature
from weathertotalsjson_sdk.feature.retry_feature import WeathertotalsJsonRetryFeature
from weathertotalsjson_sdk.feature.test_feature import WeathertotalsJsonTestFeature
from weathertotalsjson_sdk.feature.timeout_feature import WeathertotalsJsonTimeoutFeature


_FEATURES = {
    "base": lambda: WeathertotalsJsonBaseFeature(),
    "ratelimit": lambda: WeathertotalsJsonRatelimitFeature(),
    "retry": lambda: WeathertotalsJsonRetryFeature(),
    "test": lambda: WeathertotalsJsonTestFeature(),
    "timeout": lambda: WeathertotalsJsonTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
