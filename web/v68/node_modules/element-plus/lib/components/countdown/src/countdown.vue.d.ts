import { CountdownProps } from "./countdown.js";
import * as vue from "vue";
import * as dayjs$1 from "dayjs";

//#region ../../packages/components/countdown/src/countdown.vue.d.ts
declare var __VLS_10: string, __VLS_11: {};
type __VLS_Slots = {} & { [K in NonNullable<typeof __VLS_10>]?: (props: typeof __VLS_11) => any };
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CountdownProps>, {
  format: string;
  value: number;
  valueStyle: undefined;
}>, {
  /**
   * @description current display value
   */
  displayValue: vue.ComputedRef<string>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: number) => void;
  finish: () => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CountdownProps>, {
  format: string;
  value: number;
  valueStyle: undefined;
}>>> & {
  onChange?: ((value: number) => any) | undefined;
  onFinish?: (() => any) | undefined;
}, {
  value: number | dayjs$1.Dayjs;
  format: string;
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