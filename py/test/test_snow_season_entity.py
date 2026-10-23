# SnowSeason entity test

import json
import os
import time

import pytest

from weathertotalsjson_sdk.utility.voxgig_struct import voxgig_struct as vs
from weathertotalsjson_sdk import WeathertotalsJsonSDK
from weathertotalsjson_sdk.core import helpers
from weathertotalsjson_sdk.config import shared_config
from weathertotalsjson_sdk.feature.base_feature import WeathertotalsJsonBaseFeature

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner



# main.kit.test.live.strict is true (the default is true): a live
# request that fails, or a live test missing an input it needs,
# fails the test.
# An account with no record for a test to read skips it either way.
LIVE_STRICT = True


class TestSnowSeasonEntity:

    def test_should_create_instance(self):
        testsdk = WeathertotalsJsonSDK.test(None, None)
        ent = testsdk.SnowSeason(None)
        assert ent is not None

    def test_should_refuse_an_invalid_request(self):
        if "validate" not in (shared_config().get("feature") or {}):
            pytest.skip("feature not present in this SDK: validate")
        client = WeathertotalsJsonSDK.test(
            None, {"feature": {"validate": {"active": True}}})
        with pytest.raises(Exception) as err:
            client.SnowSeason(None).load({"format": 1}, None)
        assert "validate_failed" == getattr(err.value, "code", None)

    def test_should_run_basic_flow(self):
        setup = _snow_season_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "snow_season." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        snow_season_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.snow_season")))
        snow_season_ref01_data = None
        if len(snow_season_ref01_data_raw) > 0:
            snow_season_ref01_data = helpers.to_map(snow_season_ref01_data_raw[0][1])

        # LOAD
        snow_season_ref01_ent = client.SnowSeason(None)
        snow_season_ref01_match_dt0 = {}
        snow_season_ref01_data_dt0_loaded = snow_season_ref01_ent.load(snow_season_ref01_match_dt0, None)
        assert snow_season_ref01_data_dt0_loaded is not None



def _snow_season_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/snow_season/SnowSeasonTestData.json")
    with open(entity_data_file, "r", encoding="utf-8") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = WeathertotalsJsonSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["snow_season01", "snow_season02", "snow_season03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Whether *_ENTID supplied the idmap, read before env_override consumes
    # it: without it, the ids a live flow binds are the fixture's synthetic ones.
    _entid_env_raw = os.environ.get(
        "WEATHERTOTALS_JSON_TEST_SNOW_SEASON_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "WEATHERTOTALS_JSON_TEST_SNOW_SEASON_ENTID": idmap,
        "WEATHERTOTALS_JSON_TEST_LIVE": "FALSE",
        "WEATHERTOTALS_JSON_TEST_EXPLAIN": "FALSE",
    })

    idmap_resolved = helpers.to_map(
        env.get("WEATHERTOTALS_JSON_TEST_SNOW_SEASON_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("WEATHERTOTALS_JSON_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
            },
            extra or {},
        ])
        client = WeathertotalsJsonSDK(helpers.to_map(merged_opts))

    _live = env.get("WEATHERTOTALS_JSON_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("WEATHERTOTALS_JSON_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
