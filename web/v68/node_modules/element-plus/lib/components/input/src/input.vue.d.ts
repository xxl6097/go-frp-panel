import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { InputAutoSize, InputModelModifiers, InputProps, InputType } from "./input.js";
import * as vue from "vue";
import { StyleValue } from "vue";

//#region ../../packages/components/input/src/input.vue.d.ts
declare var __VLS_1: {}, __VLS_3: {}, __VLS_16: {}, __VLS_69: {};
type __VLS_Slots = {} & {
  prepend?: (props: typeof __VLS_1) => any;
} & {
  prefix?: (props: typeof __VLS_3) => any;
} & {
  suffix?: (props: typeof __VLS_16) => any;
} & {
  append?: (props: typeof __VLS_69) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<InputProps>, {
  readonly disabled: undefined;
  readonly modelValue: "";
  readonly modelModifiers: () => {};
  readonly type: InputType;
  readonly autocomplete: "off";
  readonly clearIcon: any;
  readonly wordLimitPosition: "inside";
  readonly tabindex: 0;
  readonly validateEvent: true;
  readonly inputStyle: () => {};
  readonly rows: 2;
}>, {
  /** @description HTML input element */input: vue.ShallowRef<HTMLInputElement | undefined>; /** @description HTML textarea element */
  textarea: vue.ShallowRef<HTMLTextAreaElement | undefined>; /** @description HTML element, input or textarea */
  ref: vue.ComputedRef<HTMLInputElement | HTMLTextAreaElement | undefined>; /** @description style of textarea. */
  textareaStyle: vue.ComputedRef<StyleValue>; /** @description from props (used on unit test) */
  autosize: vue.Ref<InputAutoSize | undefined>; /** @description is input composing */
  isComposing: vue.Ref<boolean>; /** @description HTML input element native method */
  focus: () => void | undefined; /** @description HTML input element native method */
  blur: () => void | undefined; /** @description HTML input element native method */
  select: () => void; /** @description clear input value */
  clear: (evt?: MouseEvent) => void; /** @description resize textarea. */
  resizeTextarea: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: string, evt?: Event | undefined) => void;
  "update:modelValue": (value: string) => void;
  input: (value: string) => void;
  focus: (evt: FocusEvent) => void;
  blur: (evt: FocusEvent) => void;
  clear: (evt: MouseEvent | undefined) => void;
  mouseleave: (evt: MouseEvent) => void;
  mouseenter: (evt: MouseEvent) => void;
  keydown: (evt: Event | KeyboardEvent) => void;
  compositionstart: (evt: CompositionEvent) => void;
  compositionupdate: (evt: CompositionEvent) => void;
  compositionend: (evt: CompositionEvent) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<InputProps>, {
  readonly disabled: undefined;
  readonly modelValue: "";
  readonly modelModifiers: () => {};
  readonly type: InputType;
  readonly autocomplete: "off";
  readonly clearIcon: any;
  readonly wordLimitPosition: "inside";
  readonly tabindex: 0;
  readonly validateEvent: true;
  readonly inputStyle: () => {};
  readonly rows: 2;
}>>> & {
  onChange?: ((value: string, evt?: Event | undefined) => any) | undefined;
  onCompositionend?: ((evt: CompositionEvent) => any) | undefined;
  onCompositionstart?: ((evt: CompositionEvent) => any) | undefined;
  onCompositionupdate?: ((evt: CompositionEvent) => any) | undefined;
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  onInput?: ((value: string) => any) | undefined;
  onKeydown?: ((evt: Event | KeyboardEvent) => any) | undefined;
  onMouseenter?: ((evt: MouseEvent) => any) | undefined;
  onMouseleave?: ((evt: MouseEvent) => any) | undefined;
  "onUpdate:modelValue"?: ((value: string) => any) | undefined;
  onClear?: ((evt: MouseEvent | undefined) => any) | undefined;
}, {
  type: InputType;
  modelValue: string | number | null;
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