import { TransferDataItem, TransferDirection, TransferFormat, TransferKey, TransferProps, TransferPropsAlias } from "./transfer.js";
import { TransferPanelInstance } from "./transfer-panel.js";
import * as vue from "vue";

//#region ../../packages/components/transfer/src/transfer.vue.d.ts
declare var __VLS_12: {}, __VLS_14: {}, __VLS_65: {}, __VLS_67: {};
type __VLS_Slots = {} & {
  'left-empty'?: (props: typeof __VLS_12) => any;
} & {
  'left-footer'?: (props: typeof __VLS_14) => any;
} & {
  'right-empty'?: (props: typeof __VLS_65) => any;
} & {
  'right-footer'?: (props: typeof __VLS_67) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<TransferProps>, {
  data: () => never[];
  titles: () => [string, string];
  buttonTexts: () => [string, string];
  leftDefaultChecked: () => never[];
  rightDefaultChecked: () => never[];
  modelValue: () => never[];
  format: () => {};
  props: () => {
    label: string;
    key: string;
    disabled: string;
  };
  targetOrder: string;
  validateEvent: boolean;
}>, {
  /** @description clear the filter keyword of a certain panel */clearQuery: (which: TransferDirection) => void; /** @description left panel ref */
  leftPanel: vue.Ref<TransferPanelInstance | undefined>; /** @description right panel ref */
  rightPanel: vue.Ref<TransferPanelInstance | undefined>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (value: TransferKey[], direction: TransferDirection, movedKeys: TransferKey[]) => void;
  "update:modelValue": (value: TransferKey[]) => void;
  "left-check-change": (value: TransferKey[], movedKeys?: TransferKey[] | undefined) => void;
  "right-check-change": (value: TransferKey[], movedKeys?: TransferKey[] | undefined) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<TransferProps>, {
  data: () => never[];
  titles: () => [string, string];
  buttonTexts: () => [string, string];
  leftDefaultChecked: () => never[];
  rightDefaultChecked: () => never[];
  modelValue: () => never[];
  format: () => {};
  props: () => {
    label: string;
    key: string;
    disabled: string;
  };
  targetOrder: string;
  validateEvent: boolean;
}>>> & {
  onChange?: ((value: TransferKey[], direction: TransferDirection, movedKeys: TransferKey[]) => any) | undefined;
  "onUpdate:modelValue"?: ((value: TransferKey[]) => any) | undefined;
  "onLeft-check-change"?: ((value: TransferKey[], movedKeys?: TransferKey[] | undefined) => any) | undefined;
  "onRight-check-change"?: ((value: TransferKey[], movedKeys?: TransferKey[] | undefined) => any) | undefined;
}, {
  props: TransferPropsAlias;
  modelValue: TransferKey[];
  validateEvent: boolean;
  data: TransferDataItem[];
  titles: [string, string];
  format: TransferFormat;
  buttonTexts: [string, string];
  leftDefaultChecked: TransferKey[];
  rightDefaultChecked: TransferKey[];
  targetOrder: "original" | "push" | "unshift";
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