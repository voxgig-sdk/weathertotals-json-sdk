

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { WeathertotalsJsonSDK, BaseFeature, config, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


loadEnvLocal(__dirname + '/../../../.env.local')


describe('SnowEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when WEATHERTOTALS_JSON_TEST_LIVE=TRUE.
  afterEach(liveDelay('WEATHERTOTALS_JSON_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = WeathertotalsJsonSDK.test()
    const ent = testsdk.Snow()
    assert(null != ent)
  })


  test('validate', async (t) => {
    if (null == (config as any).feature?.validate) {
      t.skip('feature not present in this SDK: validate')
      return
    }
    const client = WeathertotalsJsonSDK.test(undefined, { feature: { validate: { active: true } } })
    await assert.rejects(client.Snow().load({"city":1,"state":"x"} as any),
      (err: any) => 'validate_failed' === err.code)
  })



  test('basic', async (t) => {

    const live = 'TRUE' === process.env.WEATHERTOTALS_JSON_TEST_LIVE
    for (const op of []) {
      if (!live && maybeSkipControl(t, 'entityOp', 'snow.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"api":{"a":true,"h":"Api","n":"api","r":true,"t":"`$STRING`","key$":"api","index$":0},"asOf":{"a":true,"fo":"date-time","h":"As Of","n":"asOf","r":true,"sh":"When the response was assembled, or the observation time where one applies.","t":"`$STRING`","key$":"asOf","index$":1},"attribution":{"a":true,"h":"Attribution","n":"attribution","r":true,"t":"`$OBJECT`","key$":"attribution","index$":2},"data":{"a":true,"h":"Data","n":"data","r":true,"t":"`$OBJECT`","key$":"data","index$":3},"docs":{"a":true,"fo":"uri","h":"Docs","n":"docs","r":true,"t":"`$STRING`","key$":"docs","index$":4},"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":5},"page":{"a":true,"fo":"uri","h":"Page","n":"page","r":true,"t":"`$STRING`","key$":"page","index$":6}},"id":{"field":"id","name":"id","parts":["state","city"],"sep":"/"},"name":"snow","op":{"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /api/v1/snow/{state}/{city}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"city","or":"city","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"state","or":"state","r":true,"t":"`$STRING`","index$":1}],"query":[{"a":true,"ex":"json","k":"query","n":"format","or":"format","r":false,"t":"`$STRING`","index$":0},{"a":true,"ex":"days","k":"query","n":"table","or":"table","r":false,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/api/v1/snow/{state}/{city}","q":{"exist":["city","state"]},"r":{},"rs":{"alternatives":[{"kind":"raw","media":"text/csv"}],"kind":"json","media":"application/json"},"s":[{"lit":"api"},{"lit":"v1"},{"lit":"snow"},{"var":"state"},{"var":"city"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"snow","name__orig":"snow","Name":"Snow","name_":"snow","name-":"snow","NAME":"SNOW","index$":4}, {"active":true,"entity":"snow","key$":"BasicSnowFlow","kind":"basic","name":"BasicSnowFlow","param":{},"step":[{"a":false,"d":{},"i":{"ref":"snow_ref01","srcdatavar":"snow_ref01_data","suffix":"_dt0"},"m":{"id":"snow01","state":"state01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-snow_ref01"}}],"unreachable":true}]}, 'Snow', {"GET /api/v1/snow/{state}/{city}":{"protocol":"http","parameters":[{"name":"state","in":"path","required":true,"description":"State slug as it appears in the page URL, lower case and hyphenated.","schema":{"type":"string","pattern":"^[a-z-]+$","examples":["massachusetts"]},"index$":0},{"name":"city","in":"path","required":true,"description":"City slug as it appears in the page URL.","schema":{"type":"string","pattern":"^[a-z-]+$","examples":["boston"]},"index$":1},{"name":"format","in":"query","required":false,"description":"csv returns one flat table as text/csv, with comment lines above the header carrying the source, the public domain notice and the page the figures mirror.","schema":{"type":"string","enum":["json","csv"],"default":"json"},"index$":2},{"name":"table","in":"query","required":false,"description":"With format=csv, which table to return.","schema":{"type":"string","enum":["days","seasons","records","thresholds"],"default":"days"},"index$":3}]}}, { strict: LIVE_STRICT, t })
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let snow_ref01_data = Object.values(setup.data.existing.snow)[0] as any

  })
})



// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const LIVE_STRICT = true

function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/snow/SnowTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = WeathertotalsJsonSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['snow01','snow02','snow03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'WEATHERTOTALS_JSON_TEST_SNOW_ENTID': idmap,
    'WEATHERTOTALS_JSON_TEST_LIVE': 'FALSE',
    'WEATHERTOTALS_JSON_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['WEATHERTOTALS_JSON_TEST_SNOW_ENTID']

  const live = 'TRUE' === env.WEATHERTOTALS_JSON_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['WEATHERTOTALS_JSON_TEST_SNOW_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new WeathertotalsJsonSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.WEATHERTOTALS_JSON_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
