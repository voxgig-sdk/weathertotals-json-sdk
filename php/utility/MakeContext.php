<?php
declare(strict_types=1);

// WeathertotalsJson SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class WeathertotalsJsonMakeContext
{
    public static function call(array $ctxmap, ?WeathertotalsJsonContext $basectx): WeathertotalsJsonContext
    {
        return new WeathertotalsJsonContext($ctxmap, $basectx);
    }
}
