import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { AuroraNow, AuroraNowLoadMatch } from '../WeathertotalsJsonTypes';
declare class AuroraNowEntity extends WeathertotalsJsonEntityBase<AuroraNow> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: AuroraNowEntity): AuroraNowEntity;
    load(this: any, reqmatch?: AuroraNowLoadMatch, ctrl?: Control): Promise<AuroraNowEntity>;
}
export { AuroraNowEntity };
