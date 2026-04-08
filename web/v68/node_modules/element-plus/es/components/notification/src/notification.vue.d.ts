import { IconPropType } from "../../../utils/vue/icon.js";
import "../../../utils/index.js";
import { NotificationPosition, NotificationProps, NotificationType } from "./notification.js";
import * as vue from "vue";

//#region ../../packages/components/notification/src/notification.vue.d.ts
declare function close(): void;
declare var __VLS_21: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_21) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<NotificationProps>, {
  customClass: string;
  duration: number;
  id: string;
  message: string;
  offset: number;
  onClick: () => undefined;
  position: string;
  showClose: boolean;
  title: string;
  type: string;
  closeIcon: any;
}>, {
  visible: vue.Ref<boolean>; /** @description close notification */
  close: typeof close;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  destroy: () => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<NotificationProps>, {
  customClass: string;
  duration: number;
  id: string;
  message: string;
  offset: number;
  onClick: () => undefined;
  position: string;
  showClose: boolean;
  title: string;
  type: string;
  closeIcon: any;
}>>> & {
  onDestroy?: (() => any) | undefined;
}, {
  offset: number;
  position: NotificationPosition;
  type: NotificationType;
  title: string;
  id: string;
  onClick: () => void;
  message: string | vue.VNode | (() => vue.VNode);
  closeIcon: IconPropType;
  showClose: boolean;
  duration: number;
  customClass: string;
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