import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { HailWeek, HailWeekLoadMatch } from '../WeathertotalsJsonTypes';
declare class HailWeekEntity extends WeathertotalsJsonEntityBase<HailWeek> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: HailWeekEntity): HailWeekEntity;
    load(this: any, reqmatch?: HailWeekLoadMatch, ctrl?: Control): Promise<HailWeekEntity>;
}
export { HailWeekEntity };
