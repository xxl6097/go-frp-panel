import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { PopperEffect, PopperInstance } from "../../popper/src/popper.js";
import { ButtonType } from "../../button/src/button.js";
import "../../button/index.js";
import { PopconfirmProps } from "./popconfirm.js";
import "../../../index.js";
import * as vue from "vue";

//#region ../../packages/components/popconfirm/src/popconfirm.vue.d.ts
declare var __VLS_23: {
    confirm: (e: MouseEvent) => void;
    cancel: (e: MouseEvent) => void;
  }, __VLS_41: {};
type __VLS_Slots = {} & {
  actions?: (props: typeof __VLS_23) => any;
} & {
  reference?: (props: typeof __VLS_41) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<PopconfirmProps>, {
  confirmButtonType: string;
  cancelButtonType: string;
  icon: () => any;
  iconColor: string;
  hideAfter: number;
  effect: string;
  teleported: boolean;
  width: number;
}>, {
  popperRef: vue.ComputedRef<PopperInstance | undefined>;
  hide: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  confirm: (e: MouseEvent) => void;
  cancel: (e: MouseEvent) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<PopconfirmProps>, {
  confirmButtonType: string;
  cancelButtonType: string;
  icon: () => any;
  iconColor: string;
  hideAfter: number;
  effect: string;
  teleported: boolean;
  width: number;
}>>> & {
  onCancel?: ((e: MouseEvent) => any) | undefined;
  onConfirm?: ((e: MouseEvent) => any) | undefined;
}, {
  teleported: boolean;
  effect: PopperEffect;
  icon: IconPropType;
  hideAfter: number;
  width: string | number;
  confirmButtonType: ButtonType;
  cancelButtonType: ButtonType;
  iconColor: string;
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