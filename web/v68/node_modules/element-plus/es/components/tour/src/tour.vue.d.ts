import { TourGap, TourMask } from "./types.js";
import { TourProps } from "./tour.js";
import "../../../index.js";
import * as vue from "vue";
import * as _floating_ui_utils0 from "@floating-ui/utils";

//#region ../../packages/components/tour/src/tour.vue.d.ts
declare var __VLS_28: {}, __VLS_30: {
    current: number;
    total: number;
  };
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_28) => any;
} & {
  indicators?: (props: typeof __VLS_30) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<TourProps>, {
  current: number;
  showArrow: boolean;
  showClose: boolean;
  placement: string;
  mask: boolean;
  gap: () => {
    offset: number;
    radius: number;
  };
  scrollIntoViewOptions: () => {
    block: string;
  };
  appendTo: string;
  closeOnPressEscape: boolean;
  targetAreaClickable: boolean;
}>, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (current: number) => void;
  close: (current: number) => void;
  "update:modelValue": (value: boolean) => void;
  finish: () => void;
  "update:current": (current: number) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<TourProps>, {
  current: number;
  showArrow: boolean;
  showClose: boolean;
  placement: string;
  mask: boolean;
  gap: () => {
    offset: number;
    radius: number;
  };
  scrollIntoViewOptions: () => {
    block: string;
  };
  appendTo: string;
  closeOnPressEscape: boolean;
  targetAreaClickable: boolean;
}>>> & {
  onChange?: ((current: number) => any) | undefined;
  onClose?: ((current: number) => any) | undefined;
  "onUpdate:modelValue"?: ((value: boolean) => any) | undefined;
  onFinish?: (() => any) | undefined;
  "onUpdate:current"?: ((current: number) => any) | undefined;
}, {
  appendTo: string | HTMLElement;
  placement: _floating_ui_utils0.Placement;
  showArrow: boolean;
  gap: TourGap;
  mask: TourMask;
  closeOnPressEscape: boolean;
  showClose: boolean;
  scrollIntoViewOptions: boolean | ScrollIntoViewOptions;
  current: number;
  targetAreaClickable: boolean;
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