import { BorderStyle, DividerProps } from "./divider.js";
import * as vue from "vue";

//#region ../../packages/components/divider/src/divider.vue.d.ts
declare var __VLS_1: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_1) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<DividerProps>, {
  direction: string;
  contentPosition: string;
  borderStyle: string;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<DividerProps>, {
  direction: string;
  contentPosition: string;
  borderStyle: string;
}>>>, {
  borderStyle: BorderStyle;
  direction: "horizontal" | "vertical";
  contentPosition: "left" | "center" | "right";
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