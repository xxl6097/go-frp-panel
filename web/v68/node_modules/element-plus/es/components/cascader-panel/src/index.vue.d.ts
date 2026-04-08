import { Node } from "./node.js";
import { CascaderNodePathValue, CascaderNodeValue, CascaderOption, CascaderProps, CascaderValue, ExpandTrigger, LazyLoad, isDisabled, isLeaf } from "./types.js";
import { CascaderMenuInstance } from "./instance.js";
import { CascaderPanelProps } from "./config.js";
import * as vue from "vue";

//#region ../../packages/components/cascader-panel/src/index.vue.d.ts
declare var __VLS_8: {};
type __VLS_Slots = {} & {
  empty?: (props: typeof __VLS_8) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CascaderPanelProps>, {
  options: () => CascaderOption[];
  props: () => CascaderProps;
  border: boolean;
}>, {
  menuList: vue.Ref<CascaderMenuInstance[]>;
  menus: vue.Ref<{
    readonly uid: number;
    readonly level: number;
    readonly value: CascaderNodeValue;
    readonly label: string;
    readonly pathNodes: any[];
    readonly pathValues: CascaderNodeValue[];
    readonly pathLabels: string[];
    childrenData: {
      [x: string]: unknown;
      label?: string | undefined;
      value?: CascaderNodeValue | undefined;
      children?: any[] | undefined;
      disabled?: boolean | undefined;
      leaf?: boolean | undefined;
    }[] | undefined;
    children: any[];
    text: string;
    loaded: boolean;
    checked: boolean;
    indeterminate: boolean;
    loading: boolean;
    readonly data: {
      [x: string]: unknown;
      label?: string | undefined;
      value?: CascaderNodeValue | undefined;
      children?: any[] | undefined;
      disabled?: boolean | undefined;
      leaf?: boolean | undefined;
    };
    readonly config: {
      expandTrigger: ExpandTrigger;
      multiple: boolean;
      checkStrictly: boolean;
      emitPath: boolean;
      lazy: boolean;
      lazyLoad: LazyLoad;
      value: string;
      label: string;
      children: string;
      disabled: string | isDisabled;
      leaf: string | isLeaf;
      hoverThreshold: number;
      checkOnClickNode: boolean;
      checkOnClickLeaf: boolean;
      showPrefix: boolean;
    };
    readonly parent?: any | undefined;
    readonly root: boolean;
    readonly isDisabled: boolean;
    readonly isLeaf: boolean;
    readonly valueByOption: CascaderNodeValue | CascaderNodeValue[];
    appendChild: (childData: CascaderOption) => Node;
    calcText: (allLevels: boolean, separator: string) => string;
    broadcast: (checked: boolean) => void;
    emit: () => void;
    onParentCheck: (checked: boolean) => void;
    onChildCheck: () => void;
    setCheckState: (checked: boolean) => void;
    doCheck: (checked: boolean) => void;
  }[][]>;
  checkedNodes: vue.Ref<{
    readonly uid: number;
    readonly level: number;
    readonly value: CascaderNodeValue;
    readonly label: string;
    readonly pathNodes: any[];
    readonly pathValues: CascaderNodeValue[];
    readonly pathLabels: string[];
    childrenData: {
      [x: string]: unknown;
      label?: string | undefined;
      value?: CascaderNodeValue | undefined;
      children?: any[] | undefined;
      disabled?: boolean | undefined;
      leaf?: boolean | undefined;
    }[] | undefined;
    children: any[];
    text: string;
    loaded: boolean;
    checked: boolean;
    indeterminate: boolean;
    loading: boolean;
    readonly data: {
      [x: string]: unknown;
      label?: string | undefined;
      value?: CascaderNodeValue | undefined;
      children?: any[] | undefined;
      disabled?: boolean | undefined;
      leaf?: boolean | undefined;
    };
    readonly config: {
      expandTrigger: ExpandTrigger;
      multiple: boolean;
      checkStrictly: boolean;
      emitPath: boolean;
      lazy: boolean;
      lazyLoad: LazyLoad;
      value: string;
      label: string;
      children: string;
      disabled: string | isDisabled;
      leaf: string | isLeaf;
      hoverThreshold: number;
      checkOnClickNode: boolean;
      checkOnClickLeaf: boolean;
      showPrefix: boolean;
    };
    readonly parent?: any | undefined;
    readonly root: boolean;
    readonly isDisabled: boolean;
    readonly isLeaf: boolean;
    readonly valueByOption: CascaderNodeValue | CascaderNodeValue[];
    appendChild: (childData: CascaderOption) => Node;
    calcText: (allLevels: boolean, separator: string) => string;
    broadcast: (checked: boolean) => void;
    emit: () => void;
    onParentCheck: (checked: boolean) => void;
    onChildCheck: () => void;
    setCheckState: (checked: boolean) => void;
    doCheck: (checked: boolean) => void;
  }[]>;
  handleKeyDown: (e: KeyboardEvent) => void;
  handleCheckChange: (node: Node, checked: boolean, emitClose?: boolean) => void;
  getFlattedNodes: (leafOnly: boolean) => Node[];
  /**
   * @description get an array of currently selected node,(leafOnly) whether only return the leaf checked nodes, default is `false`
   */
  getCheckedNodes: (leafOnly: boolean) => Node[];
  /**
   * @description clear checked nodes
   */
  clearCheckedNodes: () => void;
  calculateCheckedValue: () => void;
  scrollToExpandingNode: () => void;
  loadLazyRootNodes: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: CascaderValue | null | undefined) => void;
  close: () => void;
  "update:modelValue": (value: CascaderValue | null | undefined) => void;
  "expand-change": (value: CascaderNodePathValue) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CascaderPanelProps>, {
  options: () => CascaderOption[];
  props: () => CascaderProps;
  border: boolean;
}>>> & {
  onChange?: ((value: CascaderValue | null | undefined) => any) | undefined;
  onClose?: (() => any) | undefined;
  "onUpdate:modelValue"?: ((value: CascaderValue | null | undefined) => any) | undefined;
  "onExpand-change"?: ((value: CascaderNodePathValue) => any) | undefined;
}, {
  props: CascaderProps;
  border: boolean;
  options: CascaderOption[];
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