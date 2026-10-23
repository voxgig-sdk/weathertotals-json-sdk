import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { Yesterday, YesterdayLoadMatch } from '../WeathertotalsJsonTypes';
declare class YesterdayEntity extends WeathertotalsJsonEntityBase<Yesterday> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: YesterdayEntity): YesterdayEntity;
    load(this: any, reqmatch?: YesterdayLoadMatch, ctrl?: Control): Promise<YesterdayEntity>;
}
export { YesterdayEntity };
