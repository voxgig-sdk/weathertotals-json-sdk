import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { Storm, StormLoadMatch } from '../WeathertotalsJsonTypes';
declare class StormEntity extends WeathertotalsJsonEntityBase<Storm> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: StormEntity): StormEntity;
    load(this: any, reqmatch?: StormLoadMatch, ctrl?: Control): Promise<StormEntity>;
}
export { StormEntity };
