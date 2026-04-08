import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { RateProps } from "./rate.js";
import * as vue from "vue";

//#region ../../packages/components/rate/src/rate.vue.d.ts
declare function setCurrentValue(value: number, event?: MouseEvent): void;
declare function resetCurrentValue(): void;
declare const __VLS_export: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<RateProps>, {
  modelValue: number;
  id: undefined;
  lowThreshold: number;
  highThreshold: number;
  max: number;
  colors: () => string[];
  voidColor: string;
  disabledVoidColor: string;
  icons: () => any[];
  voidIcon: () => any;
  disabledVoidIcon: () => any;
  disabled: undefined;
  textColor: string;
  texts: () => string[];
  scoreTemplate: string;
}>, {
  /** @description set current value */setCurrentValue: typeof setCurrentValue; /** @description reset current value */
  resetCurrentValue: typeof resetCurrentValue;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: number) => void;
  "update:modelValue": (value: number) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<RateProps>, {
  modelValue: number;
  id: undefined;
  lowThreshold: number;
  highThreshold: number;
  max: number;
  colors: () => string[];
  voidColor: string;
  disabledVoidColor: string;
  icons: () => any[];
  voidIcon: () => any;
  disabledVoidIcon: () => any;
  disabled: undefined;
  textColor: string;
  texts: () => string[];
  scoreTemplate: string;
}>>> & {
  onChange?: ((value: number) => any) | undefined;
  "onUpdate:modelValue"?: ((value: number) => any) | undefined;
}, {
  modelValue: number;
  id: string;
  disabled: boolean;
  max: number;
  textColor: string;
  lowThreshold: number;
  highThreshold: number;
  colors: string[] | Record<number, string>;
  voidColor: string;
  disabledVoidColor: string;
  icons: Array<IconPropType> | Record<number, IconPropType>;
  voidIcon: IconPropType;
  disabledVoidIcon: IconPropType;
  texts: string[];
  scoreTemplate: string;
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