import * as vue from "vue";
import * as vue_jsx_runtime0 from "vue/jsx-runtime";

//#region ../../packages/components/table-v2/src/components/auto-resizer.d.ts
declare const AutoResizer: vue.DefineComponent<{
  readonly disableWidth: BooleanConstructor;
  readonly disableHeight: BooleanConstructor;
  readonly onResize: {
    readonly type: vue.PropType<(event: {
      height: number;
      width: number;
    }) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}, () => vue_jsx_runtime0.JSX.Element, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  readonly disableWidth: BooleanConstructor;
  readonly disableHeight: BooleanConstructor;
  readonly onResize: {
    readonly type: vue.PropType<(event: {
      height: number;
      width: number;
    }) => void>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}>>, {
  readonly disableWidth: boolean;
  readonly disableHeight: boolean;
}, {}>;
//#endregion
export { AutoResizer };