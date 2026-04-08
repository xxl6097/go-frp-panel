import { Layout } from "./type.js";
import { SplitterProps } from "./splitter.js";
import * as vue from "vue";

//#region ../../packages/components/splitter/src/splitter.vue.d.ts
declare var __VLS_1: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_1) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<SplitterProps>, {
  layout: string;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  resize: (index: number, sizes: number[]) => void;
  collapse: (index: number, type: "end" | "start", sizes: number[]) => void;
  resizeStart: (index: number, sizes: number[]) => void;
  resizeEnd: (index: number, sizes: number[]) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<SplitterProps>, {
  layout: string;
}>>> & {
  onResize?: ((index: number, sizes: number[]) => any) | undefined;
  onCollapse?: ((index: number, type: "end" | "start", sizes: number[]) => any) | undefined;
  onResizeStart?: ((index: number, sizes: number[]) => any) | undefined;
  onResizeEnd?: ((index: number, sizes: number[]) => any) | undefined;
}, {
  layout: Layout;
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