import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { ButtonNativeType, ButtonProps, ButtonType } from "./button.js";
import * as vue from "vue";

//#region ../../packages/components/button/src/button.vue.d.ts
declare var __VLS_11: {}, __VLS_35: {}, __VLS_37: {};
type __VLS_Slots = {} & {
  loading?: (props: typeof __VLS_11) => any;
} & {
  icon?: (props: typeof __VLS_35) => any;
} & {
  default?: (props: typeof __VLS_37) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<ButtonProps>, {
  disabled: undefined;
  type: string;
  nativeType: string;
  loadingIcon: any;
  plain: undefined;
  text: undefined;
  round: undefined;
  dashed: undefined;
  autoInsertSpace: undefined;
  tag: string;
}>, {
  /** @description button html element */ref: vue.Ref<HTMLButtonElement | undefined>; /** @description button size */
  size: vue.ComputedRef<"" | "default" | "small" | "large">; /** @description button type */
  type: vue.ComputedRef<"default" | "" | "info" | "primary" | "success" | "warning" | "text" | "danger">; /** @description button disabled */
  disabled: vue.ComputedRef<boolean>; /** @description whether adding space */
  shouldAddSpace: vue.ComputedRef<boolean>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  click: (evt: MouseEvent) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<ButtonProps>, {
  disabled: undefined;
  type: string;
  nativeType: string;
  loadingIcon: any;
  plain: undefined;
  text: undefined;
  round: undefined;
  dashed: undefined;
  autoInsertSpace: undefined;
  tag: string;
}>>> & {
  onClick?: ((evt: MouseEvent) => any) | undefined;
}, {
  type: ButtonType;
  text: boolean;
  disabled: boolean;
  round: boolean;
  dashed: boolean;
  nativeType: ButtonNativeType;
  loadingIcon: IconPropType;
  plain: boolean;
  autoInsertSpace: boolean;
  tag: string | vue.Component;
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