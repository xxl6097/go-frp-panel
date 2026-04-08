import { CarouselProps } from "./carousel.js";
import * as vue from "vue";

//#region ../../packages/components/carousel/src/carousel.vue.d.ts
declare var __VLS_40: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_40) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CarouselProps>, {
  initialIndex: number;
  height: string;
  trigger: string;
  autoplay: boolean;
  interval: number;
  indicatorPosition: string;
  arrow: string;
  type: string;
  cardScale: number;
  loop: boolean;
  direction: string;
  pauseOnHover: boolean;
}>, {
  /** @description active slide index */activeIndex: vue.WritableComputedRef<number>; /** @description manually switch slide, index of the slide to be switched to, starting from 0; or the `name` of corresponding `el-carousel-item` */
  setActiveItem: (index: number | string) => void; /** @description switch to the previous slide */
  prev: () => void; /** @description switch to the next slide */
  next: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (current: number, prev: number) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CarouselProps>, {
  initialIndex: number;
  height: string;
  trigger: string;
  autoplay: boolean;
  interval: number;
  indicatorPosition: string;
  arrow: string;
  type: string;
  cardScale: number;
  loop: boolean;
  direction: string;
  pauseOnHover: boolean;
}>>> & {
  onChange?: ((current: number, prev: number) => any) | undefined;
}, {
  type: "" | "card";
  loop: boolean;
  trigger: "hover" | "click";
  direction: "horizontal" | "vertical";
  height: string;
  arrow: "always" | "hover" | "never";
  initialIndex: number;
  autoplay: boolean;
  interval: number;
  indicatorPosition: "" | "none" | "outside";
  cardScale: number;
  pauseOnHover: boolean;
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