import { EpPropFinalized, EpPropMergeType } from "../../../utils/vue/props/types.js";
import "../../../utils/index.js";
import { ElCollectionInjectionContext, ElCollectionItemInjectionContext } from "../../collection/src/tokens.js";
import "../../collection/index.js";
import * as vue from "vue";
import { ExtractPropTypes, ExtractPublicPropTypes, StyleValue } from "vue";

//#region ../../packages/components/roving-focus-group/src/roving-focus-group.d.ts
declare const rovingFocusGroupProps: {
  style: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | false | vue.CSSProperties | StyleValue[]) | (() => StyleValue) | ((new (...args: any[]) => string | false | vue.CSSProperties | StyleValue[]) | (() => StyleValue))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  currentTabId: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string) | (() => string | null) | ((new (...args: any[]) => string) | (() => string | null))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  defaultCurrentTabId: StringConstructor;
  loop: BooleanConstructor;
  dir: EpPropFinalized<StringConstructor, string, unknown, string, boolean>;
  orientation: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => "horizontal" | "vertical") | (() => "horizontal" | "vertical" | undefined) | ((new (...args: any[]) => "horizontal" | "vertical") | (() => "horizontal" | "vertical" | undefined))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  onBlur: FunctionConstructor;
  onFocus: FunctionConstructor;
  onMousedown: FunctionConstructor;
};
type ElRovingFocusGroupProps = ExtractPropTypes<typeof rovingFocusGroupProps>;
declare const ElCollection: {
    new (...args: any[]): vue.CreateComponentPublicInstance<Readonly<ExtractPropTypes<{}>>, {}, {}, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & Readonly<ExtractPropTypes<{}>>, {}, true, {}, {}, {
      P: {};
      B: {};
      D: {};
      C: {};
      M: {};
      Defaults: {};
    }, Readonly<ExtractPropTypes<{}>>, {}, {}, {}, {}, {}>;
    __isFragment?: never;
    __isTeleport?: never;
    __isSuspense?: never;
  } & vue.ComponentOptionsBase<Readonly<ExtractPropTypes<{}>>, {}, {}, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, {}, {}, string, {}> & vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & {
    name: string;
    setup(): void;
  }, ElCollectionItem: {
    new (...args: any[]): vue.CreateComponentPublicInstance<Readonly<ExtractPropTypes<{}>>, {}, {}, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & Readonly<ExtractPropTypes<{}>>, {}, true, {}, {}, {
      P: {};
      B: {};
      D: {};
      C: {};
      M: {};
      Defaults: {};
    }, Readonly<ExtractPropTypes<{}>>, {}, {}, {}, {}, {}>;
    __isFragment?: never;
    __isTeleport?: never;
    __isSuspense?: never;
  } & vue.ComponentOptionsBase<Readonly<ExtractPropTypes<{}>>, {}, {}, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, {}, {}, string, {}> & vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & {
    name: string;
    setup(_: unknown, {
      attrs
    }: vue.SetupContext): void;
  }, COLLECTION_INJECTION_KEY: vue.InjectionKey<ElCollectionInjectionContext>, COLLECTION_ITEM_INJECTION_KEY: vue.InjectionKey<ElCollectionItemInjectionContext>;
//#endregion
export { COLLECTION_INJECTION_KEY, COLLECTION_ITEM_INJECTION_KEY, ElRovingFocusGroupProps };