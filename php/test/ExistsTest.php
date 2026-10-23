<?php
declare(strict_types=1);

// WeathertotalsJson SDK exists test

require_once __DIR__ . '/../weathertotalsjson_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = WeathertotalsJsonSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
