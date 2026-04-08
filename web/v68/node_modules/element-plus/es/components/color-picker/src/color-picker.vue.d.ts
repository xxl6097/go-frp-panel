import { Color } from "../../color-picker-panel/src/utils/color.js";
import { ColorPickerProps } from "./color-picker.js";
import * as vue from "vue";

//#region ../../packages/components/color-picker/src/color-picker.vue.d.ts
declare function show(): void;
declare function hide(): void;
declare function focus(): void;
declare function blur(): void;
declare const __VLS_export: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<ColorPickerProps>, {
  readonly persistent: true;
  readonly modelValue: undefined;
  readonly disabled: undefined;
  readonly clearable: true;
  readonly popperStyle: undefined;
  readonly tabindex: 0;
  readonly teleported: true;
  readonly validateEvent: true;
  readonly valueOnClear: undefined;
}>, {
  /**
   * @description current color object
   */
  color: Color;
  /**
   * @description manually show ColorPicker
   */
  show: typeof show;
  /**
   * @description manually hide ColorPicker
   */
  hide: typeof hide;
  /**
   * @description focus the input element
   */
  focus: typeof focus;
  /**
   * @description blur the input element
   */
  blur: typeof blur;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (val: string | null) => void;
  "update:modelValue": (val: string | null) => void;
  focus: (evt: FocusEvent) => void;
  blur: (evt: FocusEvent) => void;
  clear: () => void;
  activeChange: (val: string | null) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<ColorPickerProps>, {
  readonly persistent: true;
  readonly modelValue: undefined;
  readonly disabled: undefined;
  readonly clearable: true;
  readonly popperStyle: undefined;
  readonly tabindex: 0;
  readonly teleported: true;
  readonly validateEvent: true;
  readonly valueOnClear: undefined;
}>>> & {
  onChange?: ((val: string | null) => any) | undefined;
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  "onUpdate:modelValue"?: ((val: string | null) => any) | undefined;
  onClear?: (() => any) | undefined;
  onActiveChange?: ((val: string | null) => any) | undefined;
}, {
  teleported: boolean;
  modelValue: string | null;
  popperStyle: string | false | vue.CSSProperties | vue.StyleValue[] | null;
  disabled: boolean;
  clearable: boolean;
  tabindex: string | number;
  validateEvent: boolean;
  persistent: boolean;
  valueOnClear: string | number | boolean | Function | null;
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