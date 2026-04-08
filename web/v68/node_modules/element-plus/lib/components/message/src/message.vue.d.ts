import { IconPropType } from "../../../utils/vue/icon.js";
import { Mutable } from "../../../utils/typescript.js";
import "../../../utils/index.js";
import { MessagePlacement, MessageProps, MessageType } from "./message.js";
import * as vue from "vue";

//#region ../../packages/components/message/src/message.vue.d.ts
declare function close(): void;
declare var __VLS_27: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_27) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<MessageProps>, Pick<Mutable<{
  readonly customClass: "";
  readonly dangerouslyUseHTMLString: false;
  readonly duration: 3000;
  readonly icon: undefined;
  readonly id: "";
  readonly message: "";
  readonly onClose: undefined;
  readonly showClose: false;
  readonly type: "info";
  readonly plain: false;
  readonly offset: 16;
  readonly placement: undefined;
  readonly zIndex: 0;
  readonly grouping: false;
  readonly repeatNum: 1;
  readonly appendTo: HTMLElement;
}>, "zIndex" | "offset" | "type" | "onClose" | "placement" | "id" | "icon" | "plain" | "message" | "showClose" | "duration" | "customClass" | "dangerouslyUseHTMLString" | "grouping" | "repeatNum">>, {
  visible: vue.Ref<boolean>;
  bottom: vue.ComputedRef<number>;
  close: typeof close;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  destroy: () => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<MessageProps>, Pick<Mutable<{
  readonly customClass: "";
  readonly dangerouslyUseHTMLString: false;
  readonly duration: 3000;
  readonly icon: undefined;
  readonly id: "";
  readonly message: "";
  readonly onClose: undefined;
  readonly showClose: false;
  readonly type: "info";
  readonly plain: false;
  readonly offset: 16;
  readonly placement: undefined;
  readonly zIndex: 0;
  readonly grouping: false;
  readonly repeatNum: 1;
  readonly appendTo: HTMLElement;
}>, "zIndex" | "offset" | "type" | "onClose" | "placement" | "id" | "icon" | "plain" | "message" | "showClose" | "duration" | "customClass" | "dangerouslyUseHTMLString" | "grouping" | "repeatNum">>>> & {
  onDestroy?: (() => any) | undefined;
}, {
  zIndex: number;
  offset: number;
  type: MessageType;
  onClose: () => void;
  placement: MessagePlacement;
  id: string;
  icon: IconPropType;
  plain: boolean;
  message: string | vue.VNode | (() => vue.VNode);
  showClose: boolean;
  duration: number;
  customClass: string;
  dangerouslyUseHTMLString: boolean;
  grouping: boolean;
  repeatNum: number;
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