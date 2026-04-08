import { InputNumberProps } from "./input-number.js";
import * as vue from "vue";

//#region ../../packages/components/input-number/src/input-number.vue.d.ts
declare var __VLS_1: {}, __VLS_19: {}, __VLS_52: {}, __VLS_55: {};
type __VLS_Slots = {} & {
  'decrease-icon'?: (props: typeof __VLS_1) => any;
} & {
  'increase-icon'?: (props: typeof __VLS_19) => any;
} & {
  prefix?: (props: typeof __VLS_52) => any;
} & {
  suffix?: (props: typeof __VLS_55) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<InputNumberProps>, {
  id: undefined;
  disabled: undefined;
  step: number;
  max: number;
  min: number;
  stepStrictly: boolean;
  readonly: boolean;
  controls: boolean;
  controlsPosition: string;
  valueOnClear: null;
  validateEvent: boolean;
  inputmode: undefined;
  align: string;
}>, {
  /** @description get focus the input component */focus: () => void; /** @description remove focus the input component */
  blur: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (cur: number | undefined, prev: number | undefined) => void;
  "update:modelValue": (val: number | undefined) => void;
  input: (val: number | null | undefined) => void;
  focus: (e: FocusEvent) => void;
  blur: (e: FocusEvent) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<InputNumberProps>, {
  id: undefined;
  disabled: undefined;
  step: number;
  max: number;
  min: number;
  stepStrictly: boolean;
  readonly: boolean;
  controls: boolean;
  controlsPosition: string;
  valueOnClear: null;
  validateEvent: boolean;
  inputmode: undefined;
  align: string;
}>>> & {
  onChange?: ((cur: number | undefined, prev: number | undefined) => any) | undefined;
  onFocus?: ((e: FocusEvent) => any) | undefined;
  onBlur?: ((e: FocusEvent) => any) | undefined;
  onInput?: ((val: number | null | undefined) => any) | undefined;
  "onUpdate:modelValue"?: ((val: number | undefined) => any) | undefined;
}, {
  id: string;
  disabled: boolean;
  readonly: boolean;
  validateEvent: boolean;
  inputmode: "search" | "text" | "email" | "tel" | "url" | "none" | "numeric" | "decimal";
  max: number;
  valueOnClear: "min" | "max" | number | null;
  min: number;
  align: "left" | "right" | "center";
  step: number;
  stepStrictly: boolean;
  controls: boolean;
  controlsPosition: "" | "right";
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