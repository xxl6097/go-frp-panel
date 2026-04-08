import { InputInstance } from "../../input/src/instance.js";
import "../../input/index.js";
import { Color } from "./utils/color.js";
import { ColorPickerPanelProps } from "./color-picker-panel.js";
import * as vue from "vue";

//#region ../../packages/components/color-picker-panel/src/color-picker-panel.vue.d.ts
declare function update(): void;
declare var __VLS_38: {};
type __VLS_Slots = {} & {
  footer?: (props: typeof __VLS_38) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<ColorPickerPanelProps>, {
  modelValue: undefined;
  border: boolean;
  validateEvent: boolean;
  showAlpha: boolean;
}>, {
  /**
   * @description current color object
   */
  color: Color;
  /**
   * @description custom input ref
   */
  inputRef: vue.Ref<InputInstance | undefined>;
  /**
   * @description update sub components
   */
  update: typeof update;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  "update:modelValue": (val: string | null) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<ColorPickerPanelProps>, {
  modelValue: undefined;
  border: boolean;
  validateEvent: boolean;
  showAlpha: boolean;
}>>> & {
  "onUpdate:modelValue"?: ((val: string | null) => any) | undefined;
}, {
  modelValue: string | null;
  validateEvent: boolean;
  border: boolean;
  showAlpha: boolean;
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