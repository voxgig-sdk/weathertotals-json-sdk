import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { Snow, SnowLoadMatch } from '../WeathertotalsJsonTypes';
declare class SnowEntity extends WeathertotalsJsonEntityBase<Snow> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: SnowEntity): SnowEntity;
    load(this: any, reqmatch?: SnowLoadMatch, ctrl?: Control): Promise<SnowEntity>;
}
export { SnowEntity };
