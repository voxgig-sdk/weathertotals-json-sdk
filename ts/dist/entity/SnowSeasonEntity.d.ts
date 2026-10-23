import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { SnowSeason, SnowSeasonLoadMatch } from '../WeathertotalsJsonTypes';
declare class SnowSeasonEntity extends WeathertotalsJsonEntityBase<SnowSeason> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: SnowSeasonEntity): SnowSeasonEntity;
    load(this: any, reqmatch?: SnowSeasonLoadMatch, ctrl?: Control): Promise<SnowSeasonEntity>;
}
export { SnowSeasonEntity };
