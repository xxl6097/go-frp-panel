import { AnchorProps } from "./anchor.js";
import * as vue from "vue";

//#region ../../packages/components/anchor/src/anchor.vue.d.ts
declare var __VLS_1: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_1) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<AnchorProps>, {
  offset: number;
  bound: number;
  duration: number;
  marker: boolean;
  type: string;
  direction: string;
}>, {
  scrollTo: (href?: string) => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (href: string) => void;
  click: (e: MouseEvent, href?: string | undefined) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<AnchorProps>, {
  offset: number;
  bound: number;
  duration: number;
  marker: boolean;
  type: string;
  direction: string;
}>>> & {
  onChange?: ((href: string) => any) | undefined;
  onClick?: ((e: MouseEvent, href?: string | undefined) => any) | undefined;
}, {
  offset: number;
  type: "default" | "underline";
  direction: "vertical" | "horizontal";
  marker: boolean;
  duration: number;
  bound: number;
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