import { SplitterPanelProps } from "./split-panel.js";
import * as vue from "vue";

//#region ../../packages/components/splitter/src/split-panel.vue.d.ts
declare var __VLS_1: {}, __VLS_15: {}, __VLS_18: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_1) => any;
} & {
  'start-collapsible'?: (props: typeof __VLS_15) => any;
} & {
  'end-collapsible'?: (props: typeof __VLS_18) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<SplitterPanelProps>, {
  resizable: boolean;
}>, {
  /** @description splitter-panel html element */splitterPanelRef: vue.Ref<HTMLDivElement | undefined>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  "update:size": (value: string | number) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<SplitterPanelProps>, {
  resizable: boolean;
}>>> & {
  "onUpdate:size"?: ((value: string | number) => any) | undefined;
}, {
  resizable: boolean;
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