import { SkeletonItemProps } from "./skeleton-item.js";
import * as vue from "vue";

//#region ../../packages/components/skeleton/src/skeleton-item.vue.d.ts
declare const __VLS_export: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<SkeletonItemProps>, {
  variant: string;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<SkeletonItemProps>, {
  variant: string;
}>>>, {
  variant: "circle" | "rect" | "h1" | "h3" | "text" | "caption" | "p" | "image" | "button";
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