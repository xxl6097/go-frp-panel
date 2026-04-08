import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import "./steps.js";
import { StepProps } from "./item.js";
import * as vue from "vue";
import { ComputedRef, Ref, VNode } from "vue";

//#region ../../packages/components/steps/src/item.vue.d.ts
declare var __VLS_1: {}, __VLS_36: {}, __VLS_38: {};
type __VLS_Slots = {} & {
  icon?: (props: typeof __VLS_1) => any;
} & {
  title?: (props: typeof __VLS_36) => any;
} & {
  description?: (props: typeof __VLS_38) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<StepProps>, {
  title: string;
  description: string;
  icon: string;
  status: string;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<StepProps>, {
  title: string;
  description: string;
  icon: string;
  status: string;
}>>>, {
  title: string;
  description: string;
  icon: IconPropType;
  status: "" | "wait" | "process" | "finish" | "error" | "success";
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