import { TypeComponentsMap } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { AlertProps } from "./alert.js";
import * as vue from "vue";

//#region ../../packages/components/alert/src/alert.vue.d.ts
declare var __VLS_13: {}, __VLS_20: {}, __VLS_22: {};
type __VLS_Slots = {} & {
  icon?: (props: typeof __VLS_13) => any;
} & {
  title?: (props: typeof __VLS_20) => any;
} & {
  default?: (props: typeof __VLS_22) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<AlertProps>, {
  title: string;
  description: string;
  type: string;
  closable: boolean;
  closeText: string;
  effect: string;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  close: (evt: MouseEvent) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<AlertProps>, {
  title: string;
  description: string;
  type: string;
  closable: boolean;
  closeText: string;
  effect: string;
}>>> & {
  onClose?: ((evt: MouseEvent) => any) | undefined;
}, {
  type: keyof typeof TypeComponentsMap;
  title: string;
  description: string;
  closable: boolean;
  closeText: string;
  effect: "light" | "dark";
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