import { AffixProps } from "./affix.js";
import * as vue from "vue";
import * as csstype from "csstype";

//#region ../../packages/components/affix/src/affix.vue.d.ts
declare var __VLS_7: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_7) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<AffixProps>, {
  zIndex: number;
  target: string;
  offset: number;
  position: string;
  appendTo: string;
}>, {
  /** @description update affix status */update: () => void; /** @description update rootRect info */
  updateRoot: () => Promise<void>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (fixed: boolean) => void;
  scroll: (args_0: {
    scrollTop: number;
    fixed: boolean;
  }) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<AffixProps>, {
  zIndex: number;
  target: string;
  offset: number;
  position: string;
  appendTo: string;
}>>> & {
  onChange?: ((fixed: boolean) => any) | undefined;
  onScroll?: ((args_0: {
    scrollTop: number;
    fixed: boolean;
  }) => any) | undefined;
}, {
  zIndex: csstype.Property.ZIndex;
  target: string;
  offset: number;
  position: "top" | "bottom";
  appendTo: string | HTMLElement;
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