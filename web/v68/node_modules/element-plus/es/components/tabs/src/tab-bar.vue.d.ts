import { TabPaneName, TabsPaneContext } from "./constants.js";
import { TabBarProps } from "./tab-bar.js";
import * as vue from "vue";
import { CSSProperties } from "vue";

//#region ../../packages/components/tabs/src/tab-bar.vue.d.ts
declare const __VLS_export: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<TabBarProps>, {
  tabs: () => never[];
  tabRefs: () => {};
}>, {
  /** @description tab root html element */ref: vue.Ref<HTMLDivElement | undefined>; /** @description method to manually update tab bar style */
  update: () => CSSProperties;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<TabBarProps>, {
  tabs: () => never[];
  tabRefs: () => {};
}>>>, {
  tabs: TabsPaneContext[];
  tabRefs: {
    [key: TabPaneName]: HTMLDivElement;
  };
}, {}>;
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
type __VLS_PrettifyLocal<T> = (T extends any ? { [K in keyof T]: T[K] } : { [K in keyof T as K]: T[K] }) & {};
//#endregion
export { _default };