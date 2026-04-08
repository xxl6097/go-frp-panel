import { RadioGroupProps, radioOptionProp } from "./radio-group.js";
import * as vue from "vue";

//#region ../../packages/components/radio/src/radio-group.vue.d.ts
declare var __VLS_1: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_1) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<RadioGroupProps>, {
  readonly id: undefined;
  readonly disabled: undefined;
  readonly modelValue: undefined;
  readonly fill: "";
  readonly textColor: "";
  readonly name: undefined;
  readonly validateEvent: true;
  readonly props: () => Required<radioOptionProp>;
  readonly type: "radio";
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (val: string | number | boolean | undefined) => void;
  "update:modelValue": (val: string | number | boolean | undefined) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<RadioGroupProps>, {
  readonly id: undefined;
  readonly disabled: undefined;
  readonly modelValue: undefined;
  readonly fill: "";
  readonly textColor: "";
  readonly name: undefined;
  readonly validateEvent: true;
  readonly props: () => Required<radioOptionProp>;
  readonly type: "radio";
}>>> & {
  onChange?: ((val: string | number | boolean | undefined) => any) | undefined;
  "onUpdate:modelValue"?: ((val: string | number | boolean | undefined) => any) | undefined;
}, {
  fill: string;
  props: radioOptionProp;
  type: "radio" | "button";
  modelValue: string | number | boolean;
  id: string;
  disabled: boolean;
  validateEvent: boolean;
  name: string;
  textColor: string;
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