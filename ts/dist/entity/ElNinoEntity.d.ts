import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { ElNino, ElNinoLoadMatch } from '../WeathertotalsJsonTypes';
declare class ElNinoEntity extends WeathertotalsJsonEntityBase<ElNino> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: ElNinoEntity): ElNinoEntity;
    load(this: any, reqmatch?: ElNinoLoadMatch, ctrl?: Control): Promise<ElNinoEntity>;
}
export { ElNinoEntity };
