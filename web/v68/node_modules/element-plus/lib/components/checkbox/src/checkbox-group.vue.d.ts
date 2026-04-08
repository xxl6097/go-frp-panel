import { CheckboxValueType } from "./checkbox.js";
import { CheckboxGroupProps, CheckboxGroupValueType } from "./checkbox-group.js";
import "../../../index.js";
import * as vue from "vue";

//#region ../../packages/components/checkbox/src/checkbox-group.vue.d.ts
declare var __VLS_8: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_8) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CheckboxGroupProps>, {
  modelValue: () => never[];
  disabled: undefined;
  tag: string;
  validateEvent: boolean;
  props: () => Required<{
    value?: string;
    label?: string;
    disabled?: string;
  }>;
  type: string;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (val: CheckboxValueType[]) => void;
  "update:modelValue": (val: CheckboxGroupValueType) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CheckboxGroupProps>, {
  modelValue: () => never[];
  disabled: undefined;
  tag: string;
  validateEvent: boolean;
  props: () => Required<{
    value?: string;
    label?: string;
    disabled?: string;
  }>;
  type: string;
}>>> & {
  onChange?: ((val: CheckboxValueType[]) => any) | undefined;
  "onUpdate:modelValue"?: ((val: CheckboxGroupValueType) => any) | undefined;
}, {
  props: {
    value?: string;
    label?: string;
    disabled?: string;
  };
  type: "checkbox" | "button";
  modelValue: CheckboxGroupValueType;
  disabled: boolean;
  validateEvent: boolean;
  tag: string;
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