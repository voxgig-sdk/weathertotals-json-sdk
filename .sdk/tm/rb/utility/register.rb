# WeathertotalsJson SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

WeathertotalsJsonUtility.registrar = ->(u) {
  u.clean = WeathertotalsJsonUtilities::Clean
  u.clean_add = WeathertotalsJsonUtilities::CleanAdd
  u.clean_explain = WeathertotalsJsonUtilities::CleanExplain
  u.done = WeathertotalsJsonUtilities::Done
  u.make_error = WeathertotalsJsonUtilities::MakeError
  u.feature_add = WeathertotalsJsonUtilities::FeatureAdd
  u.feature_hook = WeathertotalsJsonUtilities::FeatureHook
  u.feature_init = WeathertotalsJsonUtilities::FeatureInit
  u.fetcher = WeathertotalsJsonUtilities::Fetcher
  u.make_fetch_def = WeathertotalsJsonUtilities::MakeFetchDef
  u.make_context = WeathertotalsJsonUtilities::MakeContext
  u.make_options = WeathertotalsJsonUtilities::MakeOptions
  u.make_request = WeathertotalsJsonUtilities::MakeRequest
  u.make_response = WeathertotalsJsonUtilities::MakeResponse
  u.make_result = WeathertotalsJsonUtilities::MakeResult
  u.make_point = WeathertotalsJsonUtilities::MakePoint
  u.make_spec = WeathertotalsJsonUtilities::MakeSpec
  u.make_url = WeathertotalsJsonUtilities::MakeUrl
  u.param = WeathertotalsJsonUtilities::Param
  u.prepare_auth = WeathertotalsJsonUtilities::PrepareAuth
  u.prepare_body = WeathertotalsJsonUtilities::PrepareBody
  u.prepare_headers = WeathertotalsJsonUtilities::PrepareHeaders
  u.prepare_method = WeathertotalsJsonUtilities::PrepareMethod
  u.prepare_params = WeathertotalsJsonUtilities::PrepareParams
  u.prepare_path = WeathertotalsJsonUtilities::PreparePath
  u.prepare_query = WeathertotalsJsonUtilities::PrepareQuery
  u.graphql_body = WeathertotalsJsonUtilities::GraphqlBody
  u.graphql_errors = WeathertotalsJsonUtilities::GraphqlErrors
  u.result_basic = WeathertotalsJsonUtilities::ResultBasic
  u.result_body = WeathertotalsJsonUtilities::ResultBody
  u.result_headers = WeathertotalsJsonUtilities::ResultHeaders
  u.transform_request = WeathertotalsJsonUtilities::TransformRequest
  u.transform_response = WeathertotalsJsonUtilities::TransformResponse
}
