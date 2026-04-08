import { EpPropFinalized, EpPropMergeType } from "../../../utils/vue/props/types.js";
import "../../../utils/index.js";
import * as vue from "vue";

//#region ../../packages/components/dropdown/src/dropdown-item.vue.d.ts
declare const _default: typeof __VLS_export;
declare const __VLS_export: vue.DefineComponent<{
  readonly command: EpPropFinalized<readonly [ObjectConstructor, StringConstructor, NumberConstructor], unknown, unknown, () => {}, boolean>;
  readonly disabled: BooleanConstructor;
  readonly divided: BooleanConstructor;
  readonly textValue: StringConstructor;
  readonly icon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}, {
  handleClick: (event: PointerEvent) => void;
  handlePointerMove: (event: PointerEvent) => void;
  handlePointerLeave: (event: PointerEvent) => void;
  propsAndAttrs: vue.ComputedRef<{
    disabled: boolean;
    command: EpPropMergeType<readonly [ObjectConstructor, StringConstructor, NumberConstructor], unknown, unknown>;
    divided: boolean;
    icon: EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown> | undefined;
    textValue: string | undefined;
    onClick: ((...args: any[]) => any) | undefined;
    onPointermove: ((...args: any[]) => any) | undefined;
    onPointerleave: ((...args: any[]) => any) | undefined;
  }>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, ("click" | "pointerleave" | "pointermove")[], "click" | "pointerleave" | "pointermove", vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  readonly command: EpPropFinalized<readonly [ObjectConstructor, StringConstructor, NumberConstructor], unknown, unknown, () => {}, boolean>;
  readonly disabled: BooleanConstructor;
  readonly divided: BooleanConstructor;
  readonly textValue: StringConstructor;
  readonly icon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}>> & {
  onClick?: ((...args: any[]) => any) | undefined;
  onPointermove?: ((...args: any[]) => any) | undefined;
  onPointerleave?: ((...args: any[]) => any) | undefined;
}, {
  readonly disabled: boolean;
  readonly command: EpPropMergeType<readonly [ObjectConstructor, StringConstructor, NumberConstructor], unknown, unknown>;
  readonly divided: boolean;
}, {}>;
//#endregion
export { _default };