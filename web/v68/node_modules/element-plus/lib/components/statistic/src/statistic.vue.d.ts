import { StatisticProps } from "./statistic.js";
import * as vue from "vue";
import * as dayjs$1 from "dayjs";

//#region ../../packages/components/statistic/src/statistic.vue.d.ts
declare var __VLS_1: {}, __VLS_3: {}, __VLS_5: {};
type __VLS_Slots = {} & {
  title?: (props: typeof __VLS_1) => any;
} & {
  prefix?: (props: typeof __VLS_3) => any;
} & {
  suffix?: (props: typeof __VLS_5) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<StatisticProps>, {
  decimalSeparator: string;
  groupSeparator: string;
  precision: number;
  value: number;
  valueStyle: undefined;
}>, {
  /**
   * @description current display value
   */
  displayValue: vue.ComputedRef<string | number | dayjs$1.Dayjs>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<StatisticProps>, {
  decimalSeparator: string;
  groupSeparator: string;
  precision: number;
  value: number;
  valueStyle: undefined;
}>>>, {
  value: number | dayjs$1.Dayjs;
  precision: number;
  decimalSeparator: string;
  groupSeparator: string;
  valueStyle: string | false | vue.CSSProperties | vue.StyleValue[] | null;
}, {}>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
type __VLS_TypePropsToOption<T> = { [K in keyof T]-?: {} extends Pick<T, K> ? {
  type: vue.PropType<Required<T>[K]>;
} : {
  type: vue.PropType<T[K]>;
  required: true;
} };
type __VLS_WithDefaults<P, D> = { [K in keyof Pick<P, keyof P>]: K extends keyof D ? __VLS_PrettifyLocal<P[K] & {
  default: D[K];
}> : P[K] };
type __VLS_WithSlots<T, S> = T & {
  new (): {
    $slots: S;
  };
};
type __VLS_PrettifyLocal<T> = (T extends any ? { [K in keyof T]: T[K] } : { [K in keyof T as K]: T[K] }) & {};
//#endregion
export { _default };