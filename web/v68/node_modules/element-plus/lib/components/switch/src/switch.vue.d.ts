import { SwitchProps } from "./switch.js";
import * as vue from "vue";

//#region ../../packages/components/switch/src/switch.vue.d.ts
declare var __VLS_1: {}, __VLS_14: {}, __VLS_27: {}, __VLS_51: {}, __VLS_64: {}, __VLS_77: {};
type __VLS_Slots = {} & {
  inactive?: (props: typeof __VLS_1) => any;
} & {
  inactive?: (props: typeof __VLS_14) => any;
} & {
  active?: (props: typeof __VLS_27) => any;
} & {
  'active-action'?: (props: typeof __VLS_51) => any;
} & {
  'inactive-action'?: (props: typeof __VLS_64) => any;
} & {
  active?: (props: typeof __VLS_77) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<SwitchProps>, {
  modelValue: boolean;
  disabled: undefined;
  activeText: string;
  inactiveText: string;
  activeValue: boolean;
  inactiveValue: boolean;
  name: string;
  validateEvent: boolean;
  width: string;
}>, {
  /**
   *  @description manual focus to the switch component
   **/
  focus: () => void;
  /**
   * @description whether Switch is checked
   */
  checked: vue.ComputedRef<boolean>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (val: string | number | boolean) => void;
  "update:modelValue": (val: string | number | boolean) => void;
  input: (val: string | number | boolean) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<SwitchProps>, {
  modelValue: boolean;
  disabled: undefined;
  activeText: string;
  inactiveText: string;
  activeValue: boolean;
  inactiveValue: boolean;
  name: string;
  validateEvent: boolean;
  width: string;
}>>> & {
  onChange?: ((val: string | number | boolean) => any) | undefined;
  onInput?: ((val: string | number | boolean) => any) | undefined;
  "onUpdate:modelValue"?: ((val: string | number | boolean) => any) | undefined;
}, {
  modelValue: boolean | string | number;
  disabled: boolean;
  validateEvent: boolean;
  name: string;
  width: string | number;
  activeText: string;
  inactiveText: string;
  activeValue: boolean | string | number;
  inactiveValue: boolean | string | number;
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