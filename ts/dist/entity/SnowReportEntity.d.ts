import { WeathertotalsJsonEntityBase } from '../WeathertotalsJsonEntityBase';
import type { WeathertotalsJsonSDK } from '../WeathertotalsJsonSDK';
import type { Control } from '../types';
import type { SnowReport, SnowReportLoadMatch } from '../WeathertotalsJsonTypes';
declare class SnowReportEntity extends WeathertotalsJsonEntityBase<SnowReport> {
    constructor(client: WeathertotalsJsonSDK, entopts: any);
    make(this: SnowReportEntity): SnowReportEntity;
    load(this: any, reqmatch?: SnowReportLoadMatch, ctrl?: Control): Promise<SnowReportEntity>;
}
export { SnowReportEntity };
