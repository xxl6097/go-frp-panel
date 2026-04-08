import { TourMask } from "./types.js";
import { TourStepProps } from "./step.js";
import * as vue from "vue";
import * as _floating_ui_utils0 from "@floating-ui/utils";

//#region ../../packages/components/tour/src/step.vue.d.ts
declare var __VLS_12: {}, __VLS_14: {};
type __VLS_Slots = {} & {
  header?: (props: typeof __VLS_12) => any;
} & {
  default?: (props: typeof __VLS_14) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<TourStepProps>, {
  showClose: undefined;
  showArrow: undefined;
  placement: string;
  mask: undefined;
  scrollIntoViewOptions: undefined;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  close: () => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<TourStepProps>, {
  showClose: undefined;
  showArrow: undefined;
  placement: string;
  mask: undefined;
  scrollIntoViewOptions: undefined;
}>>> & {
  onClose?: (() => any) | undefined;
}, {
  placement: _floating_ui_utils0.Placement;
  showArrow: boolean;
  mask: TourMask;
  showClose: boolean;
  scrollIntoViewOptions: boolean | ScrollIntoViewOptions;
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