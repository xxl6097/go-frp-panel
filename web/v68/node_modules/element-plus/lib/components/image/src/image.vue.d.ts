import { ImageViewerAction } from "../../image-viewer/src/image-viewer.js";
import "../../image-viewer/index.js";
import { ImageFitType, ImageProps } from "./image.js";
import * as vue from "vue";

//#region ../../packages/components/image/src/image.vue.d.ts
declare function clickHandler(): void;
declare var __VLS_1: {}, __VLS_3: {}, __VLS_14: {}, __VLS_17: {
    activeIndex: number;
    total: number;
  }, __VLS_20: {
    actions: (action: ImageViewerAction, options?: {}) => void;
    prev: () => void;
    next: () => void;
    reset: () => void;
    activeIndex: number;
    setActiveItem: (index: number) => void;
  }, __VLS_23: {
    activeIndex: number;
    src: string;
  };
type __VLS_Slots = {} & {
  error?: (props: typeof __VLS_1) => any;
} & {
  placeholder?: (props: typeof __VLS_3) => any;
} & {
  viewer?: (props: typeof __VLS_14) => any;
} & {
  progress?: (props: typeof __VLS_17) => any;
} & {
  toolbar?: (props: typeof __VLS_20) => any;
} & {
  'viewer-error'?: (props: typeof __VLS_23) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<ImageProps>, {
  src: string;
  fit: string;
  previewSrcList: () => never[];
  initialIndex: number;
  infinite: boolean;
  closeOnPressEscape: boolean;
  zoomRate: number;
  scale: number;
  minScale: number;
  maxScale: number;
}>, {
  /** @description manually open preview */showPreview: typeof clickHandler;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  close: () => void;
  error: (evt: Event) => void;
  show: () => void;
  load: (evt: Event) => void;
  switch: (val: number) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<ImageProps>, {
  src: string;
  fit: string;
  previewSrcList: () => never[];
  initialIndex: number;
  infinite: boolean;
  closeOnPressEscape: boolean;
  zoomRate: number;
  scale: number;
  minScale: number;
  maxScale: number;
}>>> & {
  onClose?: (() => any) | undefined;
  onLoad?: ((evt: Event) => any) | undefined;
  onError?: ((evt: Event) => any) | undefined;
  onShow?: (() => any) | undefined;
  onSwitch?: ((val: number) => any) | undefined;
}, {
  infinite: boolean;
  scale: number;
  src: string;
  fit: ImageFitType;
  initialIndex: number;
  closeOnPressEscape: boolean;
  previewSrcList: string[];
  zoomRate: number;
  minScale: number;
  maxScale: number;
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