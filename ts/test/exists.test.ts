
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { WeathertotalsJsonSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = WeathertotalsJsonSDK.test()
    equal(testsdk instanceof WeathertotalsJsonSDK, true,
      'WeathertotalsJsonSDK.test() must return a client synchronously')
  })

})
