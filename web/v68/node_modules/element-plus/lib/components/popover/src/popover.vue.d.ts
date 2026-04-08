import { Arrayable } from "../../../utils/typescript.js";
import "../../../utils/index.js";
import { PopperEffect, PopperInstance } from "../../popper/src/popper.js";
import { TooltipTriggerType } from "../../tooltip/src/trigger.js";
import "../../tooltip/index.js";
import { PopoverProps } from "./popover.js";
import "../../../index.js";
import { Options, Placement } from "../../popper/index.js";
import * as vue from "vue";

//#region ../../packages/components/popover/src/popover.vue.d.ts
declare var __VLS_15: {}, __VLS_18: {
    hide: () => void;
  };
type __VLS_Slots = {} & {
  reference?: (props: typeof __VLS_15) => any;
} & {
  default?: (props: typeof __VLS_18) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<PopoverProps>, {
  readonly trigger: "hover";
  readonly triggerKeys: () => string[];
  readonly placement: "bottom";
  readonly visible: null;
  readonly popperOptions: () => {};
  readonly tabindex: 0;
  readonly content: "";
  readonly popperStyle: undefined;
  readonly enterable: true;
  readonly effect: "light";
  readonly teleported: true;
  readonly width: 150;
  readonly offset: undefined;
  readonly showAfter: 0;
  readonly hideAfter: 200;
  readonly autoClose: 0;
  readonly showArrow: true;
  readonly persistent: true;
}>, {
  /** @description popper ref */popperRef: vue.ComputedRef<PopperInstance | undefined>; /** @description hide popover */
  hide: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  "update:visible": (value: boolean) => void;
  "before-enter": () => void;
  "before-leave": () => void;
  "after-enter": () => void;
  "after-leave": () => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<PopoverProps>, {
  readonly trigger: "hover";
  readonly triggerKeys: () => string[];
  readonly placement: "bottom";
  readonly visible: null;
  readonly popperOptions: () => {};
  readonly tabindex: 0;
  readonly content: "";
  readonly popperStyle: undefined;
  readonly enterable: true;
  readonly effect: "light";
  readonly teleported: true;
  readonly width: 150;
  readonly offset: undefined;
  readonly showAfter: 0;
  readonly hideAfter: 200;
  readonly autoClose: 0;
  readonly showArrow: true;
  readonly persistent: true;
}>>> & {
  "onUpdate:visible"?: ((value: boolean) => any) | undefined;
  "onBefore-enter"?: (() => any) | undefined;
  "onBefore-leave"?: (() => any) | undefined;
  "onAfter-enter"?: (() => any) | undefined;
  "onAfter-leave"?: (() => any) | undefined;
}, {
  offset: number;
  teleported: boolean;
  effect: PopperEffect;
  placement: Placement;
  popperStyle: string | false | vue.CSSProperties | vue.StyleValue[] | null;
  tabindex: string | number;
  visible: boolean | null;
  content: string;
  enterable: boolean;
  popperOptions: Partial<Options>;
  showArrow: boolean;
  persistent: boolean;
  showAfter: number;
  hideAfter: number;
  autoClose: number;
  trigger: Arrayable<TooltipTriggerType>;
  triggerKeys: string[];
  width: string | number;
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