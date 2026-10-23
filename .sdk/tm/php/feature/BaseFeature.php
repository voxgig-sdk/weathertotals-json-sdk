<?php
declare(strict_types=1);

// WeathertotalsJson SDK base feature

class WeathertotalsJsonBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(WeathertotalsJsonContext $ctx, array $options): void {}
    public function PostConstruct(WeathertotalsJsonContext $ctx): void {}
    public function PostConstructEntity(WeathertotalsJsonContext $ctx): void {}
    public function SetData(WeathertotalsJsonContext $ctx): void {}
    public function GetData(WeathertotalsJsonContext $ctx): void {}
    public function GetMatch(WeathertotalsJsonContext $ctx): void {}
    public function SetMatch(WeathertotalsJsonContext $ctx): void {}
    public function PrePoint(WeathertotalsJsonContext $ctx): void {}
    public function PreSpec(WeathertotalsJsonContext $ctx): void {}
    public function PreRequest(WeathertotalsJsonContext $ctx): void {}
    public function PreResponse(WeathertotalsJsonContext $ctx): void {}
    public function PreResult(WeathertotalsJsonContext $ctx): void {}
    public function PreDone(WeathertotalsJsonContext $ctx): void {}
    public function PreUnexpected(WeathertotalsJsonContext $ctx): void {}
}
