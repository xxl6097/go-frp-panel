import { CheckboxProps, CheckboxValueType } from "./checkbox.js";
import * as vue from "vue";

//#region ../../packages/components/checkbox/src/checkbox.vue.d.ts
declare var __VLS_10: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_10) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CheckboxProps>, {
  readonly modelValue: undefined;
  readonly label: undefined;
  readonly value: undefined;
  readonly disabled: undefined;
  readonly name: undefined;
  readonly trueValue: undefined;
  readonly falseValue: undefined;
  readonly trueLabel: undefined;
  readonly falseLabel: undefined;
  readonly id: undefined;
  readonly validateEvent: true;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (val: CheckboxValueType) => void;
  "update:modelValue": (val: CheckboxValueType) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CheckboxProps>, {
  readonly modelValue: undefined;
  readonly label: undefined;
  readonly value: undefined;
  readonly disabled: undefined;
  readonly name: undefined;
  readonly trueValue: undefined;
  readonly falseValue: undefined;
  readonly trueLabel: undefined;
  readonly falseLabel: undefined;
  readonly id: undefined;
  readonly validateEvent: true;
}>>> & {
  onChange?: ((val: CheckboxValueType) => any) | undefined;
  "onUpdate:modelValue"?: ((val: CheckboxValueType) => any) | undefined;
}, {
  value: string | boolean | number | object;
  modelValue: number | string | boolean;
  id: string;
  disabled: boolean;
  validateEvent: boolean;
  name: string;
  label: string | boolean | number | object;
  trueValue: string | number;
  falseValue: string | number;
  trueLabel: string | number;
  falseLabel: string | number;
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