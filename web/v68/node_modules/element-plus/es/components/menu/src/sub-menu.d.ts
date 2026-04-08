import { EpPropFinalized, EpPropMergeType } from "../../../utils/vue/props/types.js";
import "../../../utils/index.js";
import * as vue from "vue";
import { CSSProperties, ExtractPropTypes, ExtractPublicPropTypes } from "vue";

//#region ../../packages/components/menu/src/sub-menu.d.ts
declare const subMenuProps: {
  readonly index: {
    readonly type: vue.PropType<string>;
    readonly required: true;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly showTimeout: NumberConstructor;
  readonly hideTimeout: NumberConstructor;
  readonly popperClass: StringConstructor;
  readonly popperStyle: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | CSSProperties) | (() => string | CSSProperties) | ((new (...args: any[]) => string | CSSProperties) | (() => string | CSSProperties))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly disabled: BooleanConstructor;
  readonly teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  readonly popperOffset: NumberConstructor;
  readonly expandCloseIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly expandOpenIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly collapseCloseIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly collapseOpenIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
};
type SubMenuProps = ExtractPropTypes<typeof subMenuProps>;
type SubMenuPropsPublic = ExtractPublicPropTypes<typeof subMenuProps>;
declare const _default: vue.DefineComponent<{
  readonly index: {
    readonly type: vue.PropType<string>;
    readonly required: true;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly showTimeout: NumberConstructor;
  readonly hideTimeout: NumberConstructor;
  readonly popperClass: StringConstructor;
  readonly popperStyle: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | CSSProperties) | (() => string | CSSProperties) | ((new (...args: any[]) => string | CSSProperties) | (() => string | CSSProperties))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly disabled: BooleanConstructor;
  readonly teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  readonly popperOffset: NumberConstructor;
  readonly expandCloseIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly expandOpenIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly collapseCloseIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly collapseOpenIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}, () => vue.VNode<vue.RendererNode, vue.RendererElement, {
  [key: string]: any;
}>, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<ExtractPropTypes<{
  readonly index: {
    readonly type: vue.PropType<string>;
    readonly required: true;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly showTimeout: NumberConstructor;
  readonly hideTimeout: NumberConstructor;
  readonly popperClass: StringConstructor;
  readonly popperStyle: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | CSSProperties) | (() => string | CSSProperties) | ((new (...args: any[]) => string | CSSProperties) | (() => string | CSSProperties))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly disabled: BooleanConstructor;
  readonly teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  readonly popperOffset: NumberConstructor;
  readonly expandCloseIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly expandOpenIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly collapseCloseIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly collapseOpenIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}>>, {
  readonly teleported: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly disabled: boolean;
}, {}>;
//#endregion
export { SubMenuProps, SubMenuPropsPublic, _default, subMenuProps };