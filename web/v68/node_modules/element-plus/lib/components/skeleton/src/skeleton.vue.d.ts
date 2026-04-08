import { SkeletonProps } from "./skeleton.js";
import * as vue from "vue";

//#region ../../packages/components/skeleton/src/skeleton.vue.d.ts
declare var __VLS_1: {
    key: number;
  }, __VLS_13: {};
type __VLS_Slots = {} & {
  template?: (props: typeof __VLS_1) => any;
} & {
  default?: (props: typeof __VLS_13) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<SkeletonProps>, {
  loading: boolean;
  count: number;
  rows: number;
}>, {
  /** @description loading state */uiLoading: vue.Ref<boolean>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<SkeletonProps>, {
  loading: boolean;
  count: number;
  rows: number;
}>>>, {
  rows: number;
  loading: boolean;
  count: number;
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