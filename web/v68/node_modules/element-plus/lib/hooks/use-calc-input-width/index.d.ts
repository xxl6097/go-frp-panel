import * as vue from "vue";

//#region ../../packages/hooks/use-calc-input-width/index.d.ts
declare function useCalcInputWidth(): {
  calculatorRef: vue.ShallowRef<HTMLElement | undefined>;
  calculatorWidth: vue.Ref<number>;
  inputStyle: vue.ComputedRef<{
    minWidth: string;
  }>;
};
//#endregion
export { useCalcInputWidth };