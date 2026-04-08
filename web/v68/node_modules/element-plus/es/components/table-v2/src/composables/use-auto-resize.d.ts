import { AutoResizerProps } from "../auto-resizer.js";
import * as vue from "vue";

//#region ../../packages/components/table-v2/src/composables/use-auto-resize.d.ts
declare const useAutoResize: (props: AutoResizerProps) => {
  sizer: vue.Ref<HTMLElement | undefined>;
  width: vue.Ref<number>;
  height: vue.Ref<number>;
};
//#endregion
export { useAutoResize };