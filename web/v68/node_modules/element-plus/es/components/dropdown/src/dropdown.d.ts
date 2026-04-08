import { EpPropFinalized, EpPropMergeType } from "../../../utils/vue/props/types.js";
import { Arrayable, Nullable } from "../../../utils/typescript.js";
import "../../../utils/index.js";
import { PopperEffect } from "../../popper/src/popper.js";
import { Measurable } from "../../popper/src/constants.js";
import { ButtonProps } from "../../button/src/button.js";
import "../../button/index.js";
import { Placement as Placement$1 } from "../../popper/index.js";
import * as vue from "vue";
import { ComponentInternalInstance, ComputedRef } from "vue";
import { Options } from "@popperjs/core";

//#region ../../packages/components/dropdown/src/dropdown.d.ts
interface IElDropdownInstance {
  instance?: ComponentInternalInstance;
  dropdownSize?: ComputedRef<string>;
  handleClick?: () => void;
  commandHandler?: (...arg: any[]) => void;
  show?: () => void;
  hide?: () => void;
  trigger?: ComputedRef<string>;
  hideOnClick?: ComputedRef<boolean>;
  triggerElm?: ComputedRef<Nullable<HTMLButtonElement>>;
}
declare const dropdownProps: {
  readonly trigger: {
    readonly type: vue.PropType<Arrayable<"hover" | "click" | "contextmenu">>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    readonly __epPropKey: true;
    readonly default: "hover";
  };
  readonly triggerKeys: EpPropFinalized<(new (...args: any[]) => string[]) | (() => string[]) | ((new (...args: any[]) => string[]) | (() => string[]))[], unknown, unknown, () => string[], boolean>;
  readonly virtualTriggering: BooleanConstructor;
  readonly virtualRef: {
    readonly type: vue.PropType<Measurable>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly effect: {
    readonly default: "light";
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    readonly __epPropKey: true;
  };
  readonly type: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => "" | "default" | "info" | "primary" | "success" | "warning" | "text" | "danger") | (() => "" | "default" | "info" | "primary" | "success" | "warning" | "text" | "danger") | ((new (...args: any[]) => "" | "default" | "info" | "primary" | "success" | "warning" | "text" | "danger") | (() => "" | "default" | "info" | "primary" | "success" | "warning" | "text" | "danger"))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly placement: EpPropFinalized<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement$1) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement$1))[], unknown, unknown, "bottom", boolean>;
  readonly popperOptions: EpPropFinalized<(new (...args: any[]) => Partial<Options>) | (() => Partial<Options>) | ((new (...args: any[]) => Partial<Options>) | (() => Partial<Options>))[], unknown, unknown, () => {}, boolean>;
  readonly id: StringConstructor;
  readonly size: EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
  readonly splitButton: BooleanConstructor;
  readonly hideOnClick: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly loop: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly showArrow: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly showTimeout: EpPropFinalized<NumberConstructor, unknown, unknown, 150, boolean>;
  readonly hideTimeout: EpPropFinalized<NumberConstructor, unknown, unknown, 150, boolean>;
  readonly tabindex: EpPropFinalized<(new (...args: any[]) => string | number) | (() => string | number) | ((new (...args: any[]) => string | number) | (() => string | number))[], unknown, unknown, 0, boolean>;
  readonly maxHeight: EpPropFinalized<(new (...args: any[]) => string | number) | (() => string | number) | ((new (...args: any[]) => string | number) | (() => string | number))[], unknown, unknown, "", boolean>;
  readonly popperClass: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | any)[])[])[])[])[])[])[])[])[])[])[]) | (() => string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | any)[])[])[])[])[])[])[])[])[])[])[]) | ((new (...args: any[]) => string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | any)[])[])[])[])[])[])[])[])[])[])[]) | (() => string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | (string | {
      [x: string]: boolean;
    } | any)[])[])[])[])[])[])[])[])[])[])[]))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly popperStyle: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | false | vue.CSSProperties | vue.StyleValue[]) | (() => vue.StyleValue) | ((new (...args: any[]) => string | false | vue.CSSProperties | vue.StyleValue[]) | (() => vue.StyleValue))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly disabled: BooleanConstructor;
  readonly role: EpPropFinalized<StringConstructor, "listbox" | "grid" | "menu" | "tooltip" | "dialog" | "group" | "navigation" | "tree", unknown, "menu", boolean>;
  readonly buttonProps: {
    readonly type: vue.PropType<Partial<ButtonProps>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly appendTo: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>) | ((new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly persistent: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
};
declare const dropdownItemProps: {
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
};
declare const dropdownMenuProps: {
  onKeydown: {
    readonly type: vue.PropType<(e: KeyboardEvent) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
};
declare const FIRST_KEYS: string[];
declare const LAST_KEYS: string[];
declare const FIRST_LAST_KEYS: string[];
//#endregion
export { FIRST_KEYS, FIRST_LAST_KEYS, IElDropdownInstance, LAST_KEYS, dropdownItemProps, dropdownMenuProps, dropdownProps };