import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { Snowpack, SnowpackLoadMatch } from '../WeathertotalsJsonTypes';
declare class SnowpackEntity extends WeathertotalsJsonEntityBase<Snowpack> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: SnowpackEntity): SnowpackEntity;
    load(this: any, reqmatch?: SnowpackLoadMatch, ctrl?: Control): Promise<SnowpackEntity>;
}
export { SnowpackEntity };
