import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { InputModelModifiers, InputType } from "../../input/src/input.js";
import { InputInstance } from "../../input/src/instance.js";
import "../../input/index.js";
import { TooltipInstance } from "../../tooltip/src/tooltip.js";
import "../../tooltip/index.js";
import { MentionOption } from "./types.js";
import { filterOption } from "./helper.js";
import { MentionOptionProps, MentionProps } from "./mention.js";
import * as vue from "vue";
import { CSSProperties } from "vue";
import * as _popperjs_core0 from "@popperjs/core";

//#region ../../packages/components/mention/src/mention.vue.d.ts
declare var __VLS_15: string, __VLS_16: any, __VLS_41: string, __VLS_42: any;
type __VLS_Slots = {} & { [K in NonNullable<typeof __VLS_15>]?: (props: typeof __VLS_16) => any } & { [K in NonNullable<typeof __VLS_41>]?: (props: typeof __VLS_42) => any };
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<MentionProps>, {
  options: () => never[];
  prefix: string;
  split: string;
  filterOption: () => (pattern: string, option: MentionOption) => boolean;
  placement: string;
  offset: number;
  popperOptions: () => {};
  props: () => Required<MentionOptionProps>;
  disabled: undefined;
  modelValue: "";
  modelModifiers: () => {};
  type: InputType;
  autocomplete: "off";
  clearIcon: any;
  wordLimitPosition: "inside";
  tabindex: 0;
  validateEvent: true;
  inputStyle: () => {};
  rows: 2;
}>, {
  input: vue.Ref<InputInstance | undefined>;
  tooltip: vue.Ref<TooltipInstance | undefined>;
  dropdownVisible: vue.ComputedRef<boolean>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  search: (pattern: string, prefix: string) => void;
  "update:modelValue": (value: string) => void;
  input: (value: string) => void;
  focus: (evt: FocusEvent) => void;
  blur: (evt: FocusEvent) => void;
  select: (option: MentionOption, prefix: string) => void;
  "whole-remove": (pattern: string, prefix: string) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<MentionProps>, {
  options: () => never[];
  prefix: string;
  split: string;
  filterOption: () => (pattern: string, option: MentionOption) => boolean;
  placement: string;
  offset: number;
  popperOptions: () => {};
  props: () => Required<MentionOptionProps>;
  disabled: undefined;
  modelValue: "";
  modelModifiers: () => {};
  type: InputType;
  autocomplete: "off";
  clearIcon: any;
  wordLimitPosition: "inside";
  tabindex: 0;
  validateEvent: true;
  inputStyle: () => {};
  rows: 2;
}>>> & {
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  onInput?: ((value: string) => any) | undefined;
  onSelect?: ((option: MentionOption, prefix: string) => any) | undefined;
  "onUpdate:modelValue"?: ((value: string) => any) | undefined;
  onSearch?: ((pattern: string, prefix: string) => any) | undefined;
  "onWhole-remove"?: ((pattern: string, prefix: string) => any) | undefined;
}, {
  offset: number;
  props: MentionOptionProps;
  type: InputType;
  split: string;
  modelValue: string;
  placement: "bottom" | "top";
  disabled: boolean;
  modelModifiers: InputModelModifiers;
  autocomplete: string;
  clearIcon: IconPropType;
  wordLimitPosition: "inside" | "outside";
  tabindex: string | number;
  validateEvent: boolean;
  inputStyle: string | false | CSSProperties | vue.StyleValue[] | null;
  rows: number;
  popperOptions: Partial<_popperjs_core0.Options>;
  options: MentionOption[];
  prefix: string | string[];
  filterOption: false | typeof filterOption;
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