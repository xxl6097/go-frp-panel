import { EpPropFinalized, EpPropMergeType } from "../../../utils/vue/props/types.js";
import "../../../utils/index.js";
import { PopperEffect } from "../../popper/src/popper.js";
import { TooltipInstance } from "../../tooltip/src/tooltip.js";
import "../../tooltip/index.js";
import { Option, OptionType, SelectStates } from "./select.types.js";
import { SelectDropdownInstance } from "./select-dropdown.js";
import { SelectV2Props, TagTooltipProps } from "./defaults.js";
import { Props } from "./useProps.js";
import "../../../index.js";
import { Options, Placement } from "../../popper/index.js";
import * as vue from "vue";
import * as _vueuse_shared0 from "@vueuse/shared";

//#region ../../packages/components/select-v2/src/select.vue.d.ts
declare const _default: typeof __VLS_export;
declare const __VLS_export: vue.DefineComponent<{
  readonly ariaLabel: StringConstructor;
  readonly emptyValues: ArrayConstructor;
  readonly valueOnClear: EpPropFinalized<(new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null) | ((new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null))[], unknown, unknown, undefined, boolean>;
  readonly allowCreate: BooleanConstructor;
  readonly autocomplete: EpPropFinalized<(new (...args: any[]) => "none" | "both" | "inline" | "list") | (() => "none" | "both" | "inline" | "list") | ((new (...args: any[]) => "none" | "both" | "inline" | "list") | (() => "none" | "both" | "inline" | "list"))[], unknown, unknown, "none", boolean>;
  readonly automaticDropdown: BooleanConstructor;
  readonly clearable: BooleanConstructor;
  readonly clearIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly effect: EpPropFinalized<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown, "light", boolean>;
  readonly collapseTags: BooleanConstructor;
  readonly collapseTagsTooltip: BooleanConstructor;
  readonly tagTooltip: EpPropFinalized<(new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps) | ((new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps))[], unknown, unknown, () => {}, boolean>;
  readonly maxCollapseTags: EpPropFinalized<NumberConstructor, unknown, unknown, 1, boolean>;
  readonly defaultFirstOption: BooleanConstructor;
  readonly disabled: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  readonly estimatedOptionHeight: EpPropFinalized<NumberConstructor, unknown, unknown, undefined, boolean>;
  readonly filterable: BooleanConstructor;
  readonly filterMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly height: EpPropFinalized<NumberConstructor, unknown, unknown, 274, boolean>;
  readonly itemHeight: EpPropFinalized<NumberConstructor, unknown, unknown, 34, boolean>;
  readonly id: StringConstructor;
  readonly loading: BooleanConstructor;
  readonly loadingText: StringConstructor;
  readonly modelValue: EpPropFinalized<(new (...args: any[]) => any) | (() => any) | {
    (): any;
    new (): any;
    readonly prototype: any;
  } | ((new (...args: any[]) => any) | (() => any) | {
    (): any;
    new (): any;
    readonly prototype: any;
  })[], unknown, unknown, undefined, boolean>;
  readonly multiple: BooleanConstructor;
  readonly multipleLimit: EpPropFinalized<NumberConstructor, unknown, unknown, 0, boolean>;
  readonly name: StringConstructor;
  readonly noDataText: StringConstructor;
  readonly noMatchText: StringConstructor;
  readonly remoteMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly reserveKeyword: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly options: {
    readonly type: vue.PropType<OptionType[]>;
    readonly required: true;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly placeholder: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly persistent: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
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
  readonly popperOptions: EpPropFinalized<(new (...args: any[]) => Partial<Options>) | (() => Partial<Options>) | ((new (...args: any[]) => Partial<Options>) | (() => Partial<Options>))[], unknown, unknown, () => Partial<Options>, boolean>;
  readonly remote: BooleanConstructor;
  readonly debounce: EpPropFinalized<NumberConstructor, unknown, unknown, 300, boolean>;
  readonly size: {
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "" | "default" | "small" | "large", never>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly props: EpPropFinalized<(new (...args: any[]) => Props) | (() => Props) | ((new (...args: any[]) => Props) | (() => Props))[], unknown, unknown, () => Required<Props>, boolean>;
  readonly valueKey: EpPropFinalized<StringConstructor, unknown, unknown, "value", boolean>;
  readonly scrollbarAlwaysOn: BooleanConstructor;
  readonly validateEvent: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly offset: EpPropFinalized<NumberConstructor, unknown, unknown, 12, boolean>;
  readonly remoteShowSuffix: BooleanConstructor;
  readonly showArrow: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly placement: EpPropFinalized<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement))[], Placement, unknown, "bottom-start", boolean>;
  readonly fallbackPlacements: EpPropFinalized<(new (...args: any[]) => Placement[]) | (() => Placement[]) | ((new (...args: any[]) => Placement[]) | (() => Placement[]))[], unknown, unknown, readonly ["bottom-start", "top-start", "right", "left"], boolean>;
  readonly tagType: {
    readonly default: "info";
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "info" | "primary" | "success" | "warning" | "danger", unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    readonly __epPropKey: true;
  };
  readonly tagEffect: {
    readonly default: "light";
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "light" | "dark" | "plain", unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    readonly __epPropKey: true;
  };
  readonly tabindex: EpPropFinalized<readonly [StringConstructor, NumberConstructor], unknown, unknown, 0, boolean>;
  readonly appendTo: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>) | ((new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly fitInputWidth: EpPropFinalized<readonly [BooleanConstructor, NumberConstructor], unknown, number | boolean, true, boolean>;
  readonly suffixIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}, {
  modelValue: vue.ComputedRef<any>;
  selectedLabel: vue.ComputedRef<string | string[]>;
  calculatorRef: vue.ShallowRef<HTMLElement | undefined>;
  inputStyle: vue.ComputedRef<{
    minWidth: string;
  }>;
  contentId: vue.Ref<string>;
  BORDER_HORIZONTAL_WIDTH: number;
  inputId: vue.Ref<string | undefined>;
  collapseTagSize: vue.ComputedRef<"default" | "small">;
  currentPlaceholder: vue.ComputedRef<string>;
  expanded: vue.Ref<boolean>;
  emptyText: vue.ComputedRef<string | null>;
  popupHeight: vue.ComputedRef<number>;
  debounce: vue.ComputedRef<number>;
  allOptions: vue.ComputedRef<OptionType[]>;
  allOptionsValueMap: vue.ComputedRef<Map<any, any>>;
  filteredOptions: vue.Ref<({
    [x: string]: any;
  } | {
    [x: string]: any;
    created?: boolean | undefined;
  })[]>;
  iconComponent: vue.ComputedRef<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown> | undefined>;
  iconReverse: vue.ComputedRef<string | undefined>;
  tagStyle: vue.ComputedRef<{
    maxWidth: string;
  }>;
  collapseTagStyle: vue.ComputedRef<{
    maxWidth: string;
  }>;
  popperSize: vue.Ref<number>;
  dropdownMenuVisible: vue.WritableComputedRef<boolean>;
  hasModelValue: vue.ComputedRef<boolean>;
  shouldShowPlaceholder: vue.ComputedRef<boolean>;
  selectDisabled: vue.ComputedRef<boolean>;
  selectSize: vue.ComputedRef<"" | "default" | "small" | "large">;
  needStatusIcon: vue.ComputedRef<boolean>;
  showClearBtn: vue.ComputedRef<boolean>;
  states: SelectStates;
  isFocused: vue.Ref<boolean>;
  nsSelect: {
    namespace: vue.ComputedRef<string>;
    b: (blockSuffix?: string) => string;
    e: (element?: string) => string;
    m: (modifier?: string) => string;
    be: (blockSuffix?: string, element?: string) => string;
    em: (element?: string, modifier?: string) => string;
    bm: (blockSuffix?: string, modifier?: string) => string;
    bem: (blockSuffix?: string, element?: string, modifier?: string) => string;
    is: {
      (name: string, state: boolean | undefined): string;
      (name: string): string;
    };
    cssVar: (object: Record<string, string>) => Record<string, string>;
    cssVarName: (name: string) => string;
    cssVarBlock: (object: Record<string, string>) => Record<string, string>;
    cssVarBlockName: (name: string) => string;
  };
  nsInput: {
    namespace: vue.ComputedRef<string>;
    b: (blockSuffix?: string) => string;
    e: (element?: string) => string;
    m: (modifier?: string) => string;
    be: (blockSuffix?: string, element?: string) => string;
    em: (element?: string, modifier?: string) => string;
    bm: (blockSuffix?: string, modifier?: string) => string;
    bem: (blockSuffix?: string, element?: string, modifier?: string) => string;
    is: {
      (name: string, state: boolean | undefined): string;
      (name: string): string;
    };
    cssVar: (object: Record<string, string>) => Record<string, string>;
    cssVarName: (name: string) => string;
    cssVarBlock: (object: Record<string, string>) => Record<string, string>;
    cssVarBlockName: (name: string) => string;
  };
  inputRef: vue.Ref<HTMLElement | undefined>;
  menuRef: vue.Ref<SelectDropdownInstance | undefined>;
  tagMenuRef: vue.Ref<HTMLElement | undefined>;
  tooltipRef: vue.Ref<TooltipInstance | undefined>;
  tagTooltipRef: vue.Ref<TooltipInstance | undefined>;
  selectRef: vue.Ref<HTMLElement | undefined>;
  wrapperRef: vue.ShallowRef<HTMLElement | undefined>;
  selectionRef: vue.Ref<HTMLElement | undefined>;
  prefixRef: vue.Ref<HTMLElement | undefined>;
  suffixRef: vue.Ref<HTMLElement | undefined>;
  collapseItemRef: vue.Ref<HTMLElement | undefined>;
  popperRef: vue.ComputedRef<HTMLElement | undefined>;
  validateState: vue.ComputedRef<"" | "success" | "error" | "validating">;
  validateIcon: vue.ComputedRef<any>;
  showTagList: vue.ComputedRef<Option[]>;
  collapseTagList: vue.ComputedRef<Option[]>;
  debouncedOnInputChange: _vueuse_shared0.PromisifyFn<() => void>;
  deleteTag: (event: MouseEvent, option: Option) => void;
  getLabel: (option: Option) => any;
  getValue: (option: Option) => any;
  getDisabled: (option: Option) => any;
  getValueKey: (item: unknown) => any;
  getIndex: (option: Option) => any;
  handleClear: () => void;
  handleClickOutside: (event: Event) => void;
  handleDel: (e: KeyboardEvent) => void;
  handleEsc: () => void;
  focus: () => void;
  blur: () => void;
  handleMenuEnter: () => Promise<Awaited<R>>;
  handleResize: () => void;
  resetSelectionWidth: () => void;
  updateTooltip: () => void;
  updateTagTooltip: () => void;
  updateOptions: () => void;
  toggleMenu: (event?: Event) => void;
  scrollTo: (index: number) => void;
  onInput: (event: Event) => void;
  onKeyboardNavigate: (direction: "forward" | "backward", hoveringIndex?: number | undefined) => void;
  onKeyboardSelect: () => void;
  onSelect: (option: Option) => void;
  onHover: (idx?: number) => void;
  handleCompositionStart: (event: CompositionEvent) => void;
  handleCompositionEnd: (event: CompositionEvent) => void;
  handleCompositionUpdate: (event: CompositionEvent) => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  "update:modelValue": (val: SelectV2Props["modelValue"]) => boolean;
  change: (val: SelectV2Props["modelValue"]) => boolean;
  'remove-tag': (val: unknown) => boolean;
  'visible-change': (visible: boolean) => boolean;
  focus: (evt: FocusEvent) => boolean;
  blur: (evt: FocusEvent) => boolean;
  clear: () => boolean;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  readonly ariaLabel: StringConstructor;
  readonly emptyValues: ArrayConstructor;
  readonly valueOnClear: EpPropFinalized<(new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null) | ((new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null))[], unknown, unknown, undefined, boolean>;
  readonly allowCreate: BooleanConstructor;
  readonly autocomplete: EpPropFinalized<(new (...args: any[]) => "none" | "both" | "inline" | "list") | (() => "none" | "both" | "inline" | "list") | ((new (...args: any[]) => "none" | "both" | "inline" | "list") | (() => "none" | "both" | "inline" | "list"))[], unknown, unknown, "none", boolean>;
  readonly automaticDropdown: BooleanConstructor;
  readonly clearable: BooleanConstructor;
  readonly clearIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly effect: EpPropFinalized<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown, "light", boolean>;
  readonly collapseTags: BooleanConstructor;
  readonly collapseTagsTooltip: BooleanConstructor;
  readonly tagTooltip: EpPropFinalized<(new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps) | ((new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps))[], unknown, unknown, () => {}, boolean>;
  readonly maxCollapseTags: EpPropFinalized<NumberConstructor, unknown, unknown, 1, boolean>;
  readonly defaultFirstOption: BooleanConstructor;
  readonly disabled: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  readonly estimatedOptionHeight: EpPropFinalized<NumberConstructor, unknown, unknown, undefined, boolean>;
  readonly filterable: BooleanConstructor;
  readonly filterMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly height: EpPropFinalized<NumberConstructor, unknown, unknown, 274, boolean>;
  readonly itemHeight: EpPropFinalized<NumberConstructor, unknown, unknown, 34, boolean>;
  readonly id: StringConstructor;
  readonly loading: BooleanConstructor;
  readonly loadingText: StringConstructor;
  readonly modelValue: EpPropFinalized<(new (...args: any[]) => any) | (() => any) | {
    (): any;
    new (): any;
    readonly prototype: any;
  } | ((new (...args: any[]) => any) | (() => any) | {
    (): any;
    new (): any;
    readonly prototype: any;
  })[], unknown, unknown, undefined, boolean>;
  readonly multiple: BooleanConstructor;
  readonly multipleLimit: EpPropFinalized<NumberConstructor, unknown, unknown, 0, boolean>;
  readonly name: StringConstructor;
  readonly noDataText: StringConstructor;
  readonly noMatchText: StringConstructor;
  readonly remoteMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly reserveKeyword: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly options: {
    readonly type: vue.PropType<OptionType[]>;
    readonly required: true;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly placeholder: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly persistent: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
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
  readonly popperOptions: EpPropFinalized<(new (...args: any[]) => Partial<Options>) | (() => Partial<Options>) | ((new (...args: any[]) => Partial<Options>) | (() => Partial<Options>))[], unknown, unknown, () => Partial<Options>, boolean>;
  readonly remote: BooleanConstructor;
  readonly debounce: EpPropFinalized<NumberConstructor, unknown, unknown, 300, boolean>;
  readonly size: {
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "" | "default" | "small" | "large", never>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly props: EpPropFinalized<(new (...args: any[]) => Props) | (() => Props) | ((new (...args: any[]) => Props) | (() => Props))[], unknown, unknown, () => Required<Props>, boolean>;
  readonly valueKey: EpPropFinalized<StringConstructor, unknown, unknown, "value", boolean>;
  readonly scrollbarAlwaysOn: BooleanConstructor;
  readonly validateEvent: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly offset: EpPropFinalized<NumberConstructor, unknown, unknown, 12, boolean>;
  readonly remoteShowSuffix: BooleanConstructor;
  readonly showArrow: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  readonly placement: EpPropFinalized<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement))[], Placement, unknown, "bottom-start", boolean>;
  readonly fallbackPlacements: EpPropFinalized<(new (...args: any[]) => Placement[]) | (() => Placement[]) | ((new (...args: any[]) => Placement[]) | (() => Placement[]))[], unknown, unknown, readonly ["bottom-start", "top-start", "right", "left"], boolean>;
  readonly tagType: {
    readonly default: "info";
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "info" | "primary" | "success" | "warning" | "danger", unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    readonly __epPropKey: true;
  };
  readonly tagEffect: {
    readonly default: "light";
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "light" | "dark" | "plain", unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    readonly __epPropKey: true;
  };
  readonly tabindex: EpPropFinalized<readonly [StringConstructor, NumberConstructor], unknown, unknown, 0, boolean>;
  readonly appendTo: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>) | ((new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  readonly fitInputWidth: EpPropFinalized<readonly [BooleanConstructor, NumberConstructor], unknown, number | boolean, true, boolean>;
  readonly suffixIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}>> & {
  onChange?: ((val: any) => any) | undefined;
  onFocus?: ((evt: FocusEvent) => any) | undefined;
  onBlur?: ((evt: FocusEvent) => any) | undefined;
  "onUpdate:modelValue"?: ((val: any) => any) | undefined;
  onClear?: (() => any) | undefined;
  "onVisible-change"?: ((visible: boolean) => any) | undefined;
  "onRemove-tag"?: ((val: unknown) => any) | undefined;
}, {
  readonly offset: number;
  readonly teleported: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly props: Props;
  readonly effect: EpPropMergeType<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown>;
  readonly valueKey: string;
  readonly modelValue: any;
  readonly debounce: number;
  readonly placement: EpPropMergeType<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement))[], Placement, unknown>;
  readonly fitInputWidth: EpPropMergeType<readonly [BooleanConstructor, NumberConstructor], unknown, number | boolean>;
  readonly disabled: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly autocomplete: EpPropMergeType<(new (...args: any[]) => "none" | "both" | "inline" | "list") | (() => "none" | "both" | "inline" | "list") | ((new (...args: any[]) => "none" | "both" | "inline" | "list") | (() => "none" | "both" | "inline" | "list"))[], unknown, unknown>;
  readonly clearable: boolean;
  readonly tabindex: EpPropMergeType<readonly [StringConstructor, NumberConstructor], unknown, unknown>;
  readonly validateEvent: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly multiple: boolean;
  readonly fallbackPlacements: Placement[];
  readonly popperOptions: Partial<Options>;
  readonly showArrow: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly persistent: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly height: number;
  readonly loading: boolean;
  readonly valueOnClear: EpPropMergeType<(new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null) | ((new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null))[], unknown, unknown>;
  readonly filterable: boolean;
  readonly collapseTags: boolean;
  readonly maxCollapseTags: number;
  readonly collapseTagsTooltip: boolean;
  readonly tagType: EpPropMergeType<StringConstructor, "info" | "primary" | "success" | "warning" | "danger", unknown>;
  readonly tagEffect: EpPropMergeType<StringConstructor, "light" | "dark" | "plain", unknown>;
  readonly automaticDropdown: boolean;
  readonly allowCreate: boolean;
  readonly remote: boolean;
  readonly multipleLimit: number;
  readonly defaultFirstOption: boolean;
  readonly reserveKeyword: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  readonly tagTooltip: TagTooltipProps;
  readonly remoteShowSuffix: boolean;
  readonly estimatedOptionHeight: number;
  readonly itemHeight: number;
  readonly scrollbarAlwaysOn: boolean;
}, {}>;
//#endregion
export { _default };