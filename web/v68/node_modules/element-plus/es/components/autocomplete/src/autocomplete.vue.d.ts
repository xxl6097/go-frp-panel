import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { InputModelModifiers, InputType } from "../../input/src/input.js";
import { InputInstance } from "../../input/src/instance.js";
import "../../input/index.js";
import { TooltipInstance } from "../../tooltip/src/tooltip.js";
import "../../tooltip/index.js";
import { AutocompleteFetchSuggestions, AutocompletePlacement, AutocompleteProps } from "./autocomplete.js";
import * as vue from "vue";
import { StyleValue } from "vue";

//#region ../../packages/components/autocomplete/src/autocomplete.vue.d.ts
declare var __VLS_29: {}, __VLS_32: {}, __VLS_35: {}, __VLS_38: {}, __VLS_41: {}, __VLS_49: {}, __VLS_62: {
    item: Record<string, any>;
  }, __VLS_64: {};
type __VLS_Slots = {} & {
  prepend?: (props: typeof __VLS_29) => any;
} & {
  append?: (props: typeof __VLS_32) => any;
} & {
  prefix?: (props: typeof __VLS_35) => any;
} & {
  suffix?: (props: typeof __VLS_38) => any;
} & {
  header?: (props: typeof __VLS_41) => any;
} & {
  loading?: (props: typeof __VLS_49) => any;
} & {
  default?: (props: typeof __VLS_62) => any;
} & {
  footer?: (props: typeof __VLS_64) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<AutocompleteProps>, {
  valueKey: string;
  modelValue: string;
  debounce: number;
  placement: string;
  fetchSuggestions: () => void;
  triggerOnFocus: boolean;
  loopNavigation: boolean;
  teleported: boolean;
  disabled: undefined;
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
  /** @description the index of the currently highlighted item */highlightedIndex: vue.Ref<number>; /** @description autocomplete whether activated */
  activated: vue.Ref<boolean>; /** @description remote search loading status */
  loading: vue.Ref<boolean>; /** @description el-input component instance */
  inputRef: vue.Ref<InputInstance | undefined>; /** @description el-tooltip component instance */
  popperRef: vue.Ref<TooltipInstance | undefined>; /** @description fetch suggestions result */
  suggestions: vue.Ref<Record<string, any>[]>; /** @description triggers when a suggestion is clicked */
  handleSelect: (item: any) => Promise<void>; /** @description handle keyboard enter event */
  handleKeyEnter: () => Promise<void>; /** @description focus the input element */
  focus: () => void; /** @description blur the input element */
  blur: () => void; /** @description close suggestion */
  close: () => void; /** @description highlight an item in a suggestion */
  highlight: (index: number) => void; /** @description loading suggestion list */
  getData: (queryString: string) => Promise<void>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: string | number) => void;
  "update:modelValue": (value: string | number) => void;
  input: (value: string | number) => void;
  focus: (evt: FocusEvent) => void;
  blur: (evt: FocusEvent) => void;
  clear: () => void;
  select: (item: Record<string, any>) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<AutocompleteProps>, {
  valueKey: string;
  modelValue: string;
  debounce: number;
  placement: string;
  fetchSuggestions: () => void;
  triggerOnFocus: boolean;
  loopNavigation: boolean;
  teleported: boolean;
  disabled: undefined;
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
  onChange?: ((value: string | number) => any) | undefined;
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  onInput?: ((value: string | number) => any) | undefined;
  onSelect?: ((item: Record<string, any>) => any) | undefined;
  "onUpdate:modelValue"?: ((value: string | number) => any) | undefined;
  onClear?: (() => any) | undefined;
}, {
  teleported: boolean;
  type: InputType;
  valueKey: string;
  modelValue: string | number;
  debounce: number;
  placement: AutocompletePlacement;
  fetchSuggestions: AutocompleteFetchSuggestions;
  triggerOnFocus: boolean;
  loopNavigation: boolean;
  disabled: boolean;
  modelModifiers: InputModelModifiers;
  autocomplete: string;
  clearIcon: IconPropType;
  wordLimitPosition: "inside" | "outside";
  tabindex: string | number;
  validateEvent: boolean;
  inputStyle: string | false | vue.CSSProperties | StyleValue[] | null;
  rows: number;
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