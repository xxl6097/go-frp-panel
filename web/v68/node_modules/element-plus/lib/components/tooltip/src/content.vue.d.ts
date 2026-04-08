import { PopperEffect } from "../../popper/src/popper.js";
import { ElTooltipContentProps } from "./content.js";
import { PopperContentInstance } from "../../popper/src/content.js";
import { Options, Placement } from "../../popper/index.js";
import * as vue from "vue";

//#region ../../packages/components/tooltip/src/content.vue.d.ts
declare var __VLS_32: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_32) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<ElTooltipContentProps>, {
  readonly content: "";
  readonly visible: null;
  readonly teleported: true;
  readonly effect: "dark";
  readonly enterable: true;
  readonly stopPopperMouseEvent: true;
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
  readonly showAfter: 0;
  readonly hideAfter: 200;
  readonly autoClose: 0;
}>, {
  /**
   * @description el-popper-content component instance
   */
  contentRef: vue.Ref<PopperContentInstance | undefined>;
  /**
   * @description validate current focus event is trigger inside el-popper-content
   */
  isFocusInsideContent: (event?: FocusEvent) => boolean | undefined;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<ElTooltipContentProps>, {
  readonly content: "";
  readonly visible: null;
  readonly teleported: true;
  readonly effect: "dark";
  readonly enterable: true;
  readonly stopPopperMouseEvent: true;
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
  readonly showAfter: 0;
  readonly hideAfter: 200;
  readonly autoClose: 0;
}>>>, {
  offset: number;
  teleported: boolean;
  effect: PopperEffect;
  placement: Placement;
  popperStyle: string | false | vue.CSSProperties | vue.StyleValue[] | null;
  visible: boolean | null;
  content: string;
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
  popperOptions: Partial<Options>;
  strategy: "fixed" | "absolute";
  arrowOffset: number;
  showAfter: number;
  hideAfter: number;
  autoClose: number;
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