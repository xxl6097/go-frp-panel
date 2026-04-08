import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { PopperEffect } from "../../popper/src/popper.js";
import { Node } from "../../cascader-panel/src/node.js";
import { CascaderNodePathValue, CascaderNodeValue, CascaderOption, CascaderProps, CascaderValue } from "../../cascader-panel/src/types.js";
import { CascaderPanelInstance } from "../../cascader-panel/src/instance.js";
import "../../cascader-panel/index.js";
import { CascaderComponentProps } from "./cascader.js";
import { Placement } from "../../popper/index.js";
import * as vue from "vue";
import { ComputedRef, StyleValue } from "vue";

//#region ../../packages/components/cascader/src/cascader.vue.d.ts
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CascaderComponentProps>, {
  options: () => never[];
  props: () => {};
  disabled: undefined;
  clearIcon: any;
  filterMethod: (node: any, keyword: any) => any;
  separator: string;
  showAllLevels: boolean;
  maxCollapseTags: number;
  debounce: number;
  beforeFilter: () => boolean;
  placement: string;
  fallbackPlacements: () => string[];
  teleported: boolean;
  effect: string;
  tagType: string;
  tagEffect: string;
  validateEvent: boolean;
  persistent: boolean;
  showCheckedStrategy: string;
  showPrefix: boolean;
  popperStyle: undefined;
  valueOnClear: undefined;
}>, {
  /**
   * @description get an array of currently selected node,(leafOnly) whether only return the leaf checked nodes, default is `false`
   */
  getCheckedNodes: (leafOnly: boolean) => Node[] | undefined;
  /**
   * @description cascader panel ref
   */
  cascaderPanelRef: vue.Ref<CascaderPanelInstance | undefined>;
  /**
   * @description toggle the visible of popper
   */
  togglePopperVisible: (visible?: boolean) => void;
  /**
   * @description cascader content ref
   */
  contentRef: ComputedRef<HTMLElement | undefined>;
  /**
   * @description selected content text
   */
  presentText: ComputedRef<string>; /** @description focus the input element */
  focus: () => void; /** @description blur the input element */
  blur: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: CascaderValue | null | undefined) => void;
  "update:modelValue": (value: CascaderValue | null | undefined) => void;
  focus: (evt: FocusEvent) => void;
  blur: (evt: FocusEvent) => void;
  clear: () => void;
  visibleChange: (val: boolean) => void;
  expandChange: (val: CascaderValue) => void;
  removeTag: (val: CascaderNodeValue | CascaderNodePathValue) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CascaderComponentProps>, {
  options: () => never[];
  props: () => {};
  disabled: undefined;
  clearIcon: any;
  filterMethod: (node: any, keyword: any) => any;
  separator: string;
  showAllLevels: boolean;
  maxCollapseTags: number;
  debounce: number;
  beforeFilter: () => boolean;
  placement: string;
  fallbackPlacements: () => string[];
  teleported: boolean;
  effect: string;
  tagType: string;
  tagEffect: string;
  validateEvent: boolean;
  persistent: boolean;
  showCheckedStrategy: string;
  showPrefix: boolean;
  popperStyle: undefined;
  valueOnClear: undefined;
}>>> & {
  onChange?: ((value: CascaderValue | null | undefined) => any) | undefined;
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  "onUpdate:modelValue"?: ((value: CascaderValue | null | undefined) => any) | undefined;
  onClear?: (() => any) | undefined;
  onVisibleChange?: ((val: boolean) => any) | undefined;
  onExpandChange?: ((val: CascaderValue) => any) | undefined;
  onRemoveTag?: ((val: CascaderNodeValue | CascaderNodePathValue) => any) | undefined;
}, {
  teleported: boolean;
  props: CascaderProps;
  effect: PopperEffect;
  debounce: number;
  placement: Placement;
  popperStyle: string | false | vue.CSSProperties | StyleValue[] | null;
  disabled: boolean;
  clearIcon: IconPropType;
  validateEvent: boolean;
  separator: string;
  fallbackPlacements: Placement[];
  persistent: boolean;
  options: CascaderOption[];
  valueOnClear: string | number | boolean | Function | null;
  showPrefix: boolean;
  filterMethod: (node: Node, keyword: string) => boolean;
  showAllLevels: boolean;
  maxCollapseTags: number;
  beforeFilter: (value: string) => boolean | Promise<any>;
  tagType: "info" | "primary" | "success" | "warning" | "danger";
  tagEffect: "light" | "dark" | "plain";
  showCheckedStrategy: "parent" | "child";
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