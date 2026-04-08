import { CalendarDateType, CalendarProps } from "./calendar.js";
import * as vue from "vue";
import * as dayjs$1 from "dayjs";

//#region ../../packages/components/calendar/src/calendar.vue.d.ts
declare var __VLS_1: {
    date: string;
  }, __VLS_49: {
    data: {
      isSelected: boolean;
      type: string;
      day: string;
      date: Date;
    };
  }, __VLS_60: {
    data: {
      isSelected: boolean;
      type: string;
      day: string;
      date: Date;
    };
  };
type __VLS_Slots = {} & {
  header?: (props: typeof __VLS_1) => any;
} & {
  'date-cell'?: (props: typeof __VLS_49) => any;
} & {
  'date-cell'?: (props: typeof __VLS_60) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<CalendarProps>, {
  controllerType: string;
}>, {
  /** @description currently selected date */selectedDay: vue.WritableComputedRef<dayjs$1.Dayjs | undefined>; /** @description select a specific date */
  pickDay: (day: dayjs$1.Dayjs) => void; /** @description select date */
  selectDate: (type: CalendarDateType) => void; /** @description Calculate the validate date range according to the start and end dates */
  calculateValidatedDateRange: (startDayjs: dayjs$1.Dayjs, endDayjs: dayjs$1.Dayjs) => [dayjs$1.Dayjs, dayjs$1.Dayjs][];
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  "update:modelValue": (value: Date) => void;
  input: (value: Date) => void;
}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<CalendarProps>, {
  controllerType: string;
}>>> & {
  onInput?: ((value: Date) => any) | undefined;
  "onUpdate:modelValue"?: ((value: Date) => any) | undefined;
}, {
  controllerType: "button" | "select";
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