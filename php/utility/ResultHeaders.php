<?php
declare(strict_types=1);

// WeathertotalsJson SDK utility: result_headers

class WeathertotalsJsonResultHeaders
{
    public static function call(WeathertotalsJsonContext $ctx): ?WeathertotalsJsonResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
