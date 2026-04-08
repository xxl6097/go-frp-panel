import { EpPropFinalized, EpPropMergeType } from "../../../utils/vue/props/types.js";
import "../../../utils/index.js";
import { PopperEffect } from "../../popper/src/popper.js";
import { TooltipInstance } from "../../tooltip/src/tooltip.js";
import "../../tooltip/index.js";
import { ScrollbarInstance } from "../../scrollbar/src/scrollbar.js";
import "../../scrollbar/index.js";
import { Option } from "../../select-v2/src/select.types.js";
import { Props } from "../../select-v2/src/useProps.js";
import { TagTooltipProps } from "./select.js";
import { OptionBasic, OptionPublicInstance, OptionValue, SelectStates } from "./type.js";
import "../../../index.js";
import { Options, Placement } from "../../popper/index.js";
import * as vue from "vue";
import * as _vueuse_shared0 from "@vueuse/shared";

//#region ../../packages/components/select/src/select.vue.d.ts
declare const _default: typeof __VLS_export;
declare const __VLS_export: vue.DefineComponent<{
  ariaLabel: StringConstructor;
  emptyValues: ArrayConstructor;
  valueOnClear: EpPropFinalized<(new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null) | ((new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null))[], unknown, unknown, undefined, boolean>;
  name: StringConstructor;
  id: StringConstructor;
  modelValue: EpPropFinalized<(new (...args: any[]) => string | number | boolean | Record<string, any> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[]) | (() => EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[] | null) | ((new (...args: any[]) => string | number | boolean | Record<string, any> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[]) | (() => EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[] | null))[], unknown, unknown, undefined, boolean>;
  autocomplete: EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
  automaticDropdown: BooleanConstructor;
  size: {
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "" | "default" | "small" | "large", never>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  effect: EpPropFinalized<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown, string, boolean>;
  disabled: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  clearable: BooleanConstructor;
  filterable: BooleanConstructor;
  allowCreate: BooleanConstructor;
  loading: BooleanConstructor;
  popperClass: EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
  popperStyle: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | vue.CSSProperties) | (() => string | vue.CSSProperties) | ((new (...args: any[]) => string | vue.CSSProperties) | (() => string | vue.CSSProperties))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  popperOptions: EpPropFinalized<(new (...args: any[]) => Partial<Options>) | (() => Partial<Options>) | ((new (...args: any[]) => Partial<Options>) | (() => Partial<Options>))[], unknown, unknown, () => Partial<Options>, boolean>;
  remote: BooleanConstructor;
  debounce: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  loadingText: StringConstructor;
  noMatchText: StringConstructor;
  noDataText: StringConstructor;
  remoteMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  filterMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  multiple: BooleanConstructor;
  multipleLimit: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  placeholder: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  defaultFirstOption: BooleanConstructor;
  reserveKeyword: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  valueKey: EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
  collapseTags: BooleanConstructor;
  collapseTagsTooltip: BooleanConstructor;
  tagTooltip: EpPropFinalized<(new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps) | ((new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps))[], unknown, unknown, () => {}, boolean>;
  maxCollapseTags: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  persistent: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  clearIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  fitInputWidth: BooleanConstructor;
  suffixIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  tagType: {
    default: string;
    type: vue.PropType<EpPropMergeType<StringConstructor, "info" | "primary" | "success" | "warning" | "danger", unknown>>;
    required: false;
    validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  tagEffect: {
    default: string;
    type: vue.PropType<EpPropMergeType<StringConstructor, "light" | "dark" | "plain", unknown>>;
    required: false;
    validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  validateEvent: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  remoteShowSuffix: BooleanConstructor;
  showArrow: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  offset: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  placement: EpPropFinalized<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement))[], Placement, unknown, string, boolean>;
  fallbackPlacements: EpPropFinalized<(new (...args: any[]) => Placement[]) | (() => Placement[]) | ((new (...args: any[]) => Placement[]) | (() => Placement[]))[], unknown, unknown, string[], boolean>;
  tabindex: EpPropFinalized<(NumberConstructor | StringConstructor)[], unknown, unknown, number, boolean>;
  appendTo: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>) | ((new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  options: {
    readonly type: vue.PropType<Record<string, any>[]>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  props: EpPropFinalized<(new (...args: any[]) => Props) | (() => Props) | ((new (...args: any[]) => Props) | (() => Props))[], unknown, unknown, () => Required<Props>, boolean>;
}, {
  modelValue: vue.ComputedRef<string | number | boolean | any[] | Record<string, any> | null | undefined>;
  selectedLabel: vue.ComputedRef<string | string[]>;
  calculatorRef: vue.ShallowRef<HTMLElement | undefined>;
  inputStyle: vue.ComputedRef<{
    minWidth: string;
  }>;
  getLabel: (option: Option) => any;
  getValue: (option: Option) => any;
  getOptions: (option: Option) => any;
  getDisabled: (option: Option) => any;
  getOptionProps: (option: Record<string, any>) => {
    label: any;
    value: any;
    disabled: any;
  };
  inputId: vue.Ref<string | undefined>;
  contentId: vue.Ref<string>;
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
  states: {
    inputValue: string;
    options: Map<EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>, OptionPublicInstance> & Omit<Map<EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>, OptionPublicInstance>, keyof Map<any, any>>;
    cachedOptions: Map<EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>, OptionPublicInstance> & Omit<Map<EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>, OptionPublicInstance>, keyof Map<any, any>>;
    optionValues: OptionValue[];
    selected: {
      index: number;
      value: OptionValue;
      currentLabel: OptionPublicInstance["currentLabel"];
      isDisabled?: OptionPublicInstance["isDisabled"] | undefined;
    }[];
    hoveringIndex: number;
    inputHovering: boolean;
    selectionWidth: number;
    collapseItemWidth: number;
    previousQuery: string | null;
    selectedLabel: string;
    menuVisibleOnFocus: boolean;
    isBeforeHide: boolean;
  };
  isFocused: vue.Ref<boolean>;
  expanded: vue.Ref<boolean>;
  optionsArray: vue.ComputedRef<OptionPublicInstance[]>;
  hoverOption: vue.Ref<any>;
  selectSize: vue.ComputedRef<"" | "default" | "small" | "large">;
  filteredOptionsCount: vue.ComputedRef<number>;
  updateTooltip: () => void;
  updateTagTooltip: () => void;
  debouncedOnInputChange: _vueuse_shared0.PromisifyFn<() => void>;
  onInput: (event: Event) => void;
  deletePrevTag: (e: KeyboardEvent) => void;
  deleteTag: (event: MouseEvent, tag: OptionBasic) => void;
  deleteSelected: (event: Event) => void;
  handleOptionSelect: (option: OptionPublicInstance) => void;
  scrollToOption: (option: OptionPublicInstance | OptionPublicInstance[] | SelectStates["selected"]) => void;
  hasModelValue: vue.ComputedRef<boolean>;
  shouldShowPlaceholder: vue.ComputedRef<boolean>;
  currentPlaceholder: vue.ComputedRef<string>;
  mouseEnterEventName: vue.ComputedRef<"mouseenter" | null>;
  needStatusIcon: vue.ComputedRef<boolean>;
  showClearBtn: vue.ComputedRef<boolean>;
  iconComponent: vue.ComputedRef<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown> | undefined>;
  iconReverse: vue.ComputedRef<string>;
  validateState: vue.ComputedRef<"" | "success" | "error" | "validating">;
  validateIcon: vue.ComputedRef<"" | vue.Component>;
  showNewOption: vue.ComputedRef<boolean>;
  updateOptions: () => void;
  collapseTagSize: vue.ComputedRef<"default" | "small">;
  setSelected: () => void;
  selectDisabled: vue.ComputedRef<boolean>;
  emptyText: vue.ComputedRef<string | null>;
  handleCompositionStart: (event: CompositionEvent) => void;
  handleCompositionUpdate: (event: CompositionEvent) => void;
  handleCompositionEnd: (event: CompositionEvent) => void;
  handleKeydown: (e: KeyboardEvent) => void;
  onOptionCreate: (vm: OptionPublicInstance) => void;
  onOptionDestroy: (key: OptionValue, vm: OptionPublicInstance) => void;
  handleMenuEnter: () => void;
  focus: () => void;
  blur: () => void;
  handleClearClick: (event: Event) => void;
  handleClickOutside: (event: Event) => void;
  handleEsc: () => void;
  toggleMenu: (event?: Event) => void;
  selectOption: () => void;
  getValueKey: (item: OptionPublicInstance | SelectStates["selected"][0]) => any;
  navigateOptions: (direction: "prev" | "next") => void;
  dropdownMenuVisible: vue.WritableComputedRef<boolean>;
  showTagList: vue.ComputedRef<{
    index: number;
    value: OptionValue;
    currentLabel: OptionPublicInstance["currentLabel"];
    isDisabled?: OptionPublicInstance["isDisabled"] | undefined;
  }[]>;
  collapseTagList: vue.ComputedRef<{
    index: number;
    value: OptionValue;
    currentLabel: OptionPublicInstance["currentLabel"];
    isDisabled?: OptionPublicInstance["isDisabled"] | undefined;
  }[]>;
  popupScroll: (data: {
    scrollTop: number;
    scrollLeft: number;
  }) => void;
  getOption: (value: OptionValue) => {
    index: number;
    value: EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>;
    currentLabel: any;
  } | {
    index: number;
    value: EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>;
    currentLabel: string | number | boolean;
    readonly isDisabled: boolean;
  };
  tagStyle: vue.ComputedRef<{
    maxWidth: string;
  }>;
  collapseTagStyle: vue.ComputedRef<{
    maxWidth: string;
  }>;
  popperRef: vue.ComputedRef<HTMLElement | undefined>;
  inputRef: vue.Ref<HTMLInputElement | undefined>;
  tooltipRef: vue.Ref<TooltipInstance | undefined>;
  tagTooltipRef: vue.Ref<TooltipInstance | undefined>;
  prefixRef: vue.Ref<HTMLElement | undefined>;
  suffixRef: vue.Ref<HTMLElement | undefined>;
  selectRef: vue.Ref<HTMLElement | undefined>;
  wrapperRef: vue.ShallowRef<HTMLElement | undefined>;
  selectionRef: vue.Ref<HTMLElement | undefined>;
  scrollbarRef: vue.Ref<ScrollbarInstance | undefined>;
  menuRef: vue.Ref<HTMLElement | undefined>;
  tagMenuRef: vue.Ref<HTMLElement | undefined>;
  collapseItemRef: vue.Ref<HTMLElement | undefined>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, ("change" | "update:modelValue" | "focus" | "blur" | "clear" | "visible-change" | "remove-tag" | "popup-scroll")[], "change" | "update:modelValue" | "focus" | "blur" | "clear" | "visible-change" | "remove-tag" | "popup-scroll", vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  ariaLabel: StringConstructor;
  emptyValues: ArrayConstructor;
  valueOnClear: EpPropFinalized<(new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null) | ((new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null))[], unknown, unknown, undefined, boolean>;
  name: StringConstructor;
  id: StringConstructor;
  modelValue: EpPropFinalized<(new (...args: any[]) => string | number | boolean | Record<string, any> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[]) | (() => EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[] | null) | ((new (...args: any[]) => string | number | boolean | Record<string, any> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[]) | (() => EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[] | null))[], unknown, unknown, undefined, boolean>;
  autocomplete: EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
  automaticDropdown: BooleanConstructor;
  size: {
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "" | "default" | "small" | "large", never>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  effect: EpPropFinalized<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown, string, boolean>;
  disabled: EpPropFinalized<BooleanConstructor, unknown, unknown, undefined, boolean>;
  clearable: BooleanConstructor;
  filterable: BooleanConstructor;
  allowCreate: BooleanConstructor;
  loading: BooleanConstructor;
  popperClass: EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
  popperStyle: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | vue.CSSProperties) | (() => string | vue.CSSProperties) | ((new (...args: any[]) => string | vue.CSSProperties) | (() => string | vue.CSSProperties))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  popperOptions: EpPropFinalized<(new (...args: any[]) => Partial<Options>) | (() => Partial<Options>) | ((new (...args: any[]) => Partial<Options>) | (() => Partial<Options>))[], unknown, unknown, () => Partial<Options>, boolean>;
  remote: BooleanConstructor;
  debounce: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  loadingText: StringConstructor;
  noMatchText: StringConstructor;
  noDataText: StringConstructor;
  remoteMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  filterMethod: {
    readonly type: vue.PropType<(query: string) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  multiple: BooleanConstructor;
  multipleLimit: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  placeholder: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  defaultFirstOption: BooleanConstructor;
  reserveKeyword: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  valueKey: EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
  collapseTags: BooleanConstructor;
  collapseTagsTooltip: BooleanConstructor;
  tagTooltip: EpPropFinalized<(new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps) | ((new (...args: any[]) => TagTooltipProps) | (() => TagTooltipProps))[], unknown, unknown, () => {}, boolean>;
  maxCollapseTags: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  teleported: EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
  persistent: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  clearIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  fitInputWidth: BooleanConstructor;
  suffixIcon: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component) | ((new (...args: any[]) => (string | vue.Component) & {}) | (() => string | vue.Component))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  tagType: {
    default: string;
    type: vue.PropType<EpPropMergeType<StringConstructor, "info" | "primary" | "success" | "warning" | "danger", unknown>>;
    required: false;
    validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  tagEffect: {
    default: string;
    type: vue.PropType<EpPropMergeType<StringConstructor, "light" | "dark" | "plain", unknown>>;
    required: false;
    validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  validateEvent: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  remoteShowSuffix: BooleanConstructor;
  showArrow: EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
  offset: EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
  placement: EpPropFinalized<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement))[], Placement, unknown, string, boolean>;
  fallbackPlacements: EpPropFinalized<(new (...args: any[]) => Placement[]) | (() => Placement[]) | ((new (...args: any[]) => Placement[]) | (() => Placement[]))[], unknown, unknown, string[], boolean>;
  tabindex: EpPropFinalized<(NumberConstructor | StringConstructor)[], unknown, unknown, number, boolean>;
  appendTo: {
    readonly type: vue.PropType<EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>) | ((new (...args: any[]) => string | HTMLElement) | (() => EpPropMergeType<(new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement) | ((new (...args: any[]) => string | HTMLElement) | (() => string | HTMLElement))[], unknown, unknown>))[], unknown, unknown>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  options: {
    readonly type: vue.PropType<Record<string, any>[]>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  props: EpPropFinalized<(new (...args: any[]) => Props) | (() => Props) | ((new (...args: any[]) => Props) | (() => Props))[], unknown, unknown, () => Required<Props>, boolean>;
}>> & {
  onChange?: ((...args: any[]) => any) | undefined;
  onFocus?: ((...args: any[]) => any) | undefined;
  onBlur?: ((...args: any[]) => any) | undefined;
  "onUpdate:modelValue"?: ((...args: any[]) => any) | undefined;
  onClear?: ((...args: any[]) => any) | undefined;
  "onVisible-change"?: ((...args: any[]) => any) | undefined;
  "onRemove-tag"?: ((...args: any[]) => any) | undefined;
  "onPopup-scroll"?: ((...args: any[]) => any) | undefined;
}, {
  offset: number;
  teleported: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  props: Props;
  effect: EpPropMergeType<(new (...args: any[]) => string) | (() => PopperEffect) | ((new (...args: any[]) => string) | (() => PopperEffect))[], unknown, unknown>;
  valueKey: string;
  modelValue: EpPropMergeType<(new (...args: any[]) => string | number | boolean | Record<string, any> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[]) | (() => EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[] | null) | ((new (...args: any[]) => string | number | boolean | Record<string, any> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[]) | (() => EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown> | EpPropMergeType<(BooleanConstructor | ObjectConstructor | NumberConstructor | StringConstructor)[], unknown, unknown>[] | null))[], unknown, unknown>;
  debounce: number;
  placement: EpPropMergeType<(new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement) | ((new (...args: any[]) => "top" | "auto" | "bottom" | "bottom-start" | "left" | "right" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => Placement))[], Placement, unknown>;
  popperClass: string;
  fitInputWidth: boolean;
  disabled: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  autocomplete: string;
  clearable: boolean;
  tabindex: EpPropMergeType<(NumberConstructor | StringConstructor)[], unknown, unknown>;
  validateEvent: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  multiple: boolean;
  fallbackPlacements: Placement[];
  popperOptions: Partial<Options>;
  showArrow: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  persistent: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  loading: boolean;
  valueOnClear: EpPropMergeType<(new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null) | ((new (...args: any[]) => string | number | boolean | Function) | (() => string | number | boolean | Function | null))[], unknown, unknown>;
  filterable: boolean;
  collapseTags: boolean;
  maxCollapseTags: number;
  collapseTagsTooltip: boolean;
  tagType: EpPropMergeType<StringConstructor, "info" | "primary" | "success" | "warning" | "danger", unknown>;
  tagEffect: EpPropMergeType<StringConstructor, "light" | "dark" | "plain", unknown>;
  automaticDropdown: boolean;
  allowCreate: boolean;
  remote: boolean;
  multipleLimit: number;
  defaultFirstOption: boolean;
  reserveKeyword: EpPropMergeType<BooleanConstructor, unknown, unknown>;
  tagTooltip: TagTooltipProps;
  remoteShowSuffix: boolean;
}, {}>;
//#endregion
export { _default };