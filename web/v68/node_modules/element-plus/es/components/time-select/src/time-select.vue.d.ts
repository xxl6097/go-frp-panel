import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { PopperEffect } from "../../popper/src/popper.js";
import { TimeSelectProps } from "./time-select.js";
import "../../../index.js";
import * as vue from "vue";

//#region ../../packages/components/time-select/src/time-select.vue.d.ts
declare const __VLS_export: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<TimeSelectProps>, {
  format: string;
  disabled: undefined;
  editable: boolean;
  effect: string;
  clearable: boolean;
  start: string;
  end: string;
  step: string;
  prefixIcon: () => any;
  clearIcon: () => any;
  popperClass: string;
  valueOnClear: undefined;
  popperStyle: undefined;
}>, {
  /**
   * @description blur the Input component
   */
  blur: () => void;
  /**
   * @description focus the Input component
   */
  focus: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (...args: any[]) => void;
  "update:modelValue": (...args: any[]) => void;
  focus: (...args: any[]) => void;
  blur: (...args: any[]) => void;
  clear: (...args: any[]) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<TimeSelectProps>, {
  format: string;
  disabled: undefined;
  editable: boolean;
  effect: string;
  clearable: boolean;
  start: string;
  end: string;
  step: string;
  prefixIcon: () => any;
  clearIcon: () => any;
  popperClass: string;
  valueOnClear: undefined;
  popperStyle: undefined;
}>>> & {
  onChange?: ((...args: any[]) => any) | undefined;
  onFocus?: ((...args: any[]) => any) | undefined;
  onBlur?: ((...args: any[]) => any) | undefined;
  "onUpdate:modelValue"?: ((...args: any[]) => any) | undefined;
  onClear?: ((...args: any[]) => any) | undefined;
}, {
  effect: PopperEffect;
  popperClass: string;
  popperStyle: string | vue.CSSProperties;
  disabled: boolean;
  clearable: boolean;
  clearIcon: IconPropType;
  prefixIcon: IconPropType;
  end: string;
  start: string;
  valueOnClear: string | number | boolean | Function | null;
  editable: boolean;
  format: string;
  step: string;
}, {}>;
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
type __VLS_PrettifyLocal<T> = (T extends any ? { [K in keyof T]: T[K] } : { [K in keyof T as K]: T[K] }) & {};
//#endregion
export { _default };