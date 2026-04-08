import { PopperEffect } from "./popper.js";
import "../../../index.js";
import { PopperContentProps } from "./content.js";
import { Options as Options$1, Placement as Placement$1 } from "../index.js";
import * as vue from "vue";
import * as _popperjs_core0 from "@popperjs/core";

//#region ../../packages/components/popper/src/content.vue.d.ts
declare var __VLS_13: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_13) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<PopperContentProps>, {
  readonly effect: "dark";
  readonly enterable: true;
  readonly stopPopperMouseEvent: true;
  readonly visible: false;
  readonly pure: false;
  readonly focusOnShow: false;
  readonly trapping: false;
  readonly virtualTriggering: false;
  readonly loop: false;
  readonly style: undefined;
  readonly popperStyle: undefined;
  readonly arrowOffset: 5;
  readonly boundariesPadding: 0;
  readonly gpuAcceleration: true;
  readonly offset: 12;
  readonly placement: "bottom";
  readonly popperOptions: () => {};
  readonly strategy: "absolute";
}>, {
  /**
   * @description popper content element
   */
  popperContentRef: vue.Ref<HTMLElement | undefined>;
  /**
   * @description popperjs instance
   */
  popperInstanceRef: vue.ComputedRef<_popperjs_core0.Instance | undefined>;
  /**
   * @description method for updating popper
   */
  updatePopper: (shouldUpdateZIndex?: boolean) => void;
  /**
   * @description content style
   */
  contentStyle: vue.ComputedRef<vue.StyleValue[]>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  close: () => void;
  focus: () => void;
  blur: () => void;
  mouseleave: (evt: MouseEvent) => void;
  mouseenter: (evt: MouseEvent) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<PopperContentProps>, {
  readonly effect: "dark";
  readonly enterable: true;
  readonly stopPopperMouseEvent: true;
  readonly visible: false;
  readonly pure: false;
  readonly focusOnShow: false;
  readonly trapping: false;
  readonly virtualTriggering: false;
  readonly loop: false;
  readonly style: undefined;
  readonly popperStyle: undefined;
  readonly arrowOffset: 5;
  readonly boundariesPadding: 0;
  readonly gpuAcceleration: true;
  readonly offset: 12;
  readonly placement: "bottom";
  readonly popperOptions: () => {};
  readonly strategy: "absolute";
}>>> & {
  onClose?: (() => any) | undefined;
  onFocus?: (() => any) | undefined;
  onBlur?: (() => any) | undefined;
  onMouseenter?: ((evt: MouseEvent) => any) | undefined;
  onMouseleave?: ((evt: MouseEvent) => any) | undefined;
}, {
  offset: number;
  effect: PopperEffect;
  placement: Placement$1;
  popperStyle: string | false | vue.CSSProperties | vue.StyleValue[] | null;
  visible: boolean;
  style: string | false | vue.CSSProperties | vue.StyleValue[] | null;
  enterable: boolean;
  pure: boolean;
  focusOnShow: boolean;
  trapping: boolean;
  stopPopperMouseEvent: boolean;
  virtualTriggering: boolean;
  loop: boolean;
  boundariesPadding: number;
  gpuAcceleration: boolean;
  popperOptions: Partial<Options$1>;
  strategy: "fixed" | "absolute";
  arrowOffset: number;
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