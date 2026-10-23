# WeathertotalsJson SDK exists test

import pytest
from weathertotalsjson_sdk import WeathertotalsJsonSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = WeathertotalsJsonSDK.test(None, None)
        assert testsdk is not None
