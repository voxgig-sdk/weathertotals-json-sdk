import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { WhiteChristma, WhiteChristmaLoadMatch } from '../WeathertotalsJsonTypes';
declare class WhiteChristmaEntity extends WeathertotalsJsonEntityBase<WhiteChristma> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: WhiteChristmaEntity): WhiteChristmaEntity;
    load(this: any, reqmatch?: WhiteChristmaLoadMatch, ctrl?: Control): Promise<WhiteChristmaEntity>;
}
export { WhiteChristmaEntity };
