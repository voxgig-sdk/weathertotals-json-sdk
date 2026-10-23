import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { Rain, RainLoadMatch } from '../WeathertotalsJsonTypes';
declare class RainEntity extends WeathertotalsJsonEntityBase<Rain> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: RainEntity): RainEntity;
    load(this: any, reqmatch?: RainLoadMatch, ctrl?: Control): Promise<RainEntity>;
}
export { RainEntity };
