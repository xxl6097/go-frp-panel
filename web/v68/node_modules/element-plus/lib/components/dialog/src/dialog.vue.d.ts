import { DialogProps, DialogTransition } from "./dialog.js";
import * as vue from "vue";

//#region ../../packages/components/dialog/src/dialog.vue.d.ts
declare var __VLS_42: {
    close: () => void;
    titleId: string;
    titleClass: string;
  }, __VLS_44: {}, __VLS_46: {}, __VLS_49: {};
type __VLS_Slots = {} & {
  header?: (props: typeof __VLS_42) => any;
} & {
  title?: (props: typeof __VLS_44) => any;
} & {
  default?: (props: typeof __VLS_46) => any;
} & {
  footer?: (props: typeof __VLS_49) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<DialogProps>, {
  readonly appendTo: "body";
  readonly closeOnClickModal: true;
  readonly closeOnPressEscape: true;
  readonly lockScroll: true;
  readonly modal: true;
  readonly openDelay: 0;
  readonly closeDelay: 0;
  readonly headerAriaLevel: "2";
  readonly transition: undefined;
  readonly alignCenter: undefined;
  readonly draggable: undefined;
  readonly overflow: undefined;
  readonly showClose: true;
  readonly title: "";
  readonly ariaLevel: "2";
}>, {
  /** @description whether the dialog is visible */visible: vue.Ref<boolean>;
  dialogContentRef: vue.Ref<any>;
  resetPosition: () => void;
  handleClose: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  close: () => void;
  "update:modelValue": (value: boolean) => void;
  open: () => void;
  opened: () => void;
  closed: () => void;
  openAutoFocus: () => void;
  closeAutoFocus: () => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<DialogProps>, {
  readonly appendTo: "body";
  readonly closeOnClickModal: true;
  readonly closeOnPressEscape: true;
  readonly lockScroll: true;
  readonly modal: true;
  readonly openDelay: 0;
  readonly closeDelay: 0;
  readonly headerAriaLevel: "2";
  readonly transition: undefined;
  readonly alignCenter: undefined;
  readonly draggable: undefined;
  readonly overflow: undefined;
  readonly showClose: true;
  readonly title: "";
  readonly ariaLevel: "2";
}>>> & {
  onClose?: (() => any) | undefined;
  "onUpdate:modelValue"?: ((value: boolean) => any) | undefined;
  onOpen?: (() => any) | undefined;
  onOpened?: (() => any) | undefined;
  onClosed?: (() => any) | undefined;
  onOpenAutoFocus?: (() => any) | undefined;
  onCloseAutoFocus?: (() => any) | undefined;
}, {
  appendTo: string | HTMLElement;
  title: string;
  transition: DialogTransition;
  overflow: boolean;
  closeOnClickModal: boolean;
  closeOnPressEscape: boolean;
  lockScroll: boolean;
  modal: boolean;
  openDelay: number;
  closeDelay: number;
  headerAriaLevel: string;
  alignCenter: boolean;
  draggable: boolean;
  showClose: boolean;
  ariaLevel: string;
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