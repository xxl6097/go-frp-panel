import { FormItemProps } from "./form-item.js";
import { FormValidateCallback, FormValidationResult } from "./types.js";
import * as vue from "vue";

//#region ../../packages/components/form/src/form-item.vue.d.ts
declare var __VLS_13: {
    label: string;
  }, __VLS_15: {}, __VLS_23: {
    error: string;
  };
type __VLS_Slots = {} & {
  label?: (props: typeof __VLS_13) => any;
} & {
  default?: (props: typeof __VLS_15) => any;
} & {
  error?: (props: typeof __VLS_23) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<FormItemProps>, {
  labelPosition: string;
  showMessage: boolean;
  required: undefined;
  inlineMessage: undefined;
}>, {
  /**
   * @description Form item size.
   */
  size: vue.ComputedRef<"" | "default" | "small" | "large">;
  /**
   * @description Validation message.
   */
  validateMessage: vue.Ref<string>;
  /**
   * @description Validation state.
   */
  validateState: vue.Ref<"" | "success" | "error" | "validating">;
  /**
   * @description Validate form item.
   */
  validate: (trigger: string, callback?: FormValidateCallback) => FormValidationResult;
  /**
   * @description Remove validation status of the field.
   */
  clearValidate: () => void;
  /**
   * @description Reset current field and remove validation result.
   */
  resetField: () => void;
  /**
   * @description Set initial value for this field. When `resetField` is called, the field will reset to this value.
   */
  setInitialValue: (value: any) => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<FormItemProps>, {
  labelPosition: string;
  showMessage: boolean;
  required: undefined;
  inlineMessage: undefined;
}>>>, {
  required: boolean;
  labelPosition: "left" | "right" | "top" | "";
  inlineMessage: boolean;
  showMessage: boolean;
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