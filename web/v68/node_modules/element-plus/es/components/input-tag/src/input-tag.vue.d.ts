import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { PopperEffect } from "../../popper/src/popper.js";
import { InputTagProps } from "./input-tag.js";
import "../../../index.js";
import * as vue from "vue";

//#region ../../packages/components/input-tag/src/input-tag.vue.d.ts
declare var __VLS_1: {}, __VLS_15: {
    value: string;
    index: number;
  }, __VLS_41: {
    value: string;
    index: number;
  }, __VLS_43: {};
type __VLS_Slots = {} & {
  prefix?: (props: typeof __VLS_1) => any;
} & {
  tag?: (props: typeof __VLS_15) => any;
} & {
  tag?: (props: typeof __VLS_41) => any;
} & {
  suffix?: (props: typeof __VLS_43) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<InputTagProps>, {
  tagType: string;
  tagEffect: string;
  effect: string;
  trigger: string;
  delimiter: string;
  clearIcon: any;
  disabled: undefined;
  validateEvent: boolean;
  id: undefined;
  tabindex: number;
  autocomplete: string;
  saveOnBlur: boolean;
  maxCollapseTags: number;
}>, {
  focus: () => void;
  blur: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value?: string[] | undefined) => void;
  "update:modelValue": (value?: string[] | undefined) => void;
  input: (value: string) => void;
  focus: (evt: FocusEvent) => void;
  blur: (evt: FocusEvent) => void;
  clear: () => void;
  "add-tag": (value: string | string[]) => void;
  "remove-tag": (value: string, index: number) => void;
  "drag-tag": (oldIndex: number, newIndex: number, value: string) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<InputTagProps>, {
  tagType: string;
  tagEffect: string;
  effect: string;
  trigger: string;
  delimiter: string;
  clearIcon: any;
  disabled: undefined;
  validateEvent: boolean;
  id: undefined;
  tabindex: number;
  autocomplete: string;
  saveOnBlur: boolean;
  maxCollapseTags: number;
}>>> & {
  onChange?: ((value?: string[] | undefined) => any) | undefined;
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  onInput?: ((value: string) => any) | undefined;
  "onUpdate:modelValue"?: ((value?: string[] | undefined) => any) | undefined;
  onClear?: (() => any) | undefined;
  "onAdd-tag"?: ((value: string | string[]) => any) | undefined;
  "onRemove-tag"?: ((value: string, index: number) => any) | undefined;
  "onDrag-tag"?: ((oldIndex: number, newIndex: number, value: string) => any) | undefined;
}, {
  effect: PopperEffect;
  id: string;
  disabled: boolean;
  autocomplete: string;
  clearIcon: IconPropType;
  tabindex: string | number;
  validateEvent: boolean;
  trigger: "Enter" | "Space";
  maxCollapseTags: number;
  tagType: "info" | "primary" | "success" | "warning" | "danger";
  tagEffect: "light" | "dark" | "plain";
  saveOnBlur: boolean;
  delimiter: string | RegExp;
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